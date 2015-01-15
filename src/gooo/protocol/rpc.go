package protocol

import (
	//"time"
	"gooo/session"
)

type RpcRequest struct {
	Request []byte           `json:"r"`
	Session *session.Session `json:"s"`
	Id      uint             `json:"i"`
}

type RpcResponse struct {
	Error    string                  `json:"e"`
	Data     *map[string]interface{} `json:"d"`
	Response []byte                  `json:"r"`
}
