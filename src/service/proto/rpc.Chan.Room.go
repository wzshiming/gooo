package proto

type GameInitRequest struct {
}

type SelectableRequest struct {
	Uniq   uint `json:"uniq"`
	Method uint `json:"method"`
}

//type SelectDeckRequest struct {
//	ID uint `json:"id"`
//}

//type SelectDeckResponse struct {
//	Deck  []uint `json:"deck"`
//	Extra []uint `json:"extra"`
//	Side  []uint `json:"side"`
//}
