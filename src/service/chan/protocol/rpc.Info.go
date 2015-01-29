package protocol

import ()

type RoomsRequest struct {
	Get bool `json:"g"`
}

type RoomsResponse []struct {
	Users   []uint64 `json:"u"`
	RoomId  int      `json:"r"`
	Name    string   `json:"n"`
	Started bool     `json:"s"`
}

type NameRequest struct {
	Get bool `json:"g"`
}

type NameResponse struct {
	Name string `json:"n"`
}
