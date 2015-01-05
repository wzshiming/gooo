package protocol

import (
//"time"
)

type RandRequest struct {
	Size int `json:s`
}

type RandResponse struct {
	Rands []int `json:r`
}

type SpacingRequest struct {
	Size  int `json:s`
	Space int `json:p`
}
