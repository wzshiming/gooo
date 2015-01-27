package main

import (
	"errors"
	"gooo/configs"
	"gooo/encoder"
	"gooo/protocol"
	//"gooo/router"
	infoprtc "service/info/protocol"
	//connprtc "service/connect/protocol"
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
		//name:   m.Conf.Sc[configs.Type][configs.Id].Name,
	}
	//r.callconn = router.NewCallServer("Connect", m.Conf)
	return &r
}

func (r *Use) Chan(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var p infoprtc.UseChanRequest
	encoder.Decode(args.Request, &p)

	if r.sizechan < p.Use {
		return errors.New("out of range")
	}

	var d struct {
		Auth uint32 `json:"auth"`
	}
	encoder.Decode(args.Session.Data, &d)

	*reply = protocol.RpcResponse{
		Data: &map[string]interface{}{
			"useChan": p.Use,
			"auth":    (d.Auth | r.status.Conf.St.InRoom),
		},
	}
	return nil
}
