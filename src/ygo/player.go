package ygo

import (
	"github.com/wzshiming/dispatcher"
	"github.com/wzshiming/rego"
	"github.com/wzshiming/rego/agent"
	"time"
)

type Player struct {
	//
	dispatcher.Events
	MsgChan
	Name    string         // 用户名
	Session *agent.Session // 会话

	// 规则属性
	Index    int           // 玩家索引
	game     *YGO          // 属于游戏
	OverTime time.Duration // 允许超出的时间
	WaitTime time.Duration // 每次动作等待的时间

	// 基础属性
	Hp         int     // 生命值
	Camp       int     // 阵营
	RoundSize  int     // 回合数
	DrawSize   uint    // 抽卡数
	MaxHp      uint    // 最大生命值
	MaxSdi     int     // 最大手牌
	Phases     LP_TYPE // 阶段
	PhasesNext LP_TYPE // 下一阶段

	// 卡牌区
	Deck    *Cards // 卡组 40 ~ 60
	Hand    *Cards // 手牌
	Extra   *Cards // 额外卡组 <= 15 融合怪物 同调怪物 超量怪物
	Side    *Cards // 副卡组 <= 15
	Removed *Cards // 排除卡
	Grave   *Cards // 墓地
	Mzone   *Cards // 怪物卡区 5
	Szone   *Cards // 魔法卡陷阱卡区 5
	Field   *Cards // 场地卡

	// 卡牌事件
	ToRemoved *Effects // 排除场外
	ToGrave   *Effects // 移动到墓地
	ToHand    *Effects // 返回手牌
	ToDeck    *Effects // 返回卡组
	Sustains  *Effects // 永续效果

	// 怪兽卡事件
	MonsterInitiative    *Effects // 怪兽卡发动效果
	MonsterFreedom       *Effects // 解放 送去墓地
	MonsterDestroy       *Effects // 破坏 送去墓地
	MonsterFlip          *Effects // 反转
	MonsterSummon        *Effects // 召唤
	MonsterSummonCover   *Effects // 覆盖召唤
	MonsterSummonFlip    *Effects // 反转召唤
	MonsterSummonSpecial *Effects // 特殊召唤

	EventDrawSuf    *Effects // 抽牌之后
	EventStandbyPre *Effects // 准备阶段之前
	EventStandbySuf *Effects // 准备阶段之后
	EventMainPre    *Effects // 准备阶段之前
	EventMainSuf    *Effects // 准备阶段之后

	// 魔法卡陷阱卡 事件
	MagicAndTrapInitiative *Effects // 魔法卡陷阱卡发动效果
	MagicAndTrapCover      *Effects // 魔法卡陷阱卡覆盖

	// 是否失败
	fail bool
}

