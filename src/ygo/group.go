package ygo

import (
	"github.com/wzshiming/dispatcher"
)

type Group struct {
	Cards
	dispatcher.Events
	name  LL_TYPE
	owner *Player // 所有者
}

func NewGroup(pl *Player, name LL_TYPE) *Group {
	return &Group{
		Events: dispatcher.NewForkEvent(pl.GetFork()),
		owner:  pl,
		name:   name,
		Cards:  *NewCards(),
	}
}

func (cp *Group) GetOwner() *Player {
	return cp.owner
}

func (cp *Group) SetName(name LL_TYPE) {
	cp.name = name
}

func (cp *Group) GetName() LL_TYPE {
	return cp.name
}

func (cp *Group) BeginPush(c *Card) {
	cp.Insert(c, cp.Len())
}

func (cp *Group) EndPush(c *Card) {
	cp.Insert(c, 0)
}

func (cp *Group) Insert(c *Card, index int) {
	if c.place != nil {
		c.place.PickedFor(c)
	}
	c.GetSummoner().CallAll(MoveCard(c, cp.GetName()))

	c.place = cp
	cp.Cards.Insert(c, index)
	c.Dispatch(In + string(cp.GetName()))
}

func (cp *Group) BeginPop() (c *Card) {
	return cp.Remove(cp.Len() - 1)
}

func (cp *Group) EndPop() (c *Card) {
	return cp.Remove(0)
}

func (cp *Group) Remove(index int) (c *Card) {
	c = cp.Cards.Remove(index)
	if c != nil {
		c.place = nil
		c.Dispatch(Out + string(cp.GetName()))
	}
	return
}

func (cp *Group) PickedForUniq(uniq uint) (c *Card) {
	c = cp.Cards.PickedForUniq(uniq)
	if c != nil {
		c.Dispatch(Out + string(cp.GetName()))
	}
	return
}

func (cp *Group) PickedFor(c *Card) {
	cp.Cards.PickedFor(c)
	c.Dispatch(Out + string(cp.GetName()))
}
