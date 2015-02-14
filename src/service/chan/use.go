package main

import (
	"errors"
	"gooo"
	"gooo/protocol"
)

type Use struct {
	gooo.Methods
	conf     *gooo.Configs
	status   *Status
	sizechan int
	//callconn *router.CallServer
	//name     string
}

func NewUse(m *Status) *Use {
	r := Use{
		status:   m,
		conf:     m.Conf,
		sizechan: m.Conf.TypeSize("Chan"),
	}
	r.Methods = *gooo.NewMethods(
		gooo.FlagNone,
		gooo.FlagChan|gooo.FlagRoom|gooo.FlagGame,
		"Chan",
	)
	return &r
}

func (r *Use) Chan(args gooo.RpcRequest, reply *gooo.RpcResponse) error {
	var p protocol.UseChanRequest
	gooo.Decode(args.Request, &p)

	if r.sizechan < p.Use {
		return errors.New("Out of range")
	}

	var d struct {
		Flag     uint32 `json:"flag"`
		Language string
		UseChan  uint32 `json:"useChan"`
	}
	gooo.Decode(args.Session.Data, &d)

	if d.UseChan != 0 {
		return errors.New("Have chosen")
	}

	*reply = gooo.RpcResponse{
		Data: &map[string]interface{}{
			"useChan": p.Use,
			"flag":    d.Flag | gooo.FlagChan,
		},
	}
	return nil
}
