package ygo

import (
	"time"

	"github.com/wzshiming/dispatcher"
	"github.com/wzshiming/rego"
	"github.com/wzshiming/rego/agent"
)

type Player struct {
	//
	dispatcher.Events
	MsgChan
	Name    string         // 用户名
	Session *agent.Session // 会话
	Phases  lp_type

	// 规则属性
	Index     int           // 玩家索引
	game      *YGO          // 属于游戏
	OverTime  time.Duration // 允许超出的时间
	WaitTime  time.Duration // 每次动作等待的时间
	ReplyTime time.Duration // 回应等待的时间
	PassTime  time.Duration // 经过的时间
	// 基础属性
	Hp        int  // 生命值
	Camp      int  // 阵营
	RoundSize int  // 回合数
	DrawSize  uint // 抽卡数
	MaxHp     uint // 最大生命值
	MaxSdi    int  // 最大手牌

	// 卡牌区
	Deck    *Group // 卡组 40 ~ 60
	Hand    *Group // 手牌
	Extra   *Group // 额外卡组 <= 15 融合怪物 同调怪物 超量怪物
	Side    *Group // 副卡组 <= 15
	Removed *Group // 排除卡
	Grave   *Group // 墓地
	Mzone   *Group // 怪物卡区 5
	Szone   *Group // 魔法卡陷阱卡区 5
	Field   *Group // 场地卡

	lastSummonRound int // 最后召唤回合
	//	pending   map[uint]*Card
	//	cardevent map[string][]*Card
	// 是否失败
	fail bool
}

func NewPlayer(yg *YGO) *Player {
	pl := &Player{
		Events:    dispatcher.NewForkEvent(yg.Fork),
		Camp:      1,
		Hp:        4000,
		DrawSize:  1,
		MaxHp:     ^uint(0),
		MaxSdi:    6,
		OverTime:  time.Second * 120,
		WaitTime:  time.Second * 60,
		ReplyTime: time.Second * 20,
		//		pending:   make(map[uint]*Card),
		//		cardevent: make(map[string][]*Card),
	}
	var pr uint
	pl.MsgChan = NewMsgChan(func(m MsgCode) bool {
		//rego.INFO(m)

		if m.Uniq != 0 {
			if ca := pl.Game().GetCard(m.Uniq); ca != nil {
				if m.Method == uint(LI_Over) {
					if pr != 0 {
						pl.GetTarget().CallAll(touch(pr, 1, 1, 1))
						pr = 0
					}
					pl.GetTarget().Call(touch(m.Uniq, -1, -100, 100))
				} else if m.Method == uint(LI_Out) {
					if pr == m.Uniq {
						pr = 0
					}
					pl.GetTarget().CallAll(touch(m.Uniq, 1, 1, 1))
				} else {
					return true
				}
			}
		} else {
			return true
		}
		return false
	})
	pl.Deck = NewGroup(pl, LL_Deck)
	pl.Hand = NewGroup(pl, LL_Hand)
	pl.Extra = NewGroup(pl, LL_Extra)
	pl.Removed = NewGroup(pl, LL_Removed)
	pl.Grave = NewGroup(pl, LL_Grave)
	pl.Mzone = NewGroup(pl, LL_Mzone)
	pl.Szone = NewGroup(pl, LL_Szone)
	pl.Field = NewGroup(pl, LL_Field)

	pl.AddEvent(RoundBegin, func() {
		pl.MsgPub("{self}进入第{round}回合", Arg{"round": pl.GetRound()})
	})

	pl.AddEvent(DP, func() {
		pl.MsgPub("{self}进入抽牌阶段", nil)
	})

	pl.AddEvent(SP, func() {
		pl.MsgPub("{self}进入预备阶段", nil)
	})

	pl.AddEvent(MP, func() {
		pl.MsgPub("{self}进入主阶段", nil)
	})

	pl.AddEvent(BP, func() {
		pl.MsgPub("{self}进入战斗阶段", nil)
	})

	pl.AddEvent(EP, func() {
		pl.MsgPub("{self}进入结束阶段", nil)
	})

	pl.AddEvent(DP, pl.draw)
	pl.AddEvent(SP, pl.standby)
	pl.AddEvent(MP, pl.main)
	pl.AddEvent(BP, pl.battle)
	pl.AddEvent(EP, pl.end)

	return pl
}

