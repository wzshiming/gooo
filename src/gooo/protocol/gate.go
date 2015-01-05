package protocol

import (

)

type GateRequest struct {
	Id      uint             `json: id`
}


type GateResponse struct {
    Sum      int             `json:sum`
    Response []byte          `json:response`
}