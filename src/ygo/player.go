package ygo

import (
	"errors"
	"fmt"
	"rego"
	"rego/agent"
	"time"
)

type Player struct {
	//
	Name    string         //用户名
	Session *agent.Session //会话
	Selec   chan uint

	// 规则属性
	Index    int           // 玩家索引
	Game     *YGO          // 属于游戏
	OverTime time.Duration // 允许超出的时间
	WaitTime time.Duration // 每次动作等待的时间

	// 基础属性
	Hp        int  // 生命值
	Camp      int  // 阵营
	RoundSize uint // 回合数
	DrawSize  uint // 抽卡数
	MaxHp     uint // 最大生命值
	MaxSdi    int  // 最大手牌

	// 卡牌区
	Deck    *CardPile //卡组 40 ~ 60
	Hand    *CardPile //手牌
	Extra   *CardPile //额外卡组 <= 15 融合怪物 同调怪物 超量怪物
	Side    *CardPile //副卡组 <= 15
	Removed *CardPile //排除卡
	Grave   *CardPile //墓地
	Mzone   *CardTile //怪物卡区 5
	Szone   *CardTile //魔法卡陷阱卡区 5
	Field   *CardTile //场地卡 5

	// 卡牌事件
	ToExclude  *Events // 排除场外
	ToCemetery *Events // 移动到墓地
	ToSdi      *Events // 返回手牌
	ToDeck     *Events // 返回卡组
	Sustains   *Events // 永续效果

	// 怪兽卡事件
	MonsterInitiative    *Events // 怪兽卡发动效果
	MonsterFreedom       *Events // 解放 送去墓地
	MonsterDestroy       *Events // 破坏 送去墓地
	MonsterFlip          *Events // 反转
	MonsterSummon        *Events // 召唤
	MonsterSummonCover   *Events // 覆盖召唤
	MonsterSummonFlip    *Events // 反转召唤
	MonsterSummonSpecial *Events // 特殊召唤

	// 魔法卡陷阱卡 事件
	MagicAndTrapInitiative *Events // 魔法卡陷阱卡发动效果
	MagicAndTrapCover      *Events // 魔法卡陷阱卡覆盖

	// 是否失败
	fail bool
}

func NewPlayer() *Player {
	player := &Player{
		Camp:     1,
		Hp:       4000,
		DrawSize: 1,
		MaxHp:    ^uint(0),
		MaxSdi:   6,
		OverTime: time.Second * 120,
		WaitTime: time.Second * 5,
		Deck:     NewCardPile(),
		Hand:     NewCardPile(),
		Extra:    NewCardPile(),
		Side:     NewCardPile(),
		Removed:  NewCardPile(),
		Grave:    NewCardPile(),
		Mzone:    NewCardTile(5),
		Szone:    NewCardTile(5),
		Field:    NewCardTile(1),
		Selec:    make(chan uint, 128),
	}
	return player
}
func (pl *Player) SelecAdd(u uint) {
	select {
	case pl.Selec <- u:
	default:
		return
	}
}
func (pl *Player) SelecClear() {
	for {
		select {
		case <-pl.Selec:
		default:
			return
		}
	}
}
func (pl *Player) Fail() {
	pl.fail = true
}

func (pl *Player) IsFail() bool {
	return pl.fail
}

func (pl *Player) ForEachPlayer(fun func(p *Player)) {
	pl.Game.ForEachPlayer(fun)
}

func (pl *Player) round() (err error) {
	defer func() {
		if x := recover(); err != nil {
			rego.ERR(x)
			err = errors.New(fmt.Sprintln(x))
		}
	}()
	pl.RoundSize++
	pl.CallAll("flagName", map[string]interface{}{
		"round":  pl.RoundSize,
		"player": pl.Index,
	})
	pl.draw()
	pl.standby()
	pl.main1()
	pl.battle()
	pl.main2()
	pl.end()
	return
}

func (pl *Player) draw() {
	pl.CallAll("flagStep", map[string]interface{}{
		"step": 1,
	})
	pl.Call(Message("抽牌阶段"))
	if pl.Deck.Len() == 0 {
		pl.Fail()
		return
	}
	pl.ActionDraw(2)
	select {
	case <-time.After(time.Second):
		return
	}
}

