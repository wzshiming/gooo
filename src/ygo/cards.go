package ygo

type CardPile []*Card

func NewCardPile() *CardPile {
	return &CardPile{}
}

func (cp *CardPile) BeginPush(c *Card) {
	*cp = append(*cp, c)
}

func (cp *CardPile) EndPush(c *Card) {
	*cp = append(CardPile{c}, *cp...)
}

func (cp *CardPile) BeginPop() (c *Card) {
	c = (*cp)[len(*cp)-1]
	*cp = (*cp)[:len(*cp)-2]
	return
}

func (cp *CardPile) EndPop() (c *Card) {
	c = (*cp)[0]
	*cp = (*cp)[1:]
	return
}

func (cp *CardPile) MidExtract(index int) (c *Card) {
	c = (*cp)[index]
	(*cp) = append((*cp)[:index], (*cp)[index+1:]...)
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

func (cu *CardUnfold) Existsd(index int) (c *Card) {
	c = (*cu)[index]
	return
}

func (cu *CardUnfold) Place(index int, c *Card) {
	(*cu)[index] = c
	return
}

func (cu *CardUnfold) Picked(index int) (c *Card) {
	c = (*cu)[index]
	(*cu)[index] = nil
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
