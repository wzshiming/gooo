package protocol

import (
	//"time"
	"gooo/session"
)

type RpcRequest struct {
	Request []byte           `json:"r"`
	Session *session.Session `json:"s"`
}

type RpcResponse struct {
	Error    string                  `json:"e"`
	Data     *map[string]interface{} `json:"d"`
	Coverage []byte                  `json:"c"`
	Response []byte                  `json:"r"`
}
