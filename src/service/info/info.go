package main

import (
	//"errors"
	"gooo/configs"
	"gooo/encoder"
	"gooo/protocol"
	//"gooo/router"
	//infoprtc "service/info/protocol"
	//connprtc "service/connect/protocol"
)

type Info struct {
	conf     *configs.Configs
	status   *Status
	sizechan int
	//callconn *router.CallServer
	//name     string
}

func NewInfo(m *Status) *Info {
	r := Info{
		status:   m,
		conf:     m.Conf,
		sizechan: len(m.Conf.Sc["Chan"]),
		//name:   m.Conf.Sc[configs.Type][configs.Id].Name,
	}
	//r.callconn = router.NewCallServer("Connect", m.Conf)
	return &r
}

func (r *Info) Info(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var d struct {
		Auth uint32 `json:"auth"`
	}
	encoder.Decode(args.Session.Data, &d)
	//if d.Auth != r.status.Conf.St.Login {
	//	return errors.New("Attained")
	//}

	*reply = protocol.RpcResponse{
		Data: &map[string]interface{}{
			"auth": (d.Auth |
				r.status.Conf.St.NoLogin |
				r.status.Conf.St.NoRoom |
				r.status.Conf.St.NoGame),
		},
	}
	return nil
}
