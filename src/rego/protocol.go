package rego

import (
//"net/http"
)

type Request struct {
	Request *EncodeStream `json:"r"`
	Session *Session      `json:"s"`
}

type Response struct {
	Error    string        `json:"e"`
	Data     *EncodeStream `json:"d"`
	Coverage *EncodeStream `json:"c"`
	Response *EncodeStream `json:"r"`
}

type MethodsRequest int

type MethodsResponse struct {
	Allow   uint32   `json:"allow"`
	Unallow uint32   `json:"unallow"`
	Method  []string `json:"method"`
}
