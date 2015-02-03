package main

import (
	//"fmt"
	"gooo/configs"
	"gooo/encoder"
	"gooo/protocol"
	chanprtc "service/chan/protocol"
	"time"
	//connprtc "service/connect/protocol"
)

type Info struct {
	conf          *configs.Configs
	status        *Status
	name          string
	nameResponse  []byte
	roomsResponse []byte
}

func NewInfo(m *Status) *Info {
	r := Info{
		status: m,
		conf:   m.Conf,
		name:   m.Conf.Sc[configs.Type][configs.Id].Name,
	}
	r.nameResponse, _ = encoder.Encode(chanprtc.NameResponse{
		Name: r.name,
	})
	go func() {
		for {
			r.roomsResponse, _ = encoder.Encode(r.status.ServiceChan.rooms.List())
			//fmt.Println(string(r.roomsResponse))
			time.Sleep(time.Second)
		}
	}()
	return &r
}

func (r *Info) Name(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	*reply = protocol.RpcResponse{
		Response: r.nameResponse,
	}
	return nil
}

func (r *Info) Rooms(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	*reply = protocol.RpcResponse{
		Response: r.roomsResponse,
	}
	return nil
}
