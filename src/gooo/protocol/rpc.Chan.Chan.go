package protocol

import ()

type CreateRequest struct {
	Size int `json:"s"`
}

type CreateResponse struct {
	RoomId int `json:"r"`
	SeatId int `json:"s"`
}

type JoinRequest struct {
	RoomId int `json:"r"`
}

type JoinResponse struct {
	RoomId int `json:"r"`
	SeatId int `json:"s"`
}
