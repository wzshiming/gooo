package ygo

import (
	"rego"
)

type Cards interface {
	Index(c *Card) int
	Placed(index int, c *Card)
	Picked(index int) (c *Card)
	PickedFor(c *Card)
	ForEach(fun func(*Card))
	Find(fun func(*Card) bool) (indexs []int)
}

type CardTile []*Card

func NewCardTile(max int) *CardTile {
	r := make(CardTile, max)
	return &r
}

func (cu *CardTile) Len() int {
	return len(*cu)
}

func (cu *CardTile) Exist(index int) (c *Card) {
	c = (*cu)[index]
	return
}

func (cu *CardTile) Index(c *Card) int {
	for k, v := range *cu {
		if v != nil {
			if c == v {
				return k
			}
		}
	}
	return -1
}

func (cu *CardTile) Placed(index int, c *Card) {
	c.Place = cu
	(*cu)[index] = c
	return
}

func (cu *CardTile) Picked(index int) (c *Card) {
	c = (*cu)[index]
	(*cu)[index] = nil
	c.Place = nil
	return
}

func (cu *CardTile) PickedFor(c *Card) {
	for k, v := range *cu {
		if v != nil {
			if v == c {
				c.Place = nil
				cu.Picked(k)
			}
		}
	}
	return
}

func (cu *CardTile) ForEach(fun func(*Card)) {
	for _, v := range *cu {
		if v != nil {
			fun(v)
		}
	}
}

func (cu *CardTile) Find(fun func(*Card) bool) (indexs []int) {
	for k, v := range *cu {
		if v != nil {
			if fun(v) {
				indexs = append(indexs, k)
			}
		}
	}
	return nil
}

func (cu *CardTile) Uniqs() (us []uint) {
	cu.ForEach(func(c *Card) {
		us = append(us, c.ToUint())
	})
	return
}

type CardPile struct {
	CardTile
}

func NewCardPile() *CardPile {
	return &CardPile{
		CardTile: CardTile{},
	}
}

func (cp *CardPile) Shuffle() {
	array := cp.CardTile
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if ((<-rego.LCG) % 4) != 0 {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func (cp *CardPile) BeginPush(c *Card) {
	c.Place = cp
	cp.CardTile = append(cp.CardTile, c)
}

func (cp *CardPile) EndPush(c *Card) {
	c.Place = cp
	cp.CardTile = append(CardTile{c}, cp.CardTile...)
}

func (cp *CardPile) BeginPop() (c *Card) {
	if len(cp.CardTile) == 1 {
		c = (cp.CardTile)[0]
		cp.CardTile = CardTile{}
		return
	}
	c = (cp.CardTile)[len(cp.CardTile)-1]
	cp.CardTile = (cp.CardTile)[:len(cp.CardTile)-2]
	c.Place = nil
	return
}

func (cp *CardPile) EndPop() (c *Card) {
	if len(cp.CardTile) == 1 {
		c = (cp.CardTile)[0]
		cp.CardTile = CardTile{}
		return
	}
	c = (cp.CardTile)[0]
	cp.CardTile = (cp.CardTile)[1:]
	c.Place = nil
	return
}
