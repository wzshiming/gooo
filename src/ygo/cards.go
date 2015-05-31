package ygo

import (
	"rego"
)

type CardPile struct {
	name string
	list []*Card
	join func(*Card)
}

func NewCardPile(name string) *CardPile {
	return &CardPile{
		name: name,
		list: []*Card{},
	}
}

func (cp *CardPile) SetJoin(fun func(*Card)) {
	cp.join = fun
}

func (cp *CardPile) Clear() {
	cp.list = []*Card{}
}

func (cp *CardPile) SetName(name string) {
	cp.name = name
}
func (cp *CardPile) GetName() string {
	return cp.name
}

func (cp *CardPile) Len() int {
	return len(cp.list)
}

func (cp *CardPile) Shuffle() {
	array := cp.list
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1; j++ {
			if ((<-rego.LCG) % 4) != 0 {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	//	for i := 0; i < len(array); i++ {
	//		array[i].Owner.CallAll(MoveCard(array[i], cp.GetName()))
	//	}
}

func (cp *CardPile) BeginPush(c *Card) {
	cp.Insert(c, len(cp.list))
}

func (cp *CardPile) EndPush(c *Card) {
	cp.Insert(c, 0)
}

func (cp *CardPile) Insert(c *Card, index int) {

	if c.Place != nil {
		c.Place.PickedFor(c)
	}
	c.Owner.CallAll(MoveCard(c, cp.GetName()))
	if cp.join != nil {
		cp.join(c)
	}
	c.Place = cp
	cp.list = append(cp.list[:index], append([]*Card{c}, cp.list[index:]...)...)

}

func (cp *CardPile) BeginPop() (c *Card) {
	return cp.Remove(len(cp.list) - 1)
}

func (cp *CardPile) EndPop() (c *Card) {
	return cp.Remove(0)
}

func (cp *CardPile) Remove(index int) (c *Card) {
	if len(cp.list) == 0 {
		return
	}
	c = cp.list[index]
	c.Place = nil
	cp.list = append(cp.list[:index], cp.list[index+1:]...)
	return
}

func (cp *CardPile) PickedForUniq(uniq uint) (c *Card) {
	for k, v := range cp.list {
		if v.ToUint() == uniq {
			v.Place = nil
			return cp.Remove(k)
		}
	}
	return
}

func (cp *CardPile) ExistForUniq(uniq uint) (c *Card) {
	for _, v := range cp.list {
		if v.ToUint() == uniq {
			c = v
		}
	}
	return
}

func (cp *CardPile) PickedFor(c *Card) {
	cp.PickedForUniq(c.ToUint())
}

func (cp *CardPile) Uniqs() (us []uint) {
	cp.ForEach(func(c *Card) {
		us = append(us, c.ToUint())
	})
	return
}

func (cp *CardPile) ForEach(fun func(*Card)) {
	for _, v := range cp.list {
		fun(v)
	}
}

func (cp *CardPile) Find(fun func(*Card) bool) (indexs []int) {
	for k, v := range cp.list {
		if fun(v) {
			indexs = append(indexs, k)
		}
	}
	return
}
