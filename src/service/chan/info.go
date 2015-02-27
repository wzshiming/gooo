package main

import (
	"gooo"
	"gooo/protocol"
	"time"
)

type Info struct {
	gooo.Methods
	conf          *gooo.Configs
	status        *Status
	rooms         *gooo.GameRooms
	name          string
	nameResponse  []byte
	roomsResponse []byte
	tmp           []gooo.GameRoom
}

func NewInfo(m *Status, rooms *gooo.GameRooms) *Info {
	r := Info{
		status: m,
		conf:   m.Conf,
		name:   m.Conf.Self().Name,
		rooms:  rooms,
	}
	r.Methods = *gooo.NewMethods(
		gooo.FlagChan,
		gooo.FlagNone,
		"Name",
		"Rooms",
	)
	r.nameResponse, _ = gooo.Encode(protocol.NameResponse{
		Name: r.name,
	})
	go func() {
		for {
			r.tmp = r.rooms.List()
			r.roomsResponse, _ = gooo.Encode(r.tmp)
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