func NewPlayer() *Player {
	player := &Player{
		Events:   dispatcher.NewLineEvent(),
		Camp:     1,
		Hp:       4000,
		DrawSize: 1,
		MaxHp:    ^uint(0),
		MaxSdi:   6,
		OverTime: time.Second * 120,
		WaitTime: time.Second * 10,

		MsgChan: NewMsgChan(),
	}
	player.Deck = NewCards(player, Deck)
	player.Hand = NewCards(player, Hand)
	player.Extra = NewCards(player, Extra)
	player.Removed = NewCards(player, Removed)
	player.Grave = NewCards(player, Grave)
	player.Mzone = NewCards(player, Mzone)
	player.Szone = NewCards(player, Szone)
	player.Field = NewCards(player, Field)

	player.Deck.SetJoin(func(c *Card) {
		c.FaceDownAttack()
	})
	player.Extra.SetJoin(func(c *Card) {
		c.FaceUpAttack()
	})
	player.Hand.SetJoin(func(c *Card) {
		player.Call(SetFront(c))
		player.Call(ExprCard(c, LE_FaceUpAttack))
	})
	player.Grave.SetJoin(func(c *Card) {
		c.FaceUpAttack()
	})
	return player
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

func (pl *Player) GetRound() int {
	return pl.RoundSize
}

func (pl *Player) round() (err error) {
	pl.CallAll("flagName", map[string]interface{}{
		"round":  pl.GetRound(),
		"player": pl.Index,
	})
	pl.Dispatch(DPPre, pl)
	pl.draw(LP_Draw)
	pl.Dispatch(DPSuf, pl)

	pl.Dispatch(SPPre, pl)
	pl.standby(LP_Standby)
	pl.Dispatch(SPSuf, pl)

	pl.Dispatch(MP1Pre, pl)
	if pl.main(LP_Main1) {
		pl.ToBattle()
	}
	pl.Dispatch(MP1Suf, pl)

	pl.Dispatch(BPPre, pl)
	if pl.PhasesNext == LP_Battle {
		if pl.battle(LP_Battle) {
			pl.ToMain()
		}
	}
	pl.Dispatch(BPSuf, pl)

	pl.Dispatch(MP2Pre, pl)
	if pl.PhasesNext == LP_Main2 {
		pl.main(LP_Main2)
		pl.ToEnd()
	}
	pl.Dispatch(MP2Suf, pl)

	pl.Dispatch(EPPre, pl)
	pl.end(LP_End)
	pl.Dispatch(EPSuf, pl)

	pl.RoundSize++
	return
}

func (pl *Player) ToBattle() {
	pl.PhasesNext = LP_Battle
}

func (pl *Player) ToMain() {
	pl.PhasesNext = LP_Main2
}

func (pl *Player) ToEnd() {
	pl.PhasesNext = LP_End
}

func (pl *Player) draw(lp LP_TYPE) bool {
	defer func() {
		if x := recover(); x != nil {
			rego.DebugStack()
		}
	}()
	pl.Phases = lp
	pl.CallAll(flashPhases(pl))
	pl.Call(Message("抽牌阶段"))
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
	pl.Call(Message("准备阶段"))

	return true
}

func (pl *Player) main(lp LP_TYPE) bool {
	defer func() {
		if x := recover(); x != nil {
			rego.DebugStack()
		}
	}()
	pl.Phases = lp
	pl.Call(Message("主要阶段"))
	pl.ClearCode()

	for {
		p := pl.SelectWill()
		if p.Uniq == 0 {
			break
		}
		if t := pl.Hand.ExistForUniq(p.Uniq); t != nil {
			if p.Method == 0 {
				pl.Call("optionCard", map[string]interface{}{
					"uniq": p.Uniq,
					"list": t.HandMethods(),
				})
			} else if p.Method == uint(LI_Call) {
				t.Call()
				pl.Call("finishSelect", map[string]interface{}{})
			} else if p.Method == uint(LI_Cover) {
				t.Cover()
				pl.Call("finishSelect", map[string]interface{}{})
			} else if p.Method == uint(LI_Use) {
				t.Use()
				pl.Call("finishSelect", map[string]interface{}{})
			}
		} else if t := pl.Mzone.ExistForUniq(p.Uniq); t != nil {
			t.ChangeExpression()
		} else if t := pl.Szone.ExistForUniq(p.Uniq); t != nil {
			t.EffectOnset()
		} else {
			pl.Call(Message("非法目标"))
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
	pl.Call(Message("战斗阶段"))
	pl.ClearCode()
	for {
		if i := pl.SelectWill(); i.Uniq != 0 {
			t1 := pl.Mzone.ExistForUniq(i.Uniq)
			if t1 == nil || !t1.IsFaceUpAttack() {
				pl.Call(Message("请选择怪兽区正面攻击表示的怪兽"))
				continue
			}

			if !t1.IsCanAttack() {
				pl.Call(Message("当前怪兽不能攻击"))
				continue
			}

			pl.Call(Message("选择要攻击的目标"))
			j := pl.SelectWill()
			if j.Uniq != 0 {
				t2 := pl.game.GetCard(j.Uniq)
				if pl.GetTarget().Mzone.IsExistCard(t2) {
					t1.Battle(t2)
				} else {
					pl.Call(Message("请选择对方怪兽区的怪兽"))
				}
			}
			//pl.Call(Message("战斗阶段"))
		} else {
			break
		}
	}
	//	select {
	//	case <-time.After(time.Second / 10):
	//	}
	return true
}

func (pl *Player) end(lp LP_TYPE) bool {
	defer func() {
		if x := recover(); x != nil {
			rego.DebugStack()
		}
	}()
	pl.Phases = lp
	if i := pl.Hand.Len() - pl.MaxSdi; i > 0 {
		pl.Call(Message("选择丢弃的手牌"))
		pl.SelectFunc(i, pl.Hand, ToGrave)
	}
	pl.Call(Message("对方回合"))
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

func (pl *Player) ChangeHp(i int, ca *Card) {
	if ca != nil {
		ca.Dispatch(Deduct, pl)
	}
	pl.Dispatch(HPChange, pl, i)
	pl.Hp += i
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

func (pl *Player) SelectWill() (p MsgCode) {
	pl.ClearCode()
	pl.CallAll(flashPhases(pl))
	select {
	case <-time.After(pl.WaitTime):
	case p = <-pl.GetCode():
	}
	return
}

func (pl *Player) Select() (t *Card) {
	pl.ClearCode()
	pl.CallAll(flashPhases(pl))
	select {
	case <-time.After(pl.WaitTime):
	case p := <-pl.GetCode():
		t = pl.game.GetCard(p.Uniq)
	}
	return
}

func (pl *Player) SelectSzone() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == Szone {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectSzoneTarget() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == Szone && pla.GetOwner() != pl {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectSzoneSelf() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == Szone && pla.GetOwner() == pl {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectMzone() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == Mzone {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectMzoneTarget() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == Mzone && pla.GetOwner() != pl {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectMzoneSelf() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == Mzone && pla.GetOwner() == pl {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectGrave() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == Grave {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectGraveTarget() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == Grave && pla.GetOwner() != pl {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectGraveSelf() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == Grave && pla.GetOwner() == pl {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectHand() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == Hand {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectHandTarget() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == Hand && pla.GetOwner() != pl {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectHandSelf() (t *Card) {
	if tar := pl.Select(); tar != nil {
		pla := tar.GetPlace()
		if pla.GetName() == Hand && pla.GetOwner() == pl {
			return tar
		}
	}
	return nil
}

func (pl *Player) SelectMust(cp *Cards) (t *Card) {
	pl.ClearCode()
	pl.CallAll(flashPhases(pl))
	select {
	case <-time.After(pl.WaitTime):
		t = cp.EndPop()
	case p := <-pl.GetCode():
		if t = cp.ExistForUniq(p.Uniq); t == nil {
			t = cp.EndPop()
		}
	}
	return
}

func (pl *Player) SelectFunc(size int, cp *Cards, fun func(c *Card) bool) {
	for i := 0; i != size; i++ {
		t := pl.SelectMust(cp)
		fun(t)
	}
	return
}