func (pl *Player) Dispatch(eventName string, args ...interface{}) {
	yg := pl.Game()
	yg.Chain(eventName, nil, pl, append(args))
	pl.Events.Dispatch(eventName, args...)
}

func (pl *Player) Game() *YGO {
	return pl.game
}

func (pl *Player) Msg(fmts string, a Arg) {
	if a == nil {
		a = map[string]interface{}{}
	}
	if a["self"] == nil {
		a["self"] = pl.Name
	}
	if a["rival"] == nil {
		a["rival"] = pl.GetTarget().Name
	}
	nap(1)
	pl.Call(message(fmts, a))
}

func (pl *Player) MsgPub(fmts string, a Arg) {
	if a == nil {
		a = map[string]interface{}{}
	}
	if a["self"] == nil {
		a["self"] = pl.Name
	}
	if a["rival"] == nil {
		a["rival"] = pl.GetTarget().Name
	}
	nap(2)
	pl.CallAll(message(fmts, a))
}

func (pl *Player) Fail() {
	pl.fail = true
	panic("Fail")
}

func (pl *Player) IsFail() bool {
	return pl.fail
}

func (pl *Player) ForEachPlayer(fun func(p *Player)) {
	pl.Game().ForEachPlayer(fun)
}

func (pl *Player) Chain(eventName string, ca *Card, cs *Cards, a []interface{}) bool {
	t := pl.Phases
	defer func() {
		pl.Phases = t
		pl.CallAll(flashPhases(pl))
	}()
	pl.Phases = LP_Chain
	pl.CallAll(flashPhases(pl))

	pl.ResetReplyTime()
	pl.MsgPub("等待{self}连锁", nil)
	css := NewCards()
	cs.ForEach(func(c *Card) bool {
		if c != ca && c.GetSummoner() == pl {
			css.EndPush(c)
		}
		return true
	})

	if wi := pl.SelectWill(); wi.Uniq != 0 {
		//pl.MsgPub("{self}选择{method}", Arg{"method": wi.Method})
		css.ForEach(func(c *Card) bool {
			if c.ToUint() != wi.Uniq {
				return true
			}

			e := func() {
				pl.MsgPub("{self}连锁{event}", Arg{"self": c.ToUint(), "event": eventName})
				c.Events.Dispatch(Pay, append(a, eventName)...)
				c.Events.Dispatch(Trigger, append(a, wi.Method)...)
			}
			if ca.Priority() >= c.Priority() {
				ca.OnlyOnce(eventName, e, c)
			} else {
				e()
			}
			return false

		})

		//pl.MsgPub("{self}错误的连锁", nil)
	} else {
		pl.MsgPub("{self}不连锁", nil)
	}
	return false
}

func (pl *Player) GetRound() int {
	return pl.RoundSize
}

func (pl *Player) round() (err error) {
	defer func() {
		if x := recover(); x != nil {
			if _, ok := x.(string); ok {

			} else {
				rego.DebugStack()
			}
		}
	}()
	pl.RoundSize++
	pl.CallAll(flagName(pl))
	pl.Dispatch(RoundBegin, pl)
	pl.Dispatch(DP, LP_Draw)
	pl.Dispatch(SP, LP_Standby)
	pl.Dispatch(MP, LP_Main1)
	pl.Dispatch(BP, LP_Battle)
	pl.Dispatch(MP, LP_Main2)
	pl.Dispatch(EP, LP_End)
	pl.Dispatch(RoundEnd, pl)
	return
}

func (pl *Player) draw(lp lp_type) {
	pl.Phases = lp
	pl.CallAll(flashPhases(pl))

	pl.ActionDraw(1)

}

