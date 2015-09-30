package main

import (
	"service/proto"

	"github.com/wzshiming/server/cfg"
	"github.com/wzshiming/server/dbs"
)

var db dbs.DB

func dbconn(us cfg.DbConfig) (err error) {
	db, err = dbs.Conn(us)
	db.CreateTable(&proto.Card{})
	db.CreateTable(&proto.Deck{})
	db.CreateTable(&proto.DecksInfo{})
	return
}

func getDecks(userid uint64) (r proto.DecksInfo, err error) {
	err = db.Where(&proto.DecksInfo{UserId: userid}).Find(&r).Error
	return
}

func NewDecks(userid uint64) {
	_, err := getDecks(userid)
	if err == nil {
		return
	}
	b := proto.DecksInfo{
		UserId: userid,
		Exi:    0,
		Max:    1,
		Def:    0,
	}
	db.Create(&b)
	NewDeck(userid)
	return
}

func NewDeck(userid uint64) {
	var odecks proto.DecksInfo
	if err := db.Where(&proto.DecksInfo{UserId: userid}).Find(&odecks).Error; err != nil {
		return
	}
	if odecks.Exi < odecks.Max {
		db.Save(&proto.Deck{
			UserId: userid,
			Main: []proto.Card{
				proto.Card{
					Size:  1,
					Index: 0,
					InPos: proto.InMain,
				},
			},
		})
	}
	return
}

func GetDef(userid uint64) (p proto.Deck) {
	d, err := getDecks(userid)
	if err != nil {
		return
	}
	if d.Def == 0 {
		b := GetDecks(userid)
		if len(b.List) != 0 {
			return b.List[0]
		}
	}
	return GetDeck(d.Def)
}

func SetDef(userid uint64, deckid uint64) {
	d, err := getDecks(userid)
	if err != nil {
		return
	}
	d.Def = deckid
	db.Save(&d)
}

func GetDecks(userid uint64) (odecks proto.DecksInfo) {
	if err := db.Where(&proto.DecksInfo{UserId: userid}).Find(&odecks).Error; err != nil {
		NewDecks(userid)
		if err := db.Where(&proto.DecksInfo{UserId: userid}).Find(&odecks).Error; err != nil {
			return
		}
	}

	db.Where(&proto.Deck{UserId: odecks.UserId}).Find(&odecks.List)
	for k, v := range odecks.List {
		db.Where(&proto.Card{DeckId: v.Id, InPos: proto.InMain}).Find(&odecks.List[k].Main)
		db.Where(&proto.Card{DeckId: v.Id, InPos: proto.InExtra}).Find(&odecks.List[k].Extra)
		db.Where(&proto.Card{DeckId: v.Id, InPos: proto.InSide}).Find(&odecks.List[k].Side)
	}
	return
}

func GetDeck(deckid uint64) (d proto.Deck) {
	db.Where(&proto.Deck{Id: deckid}).Find(&d)
	return
}

func SetDeck(userid uint64, deckid uint64, gr proto.Deck) {
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

	gr.UserId = userid
	if err := db.Where(&proto.Deck{Id: deckid, UserId: userid}).First(&odeck).Error; err != nil || odeck.Id != deckid {
		return
	}

	//db.Where(&proto.Card{DeckId: odeck.Id, Index: 0}).Delete(&proto.Card{})
	db.Where(&proto.Card{DeckId: odeck.Id}).Delete(&proto.Card{})

	db.Save(&gr)

	return
}

//func GetDeck(id uint64) (odecks []proto.Deck) {
//	db.Where(&proto.Deck{UserId: id}).Find(&odecks)
//	for k, _ := range odecks {
//		db.Where(&proto.Card{DeckId: odecks[k].Id, InPos: proto.InMain}).Find(&odecks[k].Main)
//		db.Where(&proto.Card{DeckId: odecks[k].Id, InPos: proto.InExtra}).Find(&odecks[k].Extra)
//		db.Where(&proto.Card{DeckId: odecks[k].Id, InPos: proto.InSide}).Find(&odecks[k].Side)
//	}
//	if len(odecks) == 0 && id != 0 {
//		odecks = append(odecks, proto.Deck{
//			Main: []proto.Card{
//				proto.Card{
//					Size:  1,
//					Index: 0,
//				},
//			},
//		})
//	}
//	return
//}

//func SetDeck(id uint64, name string, gr proto.Deck) {
//	var odeck proto.Deck
//	for k, _ := range gr.Main {
//		gr.Main[k].InPos = proto.InMain
//	}
//	for k, _ := range gr.Extra {
//		gr.Extra[k].InPos = proto.InExtra
//	}
//	for k, _ := range gr.Side {
//		gr.Side[k].InPos = proto.InSide
//	}
//	if err := db.Where(&proto.Deck{UserId: id, Name: name}).First(&odeck).Error; err != nil {
//		gr.CreatedAt = time.Now()
//		db.Model(&odeck).Save(&gr)
//	} else {
//		gr.UpdatedAt = time.Now()
//		db.Where(&proto.Card{DeckId: odeck.Id, Index: 0}).Delete(&proto.Card{})
//		db.Where(&proto.Card{DeckId: odeck.Id}).Delete(&proto.Card{})
//		db.Model(&odeck).Update(&gr)
//	}

//	return
//}

//func DelDeck(id uint64, name string) {
//	var odeck proto.Deck
//	if err := db.Where(&proto.Deck{UserId: id, Name: name}).First(&odeck).Error; err != nil {

//	} else {
//		db.Where(&proto.Card{DeckId: odeck.Id}).Delete(&proto.Card{})
//		db.Delete(&odeck)
//	}
//	return
//}
