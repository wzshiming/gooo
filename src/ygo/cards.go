package ygo

type Cards interface {
	Index(c *Card) int
	Placed(index int, c *Card)
	Picked(index int) (c *Card)
	PickedFor(c *Card)
	ForEach(fun func(*Card))
	Find(fun func(*Card) bool) (indexs []int)
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

type CardPile struct {
	CardUnfold
}

func NewCardPile() *CardPile {
	return &CardPile{
		CardUnfold: CardUnfold{},
	}
}

func (cp *CardPile) BeginPush(c *Card) {
	c.Place = cp
	cp.CardUnfold = append(cp.CardUnfold, c)
}

func (cp *CardPile) EndPush(c *Card) {
	c.Place = cp
	cp.CardUnfold = append(CardUnfold{c}, cp.CardUnfold...)
}

func (cp *CardPile) BeginPop() (c *Card) {
	c = (cp.CardUnfold)[len(cp.CardUnfold)-1]
	cp.CardUnfold = (cp.CardUnfold)[:len(cp.CardUnfold)-2]
	c.Place = nil
	return
}

func (cp *CardPile) EndPop() (c *Card) {
	c = (cp.CardUnfold)[0]
	cp.CardUnfold = (cp.CardUnfold)[1:]
	c.Place = nil
	return
}
