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
	yg := &YGO{
		Events:   dispatcher.NewLineEvent(),
		Room:     r,
		Cards:    make(map[uint]*Card),
		Survival: make(map[int]int),
		StartAt:  time.Now(),
		Players:  make(map[uint]*Player),

		pending:   make(map[uint]*Card),
		cardevent: make(map[string]map[*Card]Action),
	}
	yg.Room.ForEach(func(sess *agent.Session) {
		p := NewPlayer()
		p.Session = sess
		yg.Players[sess.ToUint()] = p
	})

	//	yg.AddEventListener(FindCards, func(cards *Cards, ld LL_TYPE, lc LC_TYPE, la LA_TYPE, lr LR_TYPE, le LE_TYPE) {
	//		t := []*Card{}
	//		if ld != LL_None {
	//			for _, v := range yg.Cards {
	//				if v.GetPlace().GetName() == ld {
	//					t = append(t, v)
	//				}
	//			}
	//		}
	//		if lc != LC_None {
	//			for _, v := range t {
	//				if v.GetPlace().GetName() == ld {
	//					t = append(t, v)
	//				}
	//			}
	//		}
	//	})

	return yg
}

func (yg *YGO) AddReply(ca *Card) {
	yg.pending[ca.ToUint()] = ca
}

func (yg *YGO) Chain(eventName string, ca *Card, args []interface{}) {
	event0 := TriggerChoose + eventName
	//	rego.INFO(eventName)
	//	rego.INFO(Take)
	//rego.INFO(eventName[:len(Take)])
	if yg.cardevent[event0] != nil {
		pl := ca.GetSummoner()
		pl.Msg("连锁事件 "+eventName+" {self}", nil)
		cs := []*Card{}
		for k, v := range yg.cardevent[event0] {
			if k != ca && v.Call(ca) {
				cs = append(cs, k)
			}
		}

		if !ca.GetSummoner().Chain(event0, ca, cs, args) {
			ca.GetSummoner().GetTarget().Chain(event0, ca, cs, args)
		}

		yg.pending = map[uint]*Card{}
	}
}

func (yg *YGO) GetCardEvents(event string) map[*Card]Action {
	return yg.cardevent[event]
}

func (yg *YGO) RegisterCardEvents(event string, ca *Card, condition Action) {
	if yg.cardevent[event] == nil {
		yg.cardevent[event] = map[*Card]Action{}
	}
	yg.cardevent[event][ca] = condition
}

func (yg *YGO) UnregisterCardEvents(event string, ca *Card) {
	if yg.cardevent[event] != nil {
		delete(yg.cardevent[event], ca)
		if len(yg.cardevent[event]) == 0 {
			delete(yg.cardevent, event)
		}
	}

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
