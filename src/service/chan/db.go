package main

import (
	"fmt"
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
	err = db.Where(&proto.DecksInfo{UserId: userid}).First(&r).Error
	return
}

// 用户第一次游戏创建卡组信息
func NewDecks(userid uint64) {
	_, err := getDecks(userid)
	if err == nil {
		return
	}
	b := proto.DecksInfo{
		UserId: userid,
		Def:    fmt.Sprint("Deck", 1),
	}
	db.Create(&b)
	NewDeck(userid)
	return
}

// 给用户添加新的卡组
func NewDeck(userid uint64) {
	odecks, err := getDecks(userid)
	if err != nil {
		return
	}
	odecks.Max++
	db.Save(&proto.Deck{
		UserId: userid,
		Name:   fmt.Sprint("Deck", odecks.Max),
	})
	db.Save(odecks)
	return
}

// 获得默认卡组
func GetDef(userid uint64) (p proto.Deck) {
	d, err := getDecks(userid)
	if err != nil {
		return
	}

	return GetDeck(userid, d.Def)
}

//func SetDef(userid uint64, deckid uint64) {
//	d, err := getDecks(userid)
//	if err != nil {
//		return
//	}
//	d.Def = deckid
//	db.Save(&d)
//}

func GetDecks(userid uint64) (odecks proto.DecksInfo) {
	if userid == 0 {
		return
	}
	if err := db.Where(&proto.DecksInfo{UserId: userid}).First(&odecks).Error; err != nil {
		NewDecks(userid)
		if err := db.Where(&proto.DecksInfo{UserId: userid}).First(&odecks).Error; err != nil {
			return
		}
	}

	db.Where(&proto.Deck{UserId: odecks.UserId}).Find(&odecks.List)
	for k, v := range odecks.List {
		db.Where(&proto.Card{DeckId: v.Id}).Find(&odecks.List[k].Main)
	}

	return
}

func GetDeck(userid uint64, name string) (odeck proto.Deck) {
	if userid == 0 {
		return
	}
	if err := db.Where(&proto.Deck{UserId: userid, Name: name}).First(&odeck).Error; err != nil {
		return
	}
	db.Where(&proto.Card{DeckId: odeck.Id}).Find(&odeck.Main)
	return
}

func SetDeck(userid uint64, name string, cas []proto.Card) {
	if userid == 0 {
		return
	}
	var odeck proto.Deck

	if err := db.Where(&proto.Deck{UserId: userid, Name: name}).First(&odeck).Error; err != nil {
		return
	}

	db.Where(&proto.Card{DeckId: odeck.Id}).Delete(&proto.Card{})
	odeck.Main = cas
	db.Save(&odeck)

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
