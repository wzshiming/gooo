package proto

import (
	"github.com/wzshiming/rego/dbs"
	"time"
)

type Deck struct {
	Id        uint64    `json:"-"`
	UserId    uint64    `json:"-";sql:"index"`
	Name      string    `json:"name";sql:"type:varchar(64);not null;index"`
	Main      []Card    `json:"main"`
	Extra     []Card    `json:"extra"`
	Side      []Card    `json:"side"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

func (s *Deck) GetMain() (r []uint) {
	for _, v := range s.Main {
		for i := uint(0); i != v.Size; i++ {
			r = append(r, v.Index)
		}
	}
	return
}

func (s *Deck) GetExtra() (r []uint) {
	for _, v := range s.Extra {
		for i := uint(0); i != v.Size; i++ {
			r = append(r, v.Index)
		}
	}
	return
}

func (s *Deck) CreateTable(db *dbs.DB) {
	db.CreateTable(s)
	db.CreateTable(&Card{})
}

func (s *Deck) DropTable(db *dbs.DB) {
	db.DropTable(s)
}

type Card struct {
	Id     uint64 `json:"-"`
	DeckId uint64 `json:"-";sql:"index:idx_deckid_index"`
	Index  uint   `json:"id";sql:"index:idx_deckid_index"`
	Size   uint   `json:"size"`
	InPos  InPos  `json:"-"`
}

type InPos uint

const (
	InNone = InPos(iota)
	InMain
	InExtra
	InSide
)
