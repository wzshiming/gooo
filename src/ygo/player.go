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
	Phases  LP_TYPE

	// 规则属性
	Index    int           // 玩家索引
	game     *YGO          // 属于游戏
	OverTime time.Duration // 允许超出的时间
	WaitTime time.Duration // 每次动作等待的时间

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

	// 是否失败
	fail bool
}

func NewPlayer() *Player {
	pl := &Player{
		Events:   dispatcher.NewLineEvent(),
		Camp:     1,
		Hp:       4000,
		DrawSize: 1,
		MaxHp:    ^uint(0),
		MaxSdi:   6,
		OverTime: time.Second * 120,
		WaitTime: time.Second * 10,
	}
	var pr uint
	pl.MsgChan = NewMsgChan(func(m MsgCode) bool {
		//rego.INFO(m)
		if m.Uniq != 0 {
			ca := pl.game.GetCard(m.Uniq)
			if ca != nil {
				if m.Method == 101 {
					if pr != 0 {
						pl.GetTarget().CallAll(Touch(pr, 1, 1, 1))
						pr = 0
					}
					//pl.Call(Touch(m.Uniq, -1, -1, 50))
					pl.GetTarget().Call(Touch(m.Uniq, -1, -100, 100))
				} else if m.Method == 102 {
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

	//	pl.Deck.SetJoin(func(c *Card) {
	//		c.FaceDownAttack()
	//	})
	pl.Extra.SetJoin(func(c *Card) {
		c.FaceUpAttack()
	})
	pl.Hand.SetJoin(func(c *Card) {
		pl.Call(SetFront(c))
		pl.Call(ExprCard(c, LE_FaceUpAttack))
	})
	pl.Grave.SetJoin(func(c *Card) {
		c.FaceUpAttack()
	})

	pl.AddEventListener(DP, pl.draw)
	pl.AddEventListener(SP, pl.standby)
	pl.AddEventListener(MP, pl.main)
	pl.AddEventListener(BP, pl.battle)
	pl.AddEventListener(EP, pl.end)
	return pl
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

	pl.Dispatch(DP, LP_Draw)
	pl.Dispatch(SP, LP_Standby)
	pl.Dispatch(MP, LP_Main1)
	pl.Dispatch(BP, LP_Battle)
	pl.Dispatch(MP, LP_Main2)
	pl.Dispatch(EP, LP_End)

	pl.RoundSize++
	return
}

func (pl *Player) draw(lp LP_TYPE) bool {
	defer func() {
		if x := recover(); x != nil {
			rego.DebugStack()
		}
	}()
	pl.Dispatch(DPPre, pl)
	pl.Phases = lp
	pl.CallAll(flashPhases(pl))
	pl.Call(Message("抽牌阶段"))
	if pl.Deck.Len() == 0 {
		pl.Fail()
		return false
	}
	pl.ActionDraw(1)
	pl.Dispatch(DPSuf, pl)
	return true
}

func (pl *Player) standby(lp LP_TYPE) bool {
	defer func() {
		if x := recover(); x != nil {
			rego.DebugStack()
		}
	}()
	pl.Dispatch(SPPre, pl)
	pl.Phases = lp
	pl.CallAll(flashPhases(pl))
	pl.Call(Message("准备阶段"))
	pl.Dispatch(SPSuf, pl)
	return true
}

func (pl *Player) main(lp LP_TYPE) bool {
	defer func() {
		if x := recover(); x != nil {
			rego.DebugStack()
		}
	}()
	pl.Dispatch(MPPre, pl)
	pl.Phases = lp
	pl.Call(Message("主要阶段"))
	//pl.ClearCode()

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
				if t.IsMonster() {
					t.Dispatch(Summon, t)
				} else {
					t.Dispatch(Use, t)
				}

			} else if p.Method == uint(LI_Cover) {
				t.Dispatch(Cover, t)
			}
		} else if t := pl.Mzone.ExistForUniq(p.Uniq); t != nil {
			t.ChangeExpression()
		} else if t := pl.Szone.ExistForUniq(p.Uniq); t != nil {
			t.Dispatch(Onset)
		} else {
			pl.Call(Message("非法目标"))
		}
	}
	pl.Dispatch(MPSuf, pl)
	return true
}

func (pl *Player) battle(lp LP_TYPE) bool {
	defer func() {
		if x := recover(); x != nil {
			rego.DebugStack()
		}
	}()
	pl.Dispatch(BPPre, pl)
	pl.Phases = lp
	pl.Call(Message("战斗阶段"))
	//pl.ClearCode()
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
			if j := pl.SelectWill(); j.Uniq != 0 {
				t2 := pl.game.GetCard(j.Uniq)
				if pl.GetTarget().Mzone.IsExistCard(t2) {
					t1.Dispatch(Battle, t2)
				} else {
					pl.Call(Message("请选择对方怪兽区的怪兽"))
				}
			}
			//pl.Call(Message("战斗阶段"))
		} else {
			break
		}
	}
	pl.Dispatch(BPSuf, pl)
	return true
}

func (pl *Player) end(lp LP_TYPE) bool {
	defer func() {
		if x := recover(); x != nil {
			rego.DebugStack()
		}
	}()
	pl.Dispatch(EPPre, pl)
	pl.Phases = lp
	if i := pl.Hand.Len() - pl.MaxSdi; i > 0 {
		pl.Call(Message("选择丢弃的手牌"))
		pl.SelectFunc(i, pl.Hand, ToGrave)
	}
	pl.Call(Message("对方回合"))
	pl.Dispatch(EPSuf, pl)
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
	//pl.ClearCode()
	pl.CallAll(flashPhases(pl))
	select {
	case <-time.After(pl.WaitTime):
	case p = <-pl.GetCode():
	}
	return
}

func (pl *Player) Select() (t *Card) {
	//pl.ClearCode()
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

func (pl *Player) SelectMust(cp *Group) (t *Card) {
	//pl.ClearCode()
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

func (pl *Player) SelectFunc(size int, cp *Group, fun func(c *Card) bool) {
	for i := 0; i != size; i++ {
		t := pl.SelectMust(cp)
		fun(t)
	}
	return
}
