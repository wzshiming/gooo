package defaul

import (
	"rego/agent"
	"ygo"
	"ygo/cards"
)

var DefaultDeck = []uint{}

func init() {
	for i := uint(0); i != 60; i++ {
		DefaultDeck = append(DefaultDeck, i+1)
	}
}

func DefaultRuleFrom(sess *agent.Session, camp int, deck []uint) *ygo.Player {
	player := &ygo.Player{
		Session:  sess,
		Camp:     camp,
		Hp:       4000,
		DrawSize: 1,
		MaxHp:    ^uint(0),
		MaxSdi:   6,
		Deck:     ygo.NewCardPile(),
		Hand:     ygo.NewCardPile(),
		Extra:    ygo.NewCardPile(),
		Side:     ygo.NewCardPile(),
		Removed:  ygo.NewCardPile(),
		Grave:    ygo.NewCardPile(),
		Mzone:    ygo.NewCardTile(5),
		Szone:    ygo.NewCardTile(5),
		Field:    ygo.NewCardTile(1),
	}
	player.Deck = cards.CardBag_test.Deck(player, deck)
	return player
}

func DefaultTest(sess *agent.Session, camp int) *ygo.Player {
	return DefaultRuleFrom(sess, camp, DefaultDeck)
}
