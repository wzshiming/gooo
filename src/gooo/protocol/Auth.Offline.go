package protocol

import (
	"gooo"
)

type InterruptRequest struct {
	Data []byte `json:"d"`
}

type InterruptResponse int

type ReconnectionRequest struct {
	UserId uint64      `json:"u"`
	Unique gooo.Unique `json:"s"`
}

type ReconnectionResponse struct {
	Data []byte `json:"d"`
}

type GetOnlineRequest struct {
	UserId uint64 `json:"u"`
}

type GetOnlineResponse struct {
	Online bool        `json:"o"`
	Unique gooo.Unique `json:"s"`
}

type OutTimeRequest struct {
	UserId uint64 `json:"u"`
}
