package ygo

import (
	//"rego"
	"rego/agent"
	"rego/misc"
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
		p.Game = yg
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
	time.Sleep(time.Second)
	yg.CallAll("init", nil)
	time.Sleep(time.Second * 3)
	i := 0
	for _, v := range yg.Players {
		ca := v.Camp
		yg.Survival[ca] = yg.Survival[ca] + 1
		v.Index = i
		v.Game = yg
		v.init()
		i++
	}
	time.Sleep(time.Second)
	for {
		for _, v := range yg.Players {
			time.Sleep(time.Second)
			if !v.IsFail() {
				v.round()
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
