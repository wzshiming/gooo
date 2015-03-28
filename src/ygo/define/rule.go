package define

import (
	"rego/agent"
	"ygo"
	"ygo/cards"
)

func DefaultRuleFrom(sess *agent.Session, camp uint, deck []uint) *ygo.Player {
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
		Mzone:    ygo.NewCardUnfold(5),
		Szone:    ygo.NewCardUnfold(5),
		Field:    ygo.NewCardUnfold(1),
	}
	player.Deck = cards.CardBag_VOL01.Deck(player, deck)
	return player
}

func DefaultTest(sess *agent.Session, camp uint) *ygo.Player {
	return DefaultRuleFrom(sess, camp, cards.CardBag_VOL01.Keys())
}
