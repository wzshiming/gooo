package main

import (
	"rego/cfg"
	"rego/dbs"
	"service/proto"
)

var db dbs.DB

func dbconn(us cfg.DbConfig) (err error) {
	db, err = dbs.Conn(us)
	var decks proto.Deck
	decks.CreateTable(&db)
	return
}

func GetDeck(id uint64) (odecks []proto.Deck) {
	db.Where(&proto.Deck{UserId: id}).Find(&odecks)
	for k, _ := range odecks {
		db.Where(&proto.CardInMain{DeckId: odecks[k].Id}).Find(&odecks[k].Main)
		db.Where(&proto.CardInExtra{DeckId: odecks[k].Id}).Find(&odecks[k].Extra)
		db.Where(&proto.CardInSide{DeckId: odecks[k].Id}).Find(&odecks[k].Side)
	}
	return
}
