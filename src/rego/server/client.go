package server

import (
	"net/rpc"
	"rego"
)

type Client struct {
	rpc.Client
}

func NewClient(addr string) *Client {
	c, err := rpc.Dial("tcp", addr)
	if err != nil {
		rego.ERR(err)
		return nil
	}
	return &Client{
		Client: *c,
	}
}
