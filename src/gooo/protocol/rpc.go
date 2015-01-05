package protocol

import (
	//"time"
	"gooo/session"
)

type RpcRequest struct {
	Request []byte           `json: request`
	Session *session.Session `json: session`
	Id      uint             `json: id`
}

type RpcResponse struct {
	Error    int                     `json:error`
	Data     *map[string]interface{} `json:data`
	Response []byte                  `json:response`
}

type OkResponse struct {
	Error int `json:error`
}
