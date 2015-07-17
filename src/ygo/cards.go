package ygo

import (
	"github.com/wzshiming/rego"
)

type Cards struct {
	list []*Card
}

func NewCards() *Cards {
	return &Cards{
		list: []*Card{},
	}
}

func (cp *Cards) Clear() {
	cp.list = []*Card{}
}

func (cp *Cards) Len() int {
	return len(cp.list)
}

func (cp *Cards) Shuffle() {
	array := cp.list
	for i := 0; i < len(array)*10; i++ {
		for j := 0; j < len(array)-1; j++ {
			if ((<-rego.LCG) % 4) != 0 {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func (cp *Cards) BeginPush(c *Card) {
	cp.Insert(c, len(cp.list))
}

func (cp *Cards) EndPush(c *Card) {
	cp.Insert(c, 0)
}

func (cp *Cards) Insert(c *Card, index int) {
	cp.list = append(cp.list[:index], append([]*Card{c}, cp.list[index:]...)...)

}

func (cp *Cards) BeginPop() (c *Card) {
	return cp.Remove(len(cp.list) - 1)
}

func (cp *Cards) EndPop() (c *Card) {
	return cp.Remove(0)
}

func (cp *Cards) Remove(index int) (c *Card) {
	if len(cp.list) == 0 {
		return
	}
	c = cp.list[index]
	cp.list = append(cp.list[:index], cp.list[index+1:]...)
	return
}

func (cp *Cards) PickedForUniq(uniq uint) (c *Card) {
	for k, v := range cp.list {
		if v.ToUint() == uniq {
			v.place = nil
			return cp.Remove(k)
		}
	}
	return
}

func (cp *Cards) ExistForUniq(uniq uint) (c *Card) {
	for _, v := range cp.list {
		if v.ToUint() == uniq {
			c = v
		}
	}
	return
}

func (cp *Cards) IsExistCard(c *Card) bool {
	return cp.ExistForUniq(c.ToUint()) == c
}

func (cp *Cards) PickedFor(c *Card) {
	cp.PickedForUniq(c.ToUint())
}

func (cp *Cards) Uniqs() (us []uint) {
	cp.ForEach(func(c *Card) bool {
		us = append(us, c.ToUint())
		return true
	})
	return
}

func (cp *Cards) ForEach(fun func(*Card) bool) {
	for _, v := range cp.list {
		if !fun(v) {
			return
		}
	}
}

func (cp *Cards) Find(fun func(*Card) bool) (indexs []int) {
	for k, v := range cp.list {
		if fun(v) {
			indexs = append(indexs, k)
		}
	}
	return
}

func (cp *Cards) Random() *Card {
	return cp.Remove(int(<-rego.LCG) % cp.Len())
}
