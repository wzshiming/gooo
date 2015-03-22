package ygo

type Cards interface {
	Index(c *Card) int
	Placed(index int, c *Card)
	Picked(index int) (c *Card)
	PickedFor(c *Card)
	ForEach(fun func(*Card))
	Find(fun func(*Card) bool) (indexs []int)
}

type CardPile []*Card

func NewCardPile() *CardPile {
	return &CardPile{}
}
func (cp *CardPile) Shuffle() {
}

func (cp *CardPile) BeginPush(c *Card) {
	c.Place = cp
	*cp = append(*cp, c)
}

func (cp *CardPile) EndPush(c *Card) {
	c.Place = cp
	*cp = append(CardPile{c}, *cp...)
}

func (cp *CardPile) BeginPop() (c *Card) {
	c = (*cp)[len(*cp)-1]
	*cp = (*cp)[:len(*cp)-2]
	c.Place = nil
	return
}

func (cp *CardPile) EndPop() (c *Card) {
	c = (*cp)[0]
	*cp = (*cp)[1:]
	c.Place = nil
	return
}

func (cp *CardPile) Index(c *Card) int {
	for k, v := range *cp {
		if c == v {
			return k
		}
	}
	return -1
}

func (cp *CardPile) Placed(index int, c *Card) {
	t := (*cp)[index:]
	c.Place = cp
	(*cp) = append((*cp)[:index], c)
	(*cp) = append((*cp)[:index], t...)
}

func (cp *CardPile) Picked(index int) (c *Card) {
	c = (*cp)[index]
	c.Place = nil
	(*cp) = append((*cp)[:index], (*cp)[index+1:]...)
	return
}

func (cp *CardPile) PickedFor(c *Card) {
	for k, v := range *cp {
		if v == c {
			c.Place = nil
			cp.Picked(k)
		}
	}
	return
}

func (cp *CardPile) ForEach(fun func(*Card)) {
	for _, v := range *cp {
		fun(v)
	}
}

func (cp *CardPile) Find(fun func(*Card) bool) (indexs []int) {
	for k, v := range *cp {
		if fun(v) {
			indexs = append(indexs, k)
		}
	}
	return nil
}

type CardUnfold []*Card

func NewCardUnfold(max int) *CardUnfold {
	r := make(CardUnfold, max)
	return &r
}

func (cu *CardUnfold) Exist(index int) (c *Card) {
	c = (*cu)[index]
	return
}

func (cu *CardUnfold) Index(c *Card) int {
	for k, v := range *cu {
		if v != nil {
			if c == v {
				return k
			}
		}
	}
	return -1
}

func (cu *CardUnfold) Placed(index int, c *Card) {
	c.Place = cu
	(*cu)[index] = c
	return
}

func (cu *CardUnfold) Picked(index int) (c *Card) {
	c = (*cu)[index]
	(*cu)[index] = nil
	c.Place = nil
	return
}

func (cu *CardUnfold) PickedFor(c *Card) {
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

func (cu *CardUnfold) ForEach(fun func(*Card)) {
	for _, v := range *cu {
		if v != nil {
			fun(v)
		}
	}
}

func (cu *CardUnfold) Find(fun func(*Card) bool) (indexs []int) {
	for k, v := range *cu {
		if v != nil {
			if fun(v) {
				indexs = append(indexs, k)
			}
		}
	}
	return nil
}
