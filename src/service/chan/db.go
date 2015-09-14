package main

import (
	"service/proto"
	"time"

	"github.com/wzshiming/server/cfg"
	"github.com/wzshiming/server/dbs"
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
		db.Where(&proto.Card{DeckId: odecks[k].Id, InPos: proto.InMain}).Find(&odecks[k].Main)
		db.Where(&proto.Card{DeckId: odecks[k].Id, InPos: proto.InExtra}).Find(&odecks[k].Extra)
		db.Where(&proto.Card{DeckId: odecks[k].Id, InPos: proto.InSide}).Find(&odecks[k].Side)
	}
	if len(odecks) == 0 && id != 0 {
		return GetDeck(0)
	}
	return
}

func SetDeck(id uint64, name string, gr proto.Deck) {
	var odeck proto.Deck
	for k, _ := range gr.Main {
		gr.Main[k].InPos = proto.InMain
	}
	for k, _ := range gr.Extra {
		gr.Extra[k].InPos = proto.InExtra
	}
	for k, _ := range gr.Side {
		gr.Side[k].InPos = proto.InSide
	}
	if err := db.Where(&proto.Deck{UserId: id, Name: name}).First(&odeck).Error; err != nil {
		gr.CreatedAt = time.Now()
		db.Model(&odeck).Save(&gr)
	} else {
		gr.UpdatedAt = time.Now()
		db.Where(&proto.Card{DeckId: odeck.Id, Index: 0}).Delete(&proto.Card{})
		db.Where(&proto.Card{DeckId: odeck.Id}).Delete(&proto.Card{})
		db.Model(&odeck).Update(&gr)
	}

	return
}

func DelDeck(id uint64, name string) {
	var odeck proto.Deck
	if err := db.Where(&proto.Deck{UserId: id, Name: name}).First(&odeck).Error; err != nil {

	} else {
		db.Where(&proto.Card{DeckId: odeck.Id}).Delete(&proto.Card{})
		db.Delete(&odeck)
	}
	return
}
