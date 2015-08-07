package ygo

import (
	"fmt"
	"service/proto"
	"time"

	"github.com/wzshiming/dispatcher"
	"github.com/wzshiming/rego/agent"
	"github.com/wzshiming/rego/misc"
)

type Arg map[string]interface{}

type YGO struct {
	dispatcher.Events
	Fork     dispatcher.Fork
	CardVer  *CardVersion
	Room     *misc.Rooms
	StartAt  time.Time
	Cards    map[uint]*Card
	Players  map[uint]*Player
	Current  *Player
	Survival map[int]int
	Over     bool
	round    []uint

	pending   map[uint]*Card
	cardevent map[string]map[*Card]Action
}

func nap(i int) {
	time.Sleep(time.Second * time.Duration(i) / 10)
}

func NewYGO(r *misc.Rooms) *YGO {
	eve := dispatcher.NewLineEvent()
	yg := &YGO{
		Events:   eve,
		Fork:     eve.GetFork(),
		Room:     r,
		Cards:    make(map[uint]*Card),
		Survival: make(map[int]int),
		StartAt:  time.Now(),
		Players:  make(map[uint]*Player),
	}

	yg.Room.ForEach(func(sess *agent.Session) {
		p := NewPlayer(yg)
		p.Session = sess
		yg.Players[sess.ToUint()] = p
	})

	return yg
}

//func (yg *YGO) Dispatch(eventName string, args ...interface{}) {
//	rego.ERR(eventName, args)
//}

func (yg *YGO) Chain(eventName string, ca *Card, pl *Player, args []interface{}) {

	if ca != nil {
		args = append(args, ca)
	}
	if pl != nil {
		args = append(args, pl)
	}

	yg.EmptyEvent(Chain)
	yg.Dispatch(eventName, args...)

	cs := []*Card{}
	yg.ForEventEach(Chain, func(i interface{}) {
		if v, ok := i.(*Card); ok {
			cs = append(cs, v)
		}
	})
	if len(cs) > 0 {
		//pl := ca.GetSummoner()
		pl.MsgPub("连锁事件 "+eventName+" {self}", nil)

		if !pl.Chain(eventName, cs, args) {
			pl.GetTarget().Chain(eventName, cs, args)
		}
	}
	yg.EmptyEvent(Chain)

}

func (yg *YGO) GetPlayer(sess *agent.Session) *Player {
	return yg.Players[sess.ToUint()]
}

func (yg *YGO) GetCard(u uint) (c *Card) {
	c = yg.Cards[u]
	return
}

func (yg *YGO) RegisterCards(c *Card) {
	yg.Cards[c.ToUint()] = c
}

func (yg *YGO) ForEachPlayer(fun func(*Player)) {
	for _, v := range yg.round {
		fun(yg.Players[v])
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

func (yg *YGO) GetPlayerForIndex(i int) *Player {
	return yg.Players[yg.round[i]]
}

func (yg *YGO) Loop() {

	nap(1) // 初始化

	for k, _ := range yg.Players {
		yg.round = append(yg.round, k)
	}

	for k, v := range yg.round {
		ca := yg.Players[v].Camp
		yg.Survival[ca] = yg.Survival[ca] + 1
		yg.Players[v].Index = k
		yg.Players[v].game = yg
		yg.Players[v].RoundSize = 0
		yg.Players[v].Name = fmt.Sprintf("player %d", k)
	}

	nap(1) // 客户端初始化
	gi := proto.GameInitResponse{}
	for _, v := range yg.round {
		pi := proto.PlayerInit{
			Hp:   yg.Players[v].Hp,
			Name: yg.Players[v].Name,
		}
		gi.Users = append(gi.Users, pi)
	}
	for _, v := range yg.round {
		gi.Index = yg.Players[v].Index
		yg.Players[v].Call("init", gi)
	}

	nap(10) // 牌组初始化
	for _, v := range yg.round {
		var s struct {
			Deck proto.Deck `json:"deck"`
		}
		yg.Players[v].Session.Data.DeJson(&s)
		//rego.INFO(string(yg.Players[v].Session.Data.Bytes()))
		yg.Players[v].initDeck(s.Deck.GetMain(), s.Deck.GetExtra())
	}

	nap(10) // 手牌初始化
	for _, v := range yg.round {
		yg.Players[v].init()
	}

	yg.Players[yg.round[0]].MsgPub("游戏开始，{self}先手！", nil)
	nap(10)
	for {
		for _, v := range yg.round {
			nap(5)
			yg.Current = yg.Players[v]
			if !yg.Current.IsFail() {
				yg.Current.round()
				//				if yg.CheckWinner(); yg.Over {
				//					return
				//				}
			} else {
				yg.Over = true
			}
			if yg.Over {
				yg.CallAll("over", nil)
				yg.Current.MsgPub("游戏结束", nil)
				return
			}
		}
	}
}
