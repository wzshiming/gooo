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
	MaxSdi    uint // 最大手牌

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
	}
	return player
}

func (pl *Player) Fail() {
	pl.Game.AddLoser(pl.Camp)
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
	pl.Game.CallAll("flagName", map[string]interface{}{
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
	pl.Game.CallAll("flagStep", map[string]interface{}{
		"step": 1,
	})
	if pl.Deck.Len() == 0 {
		pl.Fail()
		return
	}
	pl.ActionDraw(1)
	<-time.After(time.Second)
}

func (pl *Player) standby() {
	pl.Game.CallAll("flagStep", map[string]interface{}{
		"step": 2,
	})
	<-time.After(time.Second)
}

func (pl *Player) main1() {
	//	pl.Game.CallAll("moveCard", map[string]interface{}{
	//			"uniq": t.ToUint(),
	//			"pos":  "hand",
	//		})
	pl.Game.CallAll("flagStep", map[string]interface{}{
		"step": 3,
		"wait": pl.WaitTime,
	})
	for {
		select {
		case <-time.After(pl.WaitTime):
			return
		}
	}
}

func (pl *Player) battle() {
	pl.Game.CallAll("flagStep", map[string]interface{}{
		"step": 4,
		"wait": pl.WaitTime,
	})
	for {
		select {
		case <-time.After(pl.WaitTime):
			return
		}
	}
}

func (pl *Player) main2() {
	pl.Game.CallAll("flagStep", map[string]interface{}{
		"step": 5,
		"wait": pl.WaitTime,
	})
	for {
		select {
		case <-time.After(pl.WaitTime):
			return
		}
	}
}

func (pl *Player) end() {
	pl.Game.CallAll("flagStep", map[string]interface{}{
		"step": 6,
	})
	<-time.After(time.Second)
}

func (pl *Player) init() {
	pl.ActionDraw(pl.MaxSdi - 1)
}

func (pl *Player) InitDeck(a []uint) {
	pl.Deck = pl.Game.CardVer.Deck(pl, a)
	pl.Deck.ForEach(func(c *Card) {
		pl.Game.RegisterCards(c)
	})
}

func (pl *Player) ActionDraw(s uint) {
	for i := uint(0); i != s; i++ {
		if pl.Deck.Len() == 0 {
			return
		}
		t := pl.Deck.BeginPop()
		pl.Hand.EndPush(t)
		pl.Game.CallAll("moveCard", map[string]interface{}{
			"uniq": t.ToUint(),
			"pos":  "hand",
		})
		pl.Call("setFront", map[string]interface{}{
			"desk": t.Id,
			"uniq": t.ToUint(),
		})
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
