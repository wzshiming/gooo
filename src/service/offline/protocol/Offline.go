package protocol

import ()

type InterruptRequest struct {
	Data []byte `json:"d"`
}

type InterruptResponse struct {
	UserId uint64 `json:"u"`
}

type ReconnectionRequest struct {
	UserId uint64 `json:"u"`
}

type ReconnectionResponse struct {
	Data []byte `json:"d"`
}

type OutTimeRequest struct {
	UserId uint64 `json:"u"`
}
