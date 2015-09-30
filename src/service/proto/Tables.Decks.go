package proto

import "time"

type DecksInfo struct {
	Id     uint64 `json:"-";gorm:"primary_key"`
	UserId uint64 `json:"-";sql:"unique_index"`
	List   []Deck `json:"list"`
	Def    uint64 `json:"def"`
	Exi    int    `json:"Exi"`
	Max    int    `json:"max"`
}

type Deck struct {
	Id     uint64 `json:"id";gorm:"primary_key"`
	UserId uint64 `json:"-";sql:"index"`
	//DecksId uint64 `json:"-";sql:"index"`
	//Name      string    `json:"name";sql:"type:varchar(64);not null;index"`
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

type Card struct {
	Id     uint64 `json:"-";gorm:"primary_key"`
	DeckId uint64 `json:"-";sql:"index:idx_deckid_index_inpos"`
	Index  uint   `json:"id";sql:"index:idx_deckid_index_inpos"`
	Size   uint   `json:"size"`
	InPos  InPos  `json:"-";sql:"index:idx_deckid_index_inpos"`
}

type InPos uint

const (
	InNone = InPos(iota)
	InMain
	InExtra
	InSide
)
