package protocol

import ()

type InfoRequest struct {
	Get bool `json:"g"`
}

type InfoResponse struct {
	None bool `json:"n"`
}
