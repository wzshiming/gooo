package ygo

import (
	"errors"
	"fmt"
	"rego"
	"rego/agent"
	"service/proto"
	"time"
	"ygo/defaul"
)

type Player struct {
	//
	Name    string         //用户名
	Session *agent.Session //会话
	Selec   chan proto.SelectableRequest

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
	Phases    int  //阶段

	// 卡牌区
	Deck    *CardPile //卡组 40 ~ 60
	Hand    *CardPile //手牌
	Extra   *CardPile //额外卡组 <= 15 融合怪物 同调怪物 超量怪物
	Side    *CardPile //副卡组 <= 15
	Removed *CardPile //排除卡
	Grave   *CardPile //墓地
	Mzone   *CardPile //怪物卡区 5
	Szone   *CardPile //魔法卡陷阱卡区 5
	Field   *CardPile //场地卡 5

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
		WaitTime: time.Second * 3,
		Deck:     NewCardPile("deck"),
		Hand:     NewCardPile("hand"),
		Extra:    NewCardPile("extra"),
		Side:     NewCardPile("side"),
		Removed:  NewCardPile("remove"),
		Grave:    NewCardPile("grave"),
		Mzone:    NewCardPile("mzone"),
		Szone:    NewCardPile("szone"),
		Field:    NewCardPile("field"),
		Selec:    make(chan proto.SelectableRequest, 128),
	}

	player.Deck.SetJoin(func(c *Card) {
		player.CallAll(ExprCard(c, LE_FaceDown))
	})
	player.Hand.SetJoin(func(c *Card) {
		player.Call(SetFront(c))
		player.Call(ExprCard(c, LE_FaceUp))
	})
	player.Grave.SetJoin(func(c *Card) {
		player.CallAll(SetFront(c))
		c.SetLE(LE_FaceUp)
	})
	return player
}

func (pl *Player) SelecAdd(u proto.SelectableRequest) {
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
	pl.Phases = 1
	pl.draw()
	pl.Phases = 2
	pl.standby()
	pl.Phases = 3
	pl.main()
	pl.Phases = 4
	pl.battle()
	pl.Phases = 5
	pl.main()
	pl.Phases = 6
	pl.end()
	return
}

func (pl *Player) draw() {
	pl.CallAll("flagStep", map[string]interface{}{
		"step": pl.Phases,
	})
	pl.Call(Message("抽牌阶段"))
	if pl.Deck.Len() == 0 {
		pl.Fail()
		return
	}
	pl.ActionDraw(1)
	select {
	case <-time.After(time.Second):
		return
	}
}

func (pl *Player) standby() {
	pl.CallAll("flagStep", map[string]interface{}{
		"step": pl.Phases,
	})
	pl.Call(Message("准备阶段"))
	select {
	case <-time.After(time.Second):
		return
	}
}

func (pl *Player) main() {

	pl.Call(Message("主要阶段1"))
	pl.Call("useCard", map[string]interface{}{
		"master": pl.Index,
		"pos":    "hand",
	})
loop:
	for {
		pl.SelecClear()
		pl.CallAll("flagStep", map[string]interface{}{
			"step": pl.Phases,
			"wait": pl.WaitTime * 3,
		})
		var t *Card
		select {
		case <-time.After(pl.WaitTime * 3):
			break loop
		case p := <-pl.Selec:
			if t = pl.Hand.ExistForUniq(p.Uniq); t != nil {
				if p.Method == 0 {
					pl.Call("optionCard", map[string]interface{}{
						"uniq": p.Uniq,
						"list": []string{"Call", "Place", "Use", "Discard"},
					})
				} else {
					pl.Mzone.EndPush(t)
				}
			}
		}
	}
	pl.CallAll("flashCards", map[string]interface{}{
		"master": pl.Index,
		"pos":    "hand",
	})
	pl.Call("finishSelect", map[string]interface{}{})
}

func (pl *Player) battle() {
	pl.CallAll("flagStep", map[string]interface{}{
		"step": pl.Phases,
		"wait": pl.WaitTime,
	})
	pl.Call(Message("战斗阶段"))
	select {
	case <-time.After(pl.WaitTime):
		return
	}
}

func (pl *Player) end() {
	if i := pl.Hand.Len() - pl.MaxSdi; i > 0 {
		pl.SelecClear()
		pl.Call(Message("选择丢弃的手牌"))
		pl.Call("selectCard", map[string]interface{}{
			"master": pl.Index,
			"pos":    "hand",
		})
	loop:
		for {
			pl.CallAll("flagStep", map[string]interface{}{
				"step": pl.Phases,
				"wait": pl.WaitTime,
			})
			var t *Card
			select {
			case <-time.After(pl.WaitTime):
				t = pl.Hand.EndPop()
			case p := <-pl.Selec:
				if t = pl.Hand.ExistForUniq(p.Uniq); t == nil {
					t = pl.Hand.EndPop()
				}
			}
			pl.Grave.EndPush(t)

			if i--; i == 0 {
				break loop
			}
		}
		pl.CallAll("flashCards", map[string]interface{}{
			"master": pl.Index,
			"pos":    "hand",
		})
		//pl.Call("finishSelect", map[string]interface{}{})
	}

	pl.CallAll("flagStep", map[string]interface{}{
		"step": pl.Phases,
	})
	select {
	case <-time.After(time.Second):
	}
	pl.Call(Message("对方回合"))
}

func (pl *Player) init() {
	pl.ActionDraw(5)
}

func (pl *Player) initCards() {
	pl.InitDeck(defaul.DefaultDeck)
	pl.ActionShuffle()
}

func (pl *Player) InitDeck(a []uint) {
	if pl.Deck.Len() > 0 {
		return
	}
	pl.Game.CardVer.Deck(pl.Deck, pl, a)
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
	}
}

type Call struct {
	Method string      `json:"method"`
	Args   interface{} `json:"args"`
}

func (pl *Player) Call(method string, reply interface{}) error {
	//time.Sleep(time.Millisecond * 10)
	return pl.Game.Room.Push(Call{
		Method: method,
		Args:   reply,
	}, pl.Session)
}

func (pl *Player) CallAll(method string, reply interface{}) error {
	//time.Sleep(time.Millisecond * 10)
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
		"uniq":   t.ToUint(),
		"master": t.Owner.Index,
		"pos":    pos,
	}
}
