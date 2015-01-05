package protocol

import (
//"time"
)

type InitRequest struct {
	Conf  map[string][]byte `json: conf`
	State int               `json: state`
}
