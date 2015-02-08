package protocol

import ()

type RoomsRequest struct {
	Get bool `json:"g"`
}

type RoomsResponse []byte //[]room.GameRoom

type NameRequest struct {
	Get bool `json:"g"`
}

type NameResponse struct {
	Name string `json:"n"`
}
