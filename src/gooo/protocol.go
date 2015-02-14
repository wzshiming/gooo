package gooo

import ()

type RpcRequest struct {
	Request []byte   `json:"r"`
	Session *Session `json:"s"`
}

type RpcResponse struct {
	Error    string                  `json:"e"`
	Data     *map[string]interface{} `json:"d"`
	Coverage []byte                  `json:"c"`
	Response []byte                  `json:"r"`
}

type InitRequest struct {
	Conf  Configs `json:"conf"`
	State int     `json:"state"`
}

type InitResponse int

type MethodsRequest int

type MethodsResponse struct {
	Allow   uint32   `json:"allow"`
	Unallow uint32   `json:"unallow"`
	Method  []string `json:"method"`
}
