package ygo

import (
	"fmt"
	"github.com/wzshiming/dispatcher"
	"github.com/wzshiming/rego/agent"
	"github.com/wzshiming/rego/misc"
	"service/proto"
	"time"
)

type YGO struct {
	dispatcher.Events
	CardVer  *CardVersion
	Room     *misc.Rooms
	StartAt  time.Time
	Cards    map[uint]*Card
	Players  map[uint]*Player
	Survival map[int]int
	Over     bool
	round    []uint
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
		yg.Players[v].RoundSize = 1
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

	nap(10)
	for {
		for _, v := range yg.round {
			nap(5)
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
