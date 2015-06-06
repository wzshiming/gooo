package proto

import (
	"rego/dbs"
	"time"
)

type Deck struct {
	Id        uint64        `json:"-"`
	UserId    uint64        `json:"-";sql:"index"`
	Name      string        `json:"name";sql:"type:varchar(64);not null;index"`
	Main      []CardInMain  `json:"main"`
	Extra     []CardInExtra `json:"extra"`
	Side      []CardInSide  `json:"side"`
	CreatedAt time.Time     `json:"createdat"`
	UpdatedAt time.Time     `json:"updatedat"`
}

func (s *Deck) CreateTable(db *dbs.DB) {
	db.CreateTable(s)
	db.CreateTable(&CardInMain{})
	db.CreateTable(&CardInExtra{})
	db.CreateTable(&CardInSide{})

}

func (s *Deck) DropTable(db *dbs.DB) {
	db.DropTable(s)
}

type Card struct {
	Id     uint64 `json:"-"`
	DeckId uint64 `json:"-";sql:"index:idx_deckid_index"`
	Index  uint64 `json:"id";sql:"index:idx_deckid_index"`
	Size   uint64 `json:"size"`
}

type CardInMain Card
type CardInExtra Card
type CardInSide Card
