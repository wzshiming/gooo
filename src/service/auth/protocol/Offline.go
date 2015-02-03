package protocol

import (
	"gooo/session"
)

type InterruptRequest struct {
	Data []byte `json:"d"`
}

type InterruptResponse int

type ReconnectionRequest struct {
	UserId uint64         `json:"u"`
	Unique session.Unique `json:"s"`
}

type ReconnectionResponse struct {
	Data []byte `json:"d"`
}

type GetOnlineRequest struct {
	UserId uint64 `json:"u"`
}

type GetOnlineResponse struct {
	Online bool           `json:"o"`
	Unique session.Unique `json:"s"`
}

type OutTimeRequest struct {
	UserId uint64 `json:"u"`
}