func (pl *Player) standby(lp lp_type) {
	pl.Phases = lp
	pl.CallAll(flashPhases(pl))

}

func (pl *Player) main(lp lp_type) {
	pl.Phases = lp
	//pl.ClearCode()
	pl.ResetWaitTime()
	for {
		p := pl.SelectWill()
		if p.Uniq == 0 {
			if p.Method == uint(LP_Battle) && lp == LP_Main1 {
				break
			} else if p.Method == uint(LP_End) {
				if lp == LP_Main1 {
					pl.StopOnce(MP)
					pl.StopOnce(BP)
				}
				break
			}
			if p.Method == 0 {
				break
			}
			continue
		}

		if t := pl.Hand.ExistForUniq(p.Uniq); t != nil {
			if p.Method == uint(LI_Use1) {
				t.Dispatch(Use1, t)
			} else if p.Method == uint(LI_Use2) {
				t.Dispatch(Use2, t)
			}
		} else if t := pl.Mzone.ExistForUniq(p.Uniq); t != nil {
			t.Dispatch(Expression)
		} else if t := pl.Szone.ExistForUniq(p.Uniq); t != nil {
			t.Dispatch(Onset)
		} else {
			pl.Msg("非法目标", nil)
		}
	}
}

func (pl *Player) battle(lp lp_type) {
	pl.Phases = lp
	//pl.ClearCode()
	pl.ResetWaitTime()
	for {
		i := pl.SelectWill()
		if i.Uniq == 0 {
			if i.Method == uint(LP_Main2) {
				break
			} else if i.Method == uint(LP_End) {
				pl.StopOnce(MP)
				break
			}
			if i.Method == 0 {
				break
			}
			continue
		}

		t1 := pl.Mzone.ExistForUniq(i.Uniq)
		if t1 == nil || !t1.IsFaceUpAttack() {
			pl.Msg("请选择怪兽区正面攻击表示的怪兽", nil)
			continue
		}

		if !t1.IsCanAttack() {
			pl.Msg("当前怪兽不能攻击", nil)
			continue
		}
		if pl.GetTarget().Mzone.Len() != 0 {
			pl.Msg("选择要攻击的目标", nil)
			if j := pl.SelectWill(); j.Uniq != 0 {
				t2 := pl.Game().GetCard(j.Uniq)
				if pl.GetTarget().Mzone.IsExistCard(t2) {
					t1.Dispatch(Declaration, t2)
				} else {
					pl.Msg("请选择对方怪兽区的怪兽", nil)
				}
			}
		} else {
			t1.Dispatch(Declaration)
		}
	}

}

func (pl *Player) end(lp lp_type) {
	pl.Phases = lp
	if i := pl.Hand.Len() - pl.MaxSdi; i > 0 {
		pl.ResetReplyTime()
		pl.Msg("请{self}选择丢弃的手牌", nil)
		for k := 0; k != i; k++ {
			ca := pl.SelectFor(pl.Hand)
			if ca == nil {
				ca = pl.Hand.EndPop()
			}
			ca.Dispatch(Discard)
		}
	}
}

func (pl *Player) init() {
	pl.ActionDraw(5)
}

func (pl *Player) initDeck(a []uint, b []uint) {
	if pl.Deck.Len() > 0 {
		return
	}
	pl.Game().CardVer.Deck(pl.Deck, pl, a)
	pl.Game().CardVer.Deck(pl.Extra, pl, b)
	pl.ActionShuffle()
}

func (pl *Player) GetHp() int {
	return pl.Hp
}

func (pl *Player) ChangeHp(i int) {
	pl.Dispatch(HPChange, pl, i)

	if i < 0 {
		pl.MsgPub("{self}受到{num}基本分伤害！", Arg{"num": -i})
	} else if i > 0 {
		pl.MsgPub("{self}受到{num}基本分回复！", Arg{"num": i})
	}
	pl.Hp += i
	if pl.Hp < 0 {
		pl.Fail()
	} else {
		pl.CallAll(setFace(map[string]interface{}{pl.Name: pl.Hp}))
	}

}

