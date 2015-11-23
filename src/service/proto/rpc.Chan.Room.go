package proto

type GameInitRequest struct {
}

type SelectableRequest struct {
	Uniq   uint `json:"uniq"`
	Method uint `json:"method"`
}

type PlayerInit struct {
	Hp   int    `json:"hp"`
	Name string `json:"name"`
}

type GameInitResponse struct {
	// user 表示当前游戏中 卡片的id
	Index int          `json:"index"` //自身的索引
	Users []PlayerInit `json:"users"`
}

//type SelectDeckRequest struct {
//	ID uint `json:"id"`
//}

//type SelectDeckResponse struct {
//	Deck  []uint `json:"deck"`
//	Extra []uint `json:"extra"`
//	Side  []uint `json:"side"`
//}
