package ygo

import (
	"github.com/wzshiming/rego"
)

type Cards []*Card

func NewCards() *Cards {
	return &Cards{}
}

func (cp *Cards) Clear() {
	*cp = []*Card{}
}

func (cp *Cards) Clone() (p *Cards) {
	p = NewCards()
	c := append(*p, (*cp)...)
	*p = c
	return
}

func (cp *Cards) Len() int {
	return len(*cp)
}

func (cp *Cards) Get(index int) *Card {
	return (*cp)[index]
}

func (cp *Cards) Shuffle() {
	for i := 0; i != 20; i++ {
		cp.SortFor(func(c1, c2 *Card) bool {
			return ((<-rego.LCG) % 5) != 0
		})
	}
}

func (cp *Cards) SortFor(f func(c1, c2 *Card) bool) {
	array := *cp
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if f(array[j], array[j+1]) {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func (cp *Cards) BeginPush(c *Card) {
	cp.Insert(c, len(*cp))
}

func (cp *Cards) EndPush(c *Card) {
	cp.Insert(c, 0)
}

func (cp *Cards) Insert(c *Card, index int) {
	(*cp) = append((*cp)[:index], append([]*Card{c}, (*cp)[index:]...)...)

}

func (cp *Cards) BeginPop() (c *Card) {
	return cp.Remove(len(*cp) - 1)
}

func (cp *Cards) EndPop() (c *Card) {
	return cp.Remove(0)
}

func (cp *Cards) Remove(index int) (c *Card) {
	if len(*cp) == 0 {
		return
	}
	c = (*cp)[index]
	(*cp) = append((*cp)[:index], (*cp)[index+1:]...)
	return
}

func (cp *Cards) PickedForUniq(uniq uint) (c *Card) {
	for k, v := range *cp {
		if v.ToUint() == uniq {
			v.place = nil
			return cp.Remove(k)
		}
	}
	return
}

func (cp *Cards) ExistForUniq(uniq uint) (c *Card) {
	for _, v := range *cp {
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
	for _, v := range *cp {
		if !fun(v) {
			return
		}
	}
}

func (cp *Cards) Find(fun func(*Card) bool) (indexs []int) {
	for k, v := range *cp {
		if fun(v) {
			indexs = append(indexs, k)
		}
	}
	return
}

func (cp *Cards) Random() *Card {
	return cp.Remove(int(<-rego.LCG) % cp.Len())
}
