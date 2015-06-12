package ygo

import (
	"fmt"
	//"rego"
	"rego/agent"
	"rego/misc"
	"service/proto"
	"time"
)

type YGO struct {
	CardVer  *CardVersion
	Room     *misc.Rooms
	StartAt  time.Time
	Cards    map[uint]*Card
	Players  map[uint]*Player
	Survival map[int]int
	Over     bool
}

func NewYGO(r *misc.Rooms) *YGO {
	yg := &YGO{
		Room:     r,
		Cards:    make(map[uint]*Card),
		Survival: make(map[int]int),
		StartAt:  time.Now(),
		Players:  make(map[uint]*Player),
	}
	yg.Room.ForEach(func(sess *agent.Session) {
		p := NewPlayer()
		p.game = yg
		p.Session = sess
		yg.Players[sess.ToUint()] = p
	})
	return yg
}

func (yg *YGO) GetPlayer(sess *agent.Session) *Player {
	return yg.Players[sess.ToUint()]
}

func (yg *YGO) GetCard(u uint) *Card {
	return yg.Cards[u]
}

func (yg *YGO) RegisterCards(c *Card) {
	yg.Cards[c.ToUint()] = c
}

func (yg *YGO) ForEachPlayer(fun func(*Player)) {
	for _, v := range yg.Players {
		fun(v)
	}
}

func (yg *YGO) CallAll(method string, reply interface{}) error {
	yg.Room.BroadcastPush(Call{
		Method: method,
		Args:   reply,
	}, func(sess *agent.Session) {
		yg.Over = true
	})
	return nil
}

func (yg *YGO) Loop() {

	time.Sleep(time.Second) // 初始化
	round := []uint{}
	for k, v := range yg.Players {
		ca := v.Camp
		yg.Survival[ca] = yg.Survival[ca] + 1
		v.Index = len(round)
		v.game = yg
		v.Name = fmt.Sprintf("player %d", len(round))
		round = append(round, k)
	}

	time.Sleep(time.Second) // 客户端初始化
	gi := proto.GameInitResponse{}
	for _, v := range round {
		pi := proto.PlayerInit{
			Hp:   yg.Players[v].Hp,
			Name: yg.Players[v].Name,
		}
		gi.Users = append(gi.Users, pi)
	}
	for _, v := range round {
		gi.Index = yg.Players[v].Index
		yg.Players[v].Call("init", gi)
	}

	time.Sleep(time.Second) // 牌组初始化
	for _, v := range round {
		var s struct {
			Deck proto.Deck `json:"deck"`
		}
		yg.Players[v].Session.Data.DeJson(&s)
		//rego.INFO(string(yg.Players[v].Session.Data.Bytes()))
		yg.Players[v].initDeck(s.Deck.GetMain(), s.Deck.GetExtra())
	}

	time.Sleep(time.Second) // 手牌初始化
	for _, v := range round {
		yg.Players[v].init()
	}

	time.Sleep(time.Second) // 循环
	for {
		for _, v := range round {
			time.Sleep(time.Second)
			if !yg.Players[v].IsFail() {
				yg.Players[v].round()
				//				if yg.CheckWinner(); yg.Over {
				//					return
				//				}
			} else {
				yg.Over = true
			}
			if yg.Over {
				yg.CallAll("over", nil)
				yg.CallAll(Message("游戏结束"))
				return
			}
		}
	}
}
