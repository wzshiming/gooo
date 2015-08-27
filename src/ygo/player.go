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
	css := cs.Find(func(c *Card) bool {
		return c != ca && c.GetSummoner() == pl
	})

	pl.Call(trigg(css))
	if wi := pl.SelectWill(); wi.Uniq != 0 {
		if c := css.ExistForUniq(wi.Uniq); c != nil {
			e := func() {
				pl.MsgPub("{self}连锁{event}", Arg{"self": c.ToUint(), "event": eventName})
				c.Dispatch(Trigger, a...)
			}
			if ca.Priority() > c.Priority() {
				ca.OnlyOnce(eventName, e, c)
			} else {
				e()
			}
		} else {
			pl.MsgPub("{self}错误的连锁", nil)
		}
	} else {
		pl.MsgPub("{self}不连锁", nil)
	}
	pl.Call(trigg(nil))
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
		ca, u := pl.selectForWarn(pl.Hand, pl.Mzone, pl.Szone, func(c *Card) bool {
			if c.IsInSzone() && c.IsTrap() {
				return false
			} else if c.IsInMzone() && !c.IsCanChange() {
				return false
			}
			return true
		})
		if ca == nil {
			if u == uint(LP_Battle) && lp == LP_Main1 {
				break
			} else if u == uint(LP_End) {
				if lp == LP_Main1 {
					pl.StopOnce(MP)
					pl.StopOnce(BP)
				}
				break
			}
			if u == 0 {
				break
			}
			continue
		}

		if ca.IsInHand() {
			if u == uint(LI_Use1) {
				ca.Dispatch(Use1, ca)
			} else if u == uint(LI_Use2) {
				ca.Dispatch(Use2, ca)
			}
		} else if ca.IsInMzone() {
			ca.Dispatch(Expression)
		} else if ca.IsInSzone() {
			ca.Dispatch(Onset)
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
		ca, u := pl.selectForWarn(pl.Mzone, func(c *Card) bool {
			return c.IsCanAttack()
		})
		if ca == nil {
			if u == uint(LP_Main2) {
				break
			} else if u == uint(LP_End) {
				pl.StopOnce(MP)
				break
			}
			if u == 0 {
				break
			}
			continue
		}
		if !ca.IsInMzone() || !ca.IsFaceUpAttack() {
			pl.Msg("请选择怪兽区正面攻击表示的怪兽", nil)
			continue
		}
		if !ca.IsCanAttack() {
			pl.Msg("当前怪兽不能攻击", nil)
			continue
		}
		tar := pl.GetTarget()
		if tar.Mzone.Len() != 0 {
			pl.Msg("选择要攻击的目标", nil)
			if c, _ := pl.selectForWarn(tar.Mzone); c.IsInMzone() {
				ca.Dispatch(Declaration, c)
			} else {
				pl.Msg("请选择对方怪兽区的怪兽", nil)
			}
		} else {
			ca.Dispatch(Declaration)
		}
	}

}

func (pl *Player) end(lp lp_type) {
	pl.Phases = lp
	if i := pl.Hand.Len() - pl.MaxSdi; i > 0 {
		pl.ResetReplyTime()
		pl.Msg("请{self}选择丢弃的手牌", nil)
		for k := 0; k != i; k++ {
			ca := pl.SelectForWarn(pl.Hand)
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

func (pl *Player) IsCanSummon() bool {
	return pl.lastSummonRound < pl.GetRound()
}
func (pl *Player) SetCanSummon() {
	pl.lastSummonRound = 0
}

func (pl *Player) SetNotCanSummon() {
	pl.lastSummonRound = pl.GetRound()
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

func (pl *Player) Select() (*Card, uint) {
	p := pl.SelectWill()
	if p.Uniq != 0 {
		return pl.Game().GetCard(p.Uniq), p.Method
	}
	return nil, p.Method
}
func (pl *Player) selectForPopup(ci ...interface{}) (c *Card, u uint) {
	css := NewCards(ci...)
	if css.Len() == 0 {
		return
	}
	pl.Call(setPick(css, pl))
	defer pl.Call(setPick(nil, pl))
	if c, u = pl.Select(); c != nil {
		if css.IsExistCard(c) {
			return
		}
	}
	return
}
func (pl *Player) SelectForPopup(ci ...interface{}) *Card {
	c, _ := pl.selectForPopup(ci...)
	return c
}

func (pl *Player) selectForWarn(ci ...interface{}) (c *Card, u uint) {
	css := NewCards(ci...)
	if css.Len() == 0 {
		return
	}
	pl.Call(trigg(css))
	defer func() {
		pl.Call(trigg(nil))
	}()
	if c, u = pl.Select(); c != nil {
		if css.IsExistCard(c) {
			return
		}
	}
	return
}

func (pl *Player) SelectForWarn(ci ...interface{}) *Card {
	c, _ := pl.selectForWarn(ci...)
	return c
}
