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

	pending map[uint]*Card
	// 是否失败
	fail bool
}

func NewPlayer() *Player {
	pl := &Player{
		Events:    dispatcher.NewLineEvent(),
		Camp:      1,
		Hp:        4000,
		DrawSize:  1,
		MaxHp:     ^uint(0),
		MaxSdi:    6,
		OverTime:  time.Second * 120,
		WaitTime:  time.Second * 60,
		ReplyTime: time.Second * 20,
		pending:   make(map[uint]*Card),
	}
	var pr uint
	pl.MsgChan = NewMsgChan(func(m MsgCode) bool {
		//rego.INFO(m)
		if m.Uniq == 1 {
			return true
		}
		if m.Uniq != 0 {
			ca := pl.game.GetCard(m.Uniq)
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

	pl.AddEvent(Chain, func() {
		pl.MsgPub("等待{self}连锁", nil)
	})

	pl.AddEvent(DP, pl.draw)
	pl.AddEvent(SP, pl.standby)
	pl.AddEvent(MP, pl.main)
	pl.AddEvent(BP, pl.battle)
	pl.AddEvent(EP, pl.end)

	pl.AddEvent(DetectionChain, func() {
		if len(pl.pending) == 0 {
			return
		}
		pl.Dispatch(Chain)
		pl.pending = map[uint]*Card{}
	})
	pl.AddEvent(Chain, pl.chain)
	return pl
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

func (pl *Player) AddReply(ca *Card) {
	pl.pending[ca.ToUint()] = ca
}

func (pl *Player) Fail() {
	pl.fail = true
}

func (pl *Player) IsFail() bool {
	return pl.fail
}

func (pl *Player) ForEachPlayer(fun func(p *Player)) {
	pl.game.ForEachPlayer(fun)
}

func (pl *Player) Dispatch(eventName string, args ...interface{}) {
	pl.Events.Dispatch(eventName, args...)
	e := func(c *Card) bool {
		c.Dispatch(eventName, args...)
		return true
	}
	pl.Mzone.ForEach(e)
	pl.Szone.ForEach(e)
	pl.Field.ForEach(e)
	pl.Hand.ForEach(e)
	if eventName != DetectionChain {
		pl.Dispatch(DetectionChain, append(args, eventName)...)
	}
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

// 连锁
func (pl *Player) chain() {
	pl.ResetReplyTime()
	if wi := pl.SelectWill(); wi.Method != 0 {
		if ca := pl.pending[wi.Uniq]; ca != nil {
			ca.Dispatch(Trigger, wi.Method)
		}
	}
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
				t2 := pl.game.GetCard(j.Uniq)
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
	if i := pl.Hand.Len() - pl.MaxSdi; i > 0 {
		pl.Msg("选择丢弃的手牌", nil)
		pl.SelectFunc(i, pl.Hand, ToGrave)
	}
	pl.Msg("对方回合", nil)
	return true
}

func (pl *Player) init() {
	pl.ActionDraw(5)
}

func (pl *Player) initDeck(a []uint, b []uint) {
	if pl.Deck.Len() > 0 {
		return
	}
	pl.game.CardVer.Deck(pl.Deck, pl, a)
	pl.game.CardVer.Deck(pl.Extra, pl, b)
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
		return pl.game.GetPlayerForIndex(1)
	}
	return pl.game.GetPlayerForIndex(0)
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
		t := pl.Deck.BeginPop()
		pl.Hand.EndPush(t)
	}
}

func (pl *Player) Call(method string, reply interface{}) error {
	return pl.game.Room.Push(Call{
		Method: method,
		Args:   reply,
	}, pl.Session)
}

func (pl *Player) CallAll(method string, reply interface{}) error {
	return pl.game.CallAll(method, reply)
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
			t = pl.game.GetCard(p.Uniq)
			return
		}
	}
	return
}

func (pl *Player) SelectMust(cp *Group) (t *Card) {
	//pl.ClearCode()
	for {
		pl.CallAll(flashPhases(pl))
		select {
		case <-time.After(time.Second):
			pl.PassTime -= time.Second
			if pl.PassTime <= 0 {
				t = cp.EndPop()
				return
			}
		case p := <-pl.GetCode():
			if t = cp.ExistForUniq(p.Uniq); t == nil {
				t = cp.EndPop()
			}
			return
		}
	}

	return
}

func (pl *Player) SelectSzone() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == LL_Szone {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectSzoneTarget() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == LL_Szone && pla.GetOwner() != pl {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectSzoneSelf() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == LL_Szone && pla.GetOwner() == pl {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectMzone() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == LL_Mzone {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectMzoneTarget() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == LL_Mzone && pla.GetOwner() != pl {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectMzoneSelf() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == LL_Mzone && pla.GetOwner() == pl {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectGrave() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == LL_Grave {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectGraveTarget() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == LL_Grave && pla.GetOwner() != pl {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectGraveSelf() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == LL_Grave && pla.GetOwner() == pl {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectHand() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == LL_Hand {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectHandTarget() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == LL_Hand && pla.GetOwner() != pl {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectHandSelf() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == LL_Hand && pla.GetOwner() == pl {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectFunc(size int, cp *Group, fun func(c *Card) bool) {
	for i := 0; i != size; i++ {
		t := pl.SelectMust(cp)
		fun(t)
	}
	return
}