func (pl *Player) standby() {
	pl.CallAll("flagStep", map[string]interface{}{
		"step": 2,
	})
	pl.Call(Message("准备阶段"))
	select {
	case <-time.After(time.Second):
		return
	}
}

func (pl *Player) main1() {
	//	pl.CallAll("moveCard", map[string]interface{}{
	//			"uniq": t.ToUint(),
	//			"pos":  "hand",
	//		})
	pl.CallAll("flagStep", map[string]interface{}{
		"step": 3,
		"wait": pl.WaitTime,
	})
	pl.Call(Message("主要阶段1"))
	select {
	case <-time.After(pl.WaitTime):
		return
	}

}

func (pl *Player) battle() {
	pl.CallAll("flagStep", map[string]interface{}{
		"step": 4,
		"wait": pl.WaitTime,
	})
	pl.Call(Message("战斗阶段"))
	select {
	case <-time.After(pl.WaitTime):
		return
	}

}

func (pl *Player) main2() {
	pl.CallAll("flagStep", map[string]interface{}{
		"step": 5,
		"wait": pl.WaitTime,
	})
	pl.Call(Message("主要阶段2"))
	select {
	case <-time.After(pl.WaitTime):
		return
	}
}

func (pl *Player) end() {
	if i := pl.Hand.Len() - pl.MaxSdi; i > 0 {

		pl.Call(Message("选择丢弃的手牌"))
		pl.SelecClear()
	loop:
		for {
			pl.CallAll("flagStep", map[string]interface{}{
				"step": 6,
				"wait": pl.WaitTime,
			})
			pl.Call("selectCard", map[string]interface{}{
				"size": i,
				"pos":  "hand",
			})
			var t *Card
			select {
			case <-time.After(pl.WaitTime):
				t = pl.Hand.EndPop()
			case p := <-pl.Selec:
				t = pl.Hand.PickedForUniq(p)
			}
			if t != nil {
				pl.Grave.EndPush(t)
				pl.CallAll(MoveCard(t, "grave"))
				pl.CallAll(SetFront(t))
				pl.CallAll(ExprCard(t, LE_FaceUp))
				if i--; i == 0 {
					break loop
				}
			}
		}
		pl.Call("finishSelect", map[string]interface{}{})
	}
	pl.CallAll("flagStep", map[string]interface{}{
		"step": 6,
	})
	select {
	case <-time.After(time.Second):

	}
	pl.Call(Message("对方回合"))
}

func (pl *Player) init() {
	pl.ActionShuffle()
	pl.ActionDraw(pl.MaxSdi - 1)
}

func (pl *Player) InitDeck(a []uint) {
	pl.Deck = pl.Game.CardVer.Deck(pl, a)
	pl.Deck.ForEach(func(c *Card) {
		pl.Game.RegisterCards(c)
	})
	pl.ActionShuffle()
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
			return
		}
		t := pl.Deck.BeginPop()
		pl.Hand.EndPush(t)
		pl.CallAll(MoveCard(t, "hand"))
		pl.Call(SetFront(t))
		pl.Call(ExprCard(t, LE_FaceUp))
	}
}

type Call struct {
	Method string      `json:"method"`
	Args   interface{} `json:"args"`
}

func (pl *Player) Call(method string, reply interface{}) error {
	return pl.Game.Room.Push(Call{
		Method: method,
		Args:   reply,
	}, pl.Session)
}

func (pl *Player) CallAll(method string, reply interface{}) error {
	return pl.Game.CallAll(method, reply)
}

func ExprCard(t *Card, le LE_TYPE) (string, interface{}) {
	return "exprCard", map[string]interface{}{
		"uniq": t.ToUint(),
		"expr": le,
	}
}

func SetFront(t *Card) (string, interface{}) {
	return "setFront", map[string]interface{}{
		"desk": t.Id,
		"uniq": t.ToUint(),
	}
}

func Message(msg string) (string, interface{}) {
	return "message", map[string]interface{}{
		"message": msg,
	}
}

func MoveCard(t *Card, pos string) (string, interface{}) {
	return "moveCard", map[string]interface{}{
		"uniq": t.ToUint(),
		"pos":  pos,
	}
}
