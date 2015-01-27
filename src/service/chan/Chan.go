package main

import (
	//"errors"
	"gooo/configs"
	"gooo/encoder"
	"gooo/protocol"
	"gooo/router"
	chanprtc "service/chan/protocol"
	//connprtc "service/connect/protocol"
)

type Chan struct {
	conf     *configs.Configs
	status   *Status
	callconn *router.CallServer
	name     string
}

func NewChan(m *Status) *Chan {
	r := Chan{
		status: m,
		conf:   m.Conf,
		name:   m.Conf.Sc[configs.Type][configs.Id].Name,
	}
	r.callconn = router.NewCallServer("Connect", m.Conf)
	return &r
}

func (r *Chan) Name(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	p := chanprtc.NameResponse{
		Name: r.name,
	}
	b, _ := encoder.Encode(p)
	*reply = protocol.RpcResponse{
		Response: b,
	}
	return nil
}
