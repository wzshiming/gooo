package main

import (
	"gooo"
	"gooo/protocol"
	"time"
)

type Info struct {
	conf          *gooo.Configs
	status        *Status
	name          string
	nameResponse  []byte
	roomsResponse []byte
}

func NewInfo(m *Status) *Info {
	r := Info{
		status: m,
		conf:   m.Conf,
		name:   m.Conf.Sc[gooo.Type][gooo.Id].Name,
	}
	r.nameResponse, _ = gooo.Encode(protocol.NameResponse{
		Name: r.name,
	})
	go func() {
		for {
			r.roomsResponse, _ = gooo.Encode(r.status.ServiceChan.rooms.List())
			//fmt.Println(string(r.roomsResponse))
			time.Sleep(time.Second)
		}
	}()
	return &r
}

func (r *Info) Name(args gooo.RpcRequest, reply *gooo.RpcResponse) error {
	*reply = gooo.RpcResponse{
		Response: r.nameResponse,
	}
	return nil
}

func (r *Info) Rooms(args gooo.RpcRequest, reply *gooo.RpcResponse) error {
	*reply = gooo.RpcResponse{
		Response: r.roomsResponse,
	}
	return nil
}
