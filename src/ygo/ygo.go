package ygo

import (
	"time"
)

type YGO struct {
	StartAt time.Time
	Players []*Player
	Over    bool
}

func NewYGO(players ...*Player) *YGO {
	return &YGO{
		StartAt: time.Now(),
		Players: players,
	}
}

func (yg *YGO) ForEachPlayer(fun func(*Player)) {
	for k, _ := range yg.Players {
		fun(yg.Players[k])
	}
}

func (yg *YGO) Loop() {
	for k, _ := range yg.Players {
		yg.Players[k].Index = k
		yg.Players[k].Game = yg
		yg.Players[k].init()
	}
	for {
		for _, v := range yg.Players {
			v.round()
			if yg.Over {
				return
			}
		}
	}
}

func (yg *YGO) Winner(index ...int) {
	yg.Over = true
}
