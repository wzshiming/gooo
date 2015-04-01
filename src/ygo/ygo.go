package ygo

import (
	"rego"
	"time"
)

type YGO struct {
	StartAt  time.Time
	Players  []*Player
	Survival map[int]int
	Over     bool
}

func NewYGO(players ...*Player) *YGO {
	return &YGO{
		Survival: make(map[int]int),
		StartAt:  time.Now(),
		Players:  players,
	}
}

func (yg *YGO) ForEachPlayer(fun func(*Player)) {
	for k, _ := range yg.Players {
		fun(yg.Players[k])
	}
}

func (yg *YGO) Loop() {
	for k, _ := range yg.Players {
		ca := yg.Players[k].Camp
		yg.Survival[ca] = yg.Survival[ca] + 1
		yg.Players[k].Index = k
		yg.Players[k].Game = yg
		yg.Players[k].init()

	}
	for {
		for _, v := range yg.Players {
			if !v.IsFail() {
				v.round()
				if yg.CheckWinner(); yg.Over {
					return
				}
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