func (pl *Player) GetTarget() *Player {
	if pl.Index == 0 {
		return pl.Game().GetPlayerForIndex(1)
	}
	return pl.Game().GetPlayerForIndex(0)
}

func (pl *Player) ActionShuffle() {
	pl.Deck.Shuffle()
}

func (pl *Player) ActionDraw(s int) {
	if s <= 0 {
		return
	}
	for i := 0; i != s; i++ {
		if pl.Deck.Len() == 0 {
			pl.Fail()
			return
		}
		pl.Dispatch(DrawNum, pl)
		t := pl.Deck.EndPop()
		pl.Hand.EndPush(t)
	}
	pl.Dispatch(Draw, pl)
}

func (pl *Player) Call(method string, reply interface{}) error {
	return pl.Game().Room.Push(Call{
		Method: method,
		Args:   reply,
	}, pl.Session)
}

func (pl *Player) CallAll(method string, reply interface{}) error {
	return pl.Game().CallAll(method, reply)
}

func (pl *Player) OutTime() {
	pl.PassTime = 0
}

func (pl *Player) ResetReplyTime() {
	pl.PassTime = pl.ReplyTime
}

func (pl *Player) ResetWaitTime() {
	pl.PassTime = pl.WaitTime
}

func (pl *Player) SelectWill() (p MsgCode) {
	//pl.ClearCode()
	pl.CallAll(flashPhases(pl))
	for {
		select {
		case <-time.After(time.Second):
			pl.PassTime -= time.Second
			if pl.PassTime <= 0 {
				return
			}
		case p = <-pl.GetCode():
			return
		}
	}
	return
}

func (pl *Player) IsCanSummon() bool {
	return pl.lastSummonRound < pl.GetRound()
}
func (pl *Player) SetCanSummon() {
	pl.lastSummonRound = 0
}

func (pl *Player) SetNotCanSummon() {
	pl.lastSummonRound = pl.GetRound()
}

func (pl *Player) Select() (t *Card) {
	//pl.ClearCode()
	pl.CallAll(flashPhases(pl))
	for {
		select {
		case <-time.After(time.Second):
			pl.PassTime -= time.Second
			if pl.PassTime <= 0 {
				return
			}
		case p := <-pl.GetCode():
			t = pl.Game().GetCard(p.Uniq)
			return
		}
	}
	return
}

func (pl *Player) SelectForCards(ca *Cards) *Card {
	defer pl.Call(setPick(nil, pl))
	pl.Call(setPick(ca, pl))
	//pl.Msg("从下列卡牌选择", nil)
	if c := pl.Select(); c != nil {
		for _, v := range *ca {
			if v == c {
				return c
			}
		}
	}
	return nil
}

//func (pl *Player) SelectFor(cp ...*Group) *Card {
//	s := []string{}
//	for _, v := range cp {
//		s = append(s, v.GetOwner().Name+" "+string(v.GetName()))
//	}

//	pl.Msg("从下列区域选择卡牌 {c1}", Arg{"c1": s})
//	cs := NewCards()
//	for _, v := range cp {
//		v.ForEach(func(c *Card) bool {
//			cs.EndPush(c)
//			return true
//		})
//	}

//	return pl.SelectForCards(cs)
//}

func (pl *Player) SelectFor(ci ...interface{}) *Card {
	s := []string{}
	as := []Action{}
	cp := []*Group{}
	for _, v := range ci {
		switch t := v.(type) {
		case *Group:
			s = append(s, t.GetOwner().Name+" "+string(t.GetName()))
			cp = append(cp, t)
		case Action:
			as = append(as, t)
		}
	}

	pl.Msg("从下列区域选择卡牌 {c1}", Arg{"c1": s})
	if c := pl.Select(); c != nil {

		for _, v := range cp {
			if v.IsExistCard(c) {
				for _, a := range as {
					if !a.Call(c) {
						return nil
					}
				}
				return c
			}
		}
	}
	return nil
}
