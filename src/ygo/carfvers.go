package ygo

import (
	"errors"
)

type CardVersion struct {
	List map[uint]*CardOriginal
}

func NewCardVersion() *CardVersion {
	return &CardVersion{
		List: make(map[uint]*CardOriginal),
	}
}

func (cv *CardVersion) Keys() (c []uint) {
	for k, _ := range cv.List {
		c = append(c, k)
	}
	return
}

func (cv *CardVersion) Get(id uint) *CardOriginal {
	return cv.List[id]
}

func (cv *CardVersion) Register(co *CardOriginal) error {
	if co == nil {
		return errors.New("RegisterCard: Nil")
	}
	if co.Name == "" {
		return errors.New("RegisterCard: Name is empty")
	}
	if co.Id == 0 || cv.List[co.Id] != nil {
		return errors.New("RegisterCard: Duplicate ID")
	}
	cv.List[co.Id] = co
	return nil
}

func (cv *CardVersion) Sum(cv2 *CardVersion) *CardVersion {
	rcv := NewCardVersion()
	for _, v := range cv.List {
		cv.Register(v)
	}
	for _, v := range cv2.List {
		cv.Register(v)
	}
	return rcv
}

func (cv *CardVersion) Deck(cp *CardPile, player *Player, deck []uint) {
	for _, v := range deck {
		t := cv.Get(v)
		if t != nil {
			c := t.Make(player)
			cp.EndPush(c)
			player.Game.RegisterCards(c)
		}
	}
}
