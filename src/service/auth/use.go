package main

import (
	"errors"
	"gooo/configs"
	"gooo/encoder"
	"gooo/protocol"
	//"gooo/router"
	authprot "service/auth/protocol"
)

type Use struct {
	conf     *configs.Configs
	status   *Status
	sizechan int
	//callconn *router.CallServer
	//name     string
}

func NewUse(m *Status) *Use {
	r := Use{
		status:   m,
		conf:     m.Conf,
		sizechan: len(m.Conf.Sc["Chan"]),
	}
	return &r
}

func (r *Use) Chan(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var p authprot.UseChanRequest
	encoder.Decode(args.Request, &p)

	if r.sizechan < p.Use {
		return errors.New("Out of range")
	}

	var d struct {
		UseChan uint32 `json:"useChan"`
	}

	if d.UseChan != 0 {
		return errors.New("Have chosen")
	}

	encoder.Decode(args.Session.Data, &d)

	*reply = protocol.RpcResponse{
		Data: &map[string]interface{}{
			"useChan": p.Use,
		},
	}
	return nil
}
