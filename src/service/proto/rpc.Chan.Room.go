package proto

import ()

type GameInitRequest struct {
}

type PlayerInit struct {
	// user 表示当前游戏中 卡片的id
	Deck  []uint `json:"deck"`
	Extra []uint `json:"extra"`
	Side  []uint `json:"side"`
	Hp    int    `json:"hp"`
	Name  string `json:"name"`
}

type GameInitResponse struct {
	// user 表示当前游戏中 卡片的id
	Index int          `json:"index"` //自身的索引
	Users []PlayerInit `json:"users"`
}

type SelectDeckRequest struct {
	ID uint `json:id`
}

type SelectDeckResponse struct {
	Deck  []uint `json:"deck"`
	Extra []uint `json:"extra"`
	Side  []uint `json:"side"`
}
