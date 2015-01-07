package main

import (
	"gooo/protocol"
	"gooo/encoder"

	"gooo/configs"
	"gooo/helper"
)

type Information struct {
	conf     *configs.Configs
	versions []string
}

func NewInformation() (info *Information) {
	info = &Information{}

	return
}

func (i *Information) Versions(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var p protocol.VersionsRequest
	encoder.Decode(args.Request, &p)
	var r protocol.VersionsResponse

	res, _ := encoder.Encode(r)
	*reply = protocol.RpcResponse{
		Error:    0,
		Response: res,
	}
	//log.Println(t)
	return nil
}

func (r *Information) Init(args protocol.InitRequest, reply *int) error {
	if args.State == 1 {
		r.conf = configs.NewConfigs(&args.Conf)
		r.conf.StartConnect()
	}
	return nil
}

func main() {
	h := helper.NewHandeln()
	c := NewInformation()
	h.Register(c)
	h.RegisterName("Status", c)
	helper.EchoPortInfo(configs.Name, configs.Port)
	h.Start(configs.Port)
}
