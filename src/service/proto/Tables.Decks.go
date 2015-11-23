package proto

import (
	"time"
)

type InPos uint

const (
	InNone = InPos(iota)
	InMain
	InSide
)

type Card struct {
	Id     uint64 `json:"-";gorm:"primary_key"`
	DeckId uint64 `json:"-";sql:"index:idx_deckid_index_inpos"`
	Index  uint   `json:"id";sql:"index:idx_deckid_index_inpos"`
	Size   uint   `json:"size"`
	InPos  InPos  `json:"pos";sql:"index:idx_deckid_index_inpos"`
}


type DecksInfo struct {
	Id     uint64 `json:"-";gorm:"primary_key"`
	UserId uint64 `json:"-";sql:"unique_index"`
	List   []Deck `json:"list"`
	Def    string `json:"def"`
	Max    int    `json:"max"`
}

type Deck struct {
	Id     uint64 `json:"id";gorm:"primary_key"`
	UserId uint64 `json:"-";sql:"index"`
	//DecksId uint64 `json:"-";sql:"index"`
	Name string `json:"name";sql:"type:varchar(64);not null;index"`
	Main []Card `json:"main"`
	//Side      []Card    `json:"side"`
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
