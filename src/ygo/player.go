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
	Phases  LP_TYPE

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

	isChain bool
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
		if m.Uniq == 1 {
			return true
		}
		if m.Uniq != 0 {
			ca := pl.Game().GetCard(m.Uniq)
			if ca != nil {
				if m.Method == uint(LI_Over) {
					if pr != 0 {
						pl.GetTarget().CallAll(Touch(pr, 1, 1, 1))
						pr = 0
					}
					pl.GetTarget().Call(Touch(m.Uniq, -1, -100, 100))
				} else if m.Method == uint(LI_Out) {
					if pr == m.Uniq {
						pr = 0
					}
					pl.GetTarget().CallAll(Touch(m.Uniq, 1, 1, 1))
				} else {
					return true
				}
			}

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
		pl.MsgPub("{self}开始第{round}回合", Arg{"round": pl.GetRound()})
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

	pl.Call(Message(fmts, a))
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
	pl.CallAll(Message(fmts, a))
}

func (pl *Player) Fail() {
	pl.fail = true
}

func (pl *Player) IsFail() bool {
	return pl.fail
}

func (pl *Player) ForEachPlayer(fun func(p *Player)) {
	pl.Game().ForEachPlayer(fun)
}

func (pl *Player) Chain(eventName string, cs []*Card, a []interface{}) bool {
	pl.ResetReplyTime()
	pl.MsgPub("等待{self}连锁", nil)
	if wi := pl.SelectWill(); wi.Uniq != 0 {
		//pl.MsgPub("{self}选择{method}", Arg{"method": wi.Method})
		for _, v := range cs {
			if v.ToUint() == wi.Uniq {
				if v.GetSummoner() == pl {
					pl.MsgPub("{self}连锁{event}", Arg{"self": v.GetId(), "event": eventName})
					v.Events.Dispatch(Pay, append(a, eventName)...)
					v.Events.Dispatch(Trigger+eventName, append(a, wi.Method)...)
					return true
				}
			}
		}
		pl.MsgPub("{self}错误的连锁", nil)
	} else {
		pl.MsgPub("{self}不连锁", nil)
	}
	return false
}

func (pl *Player) GetRound() int {
	return pl.RoundSize
}

func (pl *Player) round() (err error) {
	pl.RoundSize++
	pl.CallAll("flagName", map[string]interface{}{
		"round":  pl.GetRound(),
		"player": pl.Index,
	})
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

func (pl *Player) draw(lp LP_TYPE) bool {
	defer func() {
		if x := recover(); x != nil {
			rego.DebugStack()
		}
	}()
	pl.Phases = lp
	pl.CallAll(flashPhases(pl))
	if pl.Deck.Len() == 0 {
		pl.Fail()
		return false
	}
	pl.ActionDraw(1)
	return true
}

func (pl *Player) standby(lp LP_TYPE) bool {
	defer func() {
		if x := recover(); x != nil {
			rego.DebugStack()
		}
	}()
	pl.Phases = lp
	pl.CallAll(flashPhases(pl))
	return true
}

func (pl *Player) main(lp LP_TYPE) bool {
	defer func() {
		if x := recover(); x != nil {
			rego.DebugStack()
		}
	}()
	pl.Phases = lp
	//pl.ClearCode()
	pl.ResetWaitTime()
	for {
		p := pl.SelectWill()
		if p.Uniq == 0 {
			break
		}
		if p.Uniq == 1 {
			if p.Method == uint(LP_Battle) && lp == LP_Main1 {
				break
			} else if p.Method == uint(LP_End) {
				if lp == LP_Main1 {
					pl.StopOnce(MP)
					pl.StopOnce(BP)
				}
				break
			}
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
	return true
}

func (pl *Player) battle(lp LP_TYPE) bool {
	defer func() {
		if x := recover(); x != nil {
			rego.DebugStack()
		}
	}()
	pl.Phases = lp
	//pl.ClearCode()
	pl.ResetWaitTime()
	for {
		i := pl.SelectWill()
		if i.Uniq == 0 {
			break
		}
		if i.Uniq == 1 {
			if i.Method == uint(LP_Main2) {
				break
			} else if i.Method == uint(LP_End) {
				pl.StopOnce(MP)
				break
			}
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
	return true
}

func (pl *Player) end(lp LP_TYPE) bool {
	defer func() {
		if x := recover(); x != nil {
			rego.DebugStack()
		}
	}()
	pl.Phases = lp
	pl.ResetReplyTime()
	pl.Msg("选择丢弃的手牌", nil)
	if i := pl.Hand.Len() - pl.MaxSdi; i > 0 {
		for k := 0; k != i; k++ {
			ca := pl.SelectFor(pl.Hand)
			if ca == nil {
				ca = pl.Hand.EndPop()
			}
			ca.Dispatch(Discard)
		}
	}
	pl.Msg("对方回合", nil)
	return true
}

func (pl *Player) init() {
	pl.ActionDraw(5)
	pl.isChain = true
}

func (pl *Player) initDeck(a []uint, b []uint) {
	if pl.Deck.Len() > 0 {
		return
	}
	pl.Game().CardVer.Deck(pl.Deck, pl, a)
	pl.Game().CardVer.Deck(pl.Extra, pl, b)
	pl.ActionShuffle()
}

func (pl *Player) ChangeHp(i int) {
	pl.Dispatch(HPChange, pl, i)

	if i < 0 {
		pl.MsgPub("{self}受到{num}基本分伤害！", Arg{"num": -i})
	} else if i > 0 {
		pl.MsgPub("{self}受到{num}基本分回复！", Arg{"num": i})
	}
	pl.Hp += i
	pl.CallAll(setFace(map[string]interface{}{pl.Name: pl.Hp}))
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
	pl.Dispatch(Draw, pl)
	for i := 0; i != s; i++ {
		if pl.Deck.Len() == 0 {
			return
		}
		pl.Dispatch(DrawNum, pl)
		t := pl.Deck.EndPop()
		pl.Hand.EndPush(t)
	}
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

	for {
		pl.CallAll(flashPhases(pl))
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

func (pl *Player) Select() (t *Card) {
	//pl.ClearCode()
	for {
		pl.CallAll(flashPhases(pl))
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

	pl.Msg("从下列卡牌选择", nil)
	if c := pl.Select(); c != nil {
		for _, v := range *ca {
			if v == c {
				return c
			}
		}
	}
	return nil
}

func (pl *Player) SelectFor(cp ...*Group) *Card {
	s := []string{}
	for _, v := range cp {
		s = append(s, v.GetOwner().Name+" "+string(v.GetName()))
	}
	pl.Msg("从下列区域选择卡牌 {c1}", Arg{"c1": s})
	if c := pl.Select(); c != nil {
		for _, v := range cp {
			if v.IsExistCard(c) {
				return c
			}
		}
	}
	return nil
}
