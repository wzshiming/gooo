package protocol

import (
//"time"
)

type SendRequest struct {
	Clients []uint `json:c`
	Data    []byte `json:d`
}

type SendResponse struct {
	Error int `json:e`
}
