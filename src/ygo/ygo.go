package ygo

import (
	"rego"
	"rego/agent"
	"rego/misc"
	"time"
)

type YGO struct {
	CardVer  *CardVersion
	Room     *misc.Rooms
	StartAt  time.Time
	Cards    map[uint]*Card
	Players  []*Player
	Survival map[int]int
	Over     bool
}

func NewYGO(r *misc.Rooms) *YGO {
	yg := &YGO{
		Room:     r,
		Cards:    make(map[uint]*Card),
		Survival: make(map[int]int),
		StartAt:  time.Now(),
		//Players:  players,
	}
	yg.Room.ForEach(func(sess *agent.Session) {
		p := NewPlayer()
		p.Game = yg
		p.Session = sess
		yg.Players = append(yg.Players, p)
	})
	return yg
}

func (yg *YGO) GetCard(u uint) *Card {
	return yg.Cards[u]
}

func (yg *YGO) RegisterCards(c *Card) {
	yg.Cards[c.ToUint()] = c
}

func (yg *YGO) AddPlayer(pl *Player) {
	yg.Players = append(yg.Players, pl)
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
	}, nil)
	return nil
}

func (yg *YGO) Loop() {
	time.Sleep(time.Second)
	yg.CallAll("init", nil)
	time.Sleep(time.Second)
	for k, _ := range yg.Players {
		ca := yg.Players[k].Camp
		yg.Survival[ca] = yg.Survival[ca] + 1
		yg.Players[k].Index = k
		yg.Players[k].Game = yg
		yg.Players[k].init()
	}
	time.Sleep(time.Second)
	for {
		for _, v := range yg.Players {
			if !v.IsFail() {
				v.round()
				//				if yg.CheckWinner(); yg.Over {
				//					return
				//				}
			}
		}
	}
}

func (yg *YGO) AddLoser(camp int) {
	yg.Survival[camp] = yg.Survival[camp] - 1
}

func (yg *YGO) CheckWinner() {
	i := 0
	ca := 0
	for k, v := range yg.Survival {
		if v > 0 {
			i += 1
			ca = k
		}
	}
	//平局
	if i == 0 {
		rego.NOTICE("Game not winner")
		yg.Over = true
	}
	//
	if i == 1 {
		rego.NOTICE("Game winners the", ca, "group")
		yg.Over = true
	}
}
