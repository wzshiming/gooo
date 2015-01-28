package main

import (
	"errors"
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
	rooms    *chanprtc.GameRooms
}

func NewChan(m *Status) *Chan {
	r := Chan{
		status: m,
		conf:   m.Conf,
		rooms:  chanprtc.NewGameRooms(100),
	}
	r.callconn = router.NewCallServer("Connect", m.Conf)
	return &r
}

func (r *Chan) Create(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var p chanprtc.CreateRequest
	encoder.Decode(args.Request, &p)
	var d struct {
		Auth     uint32 `json:"auth"`
		Language string
		UserId   uint64 `json:"userId"`
	}
	encoder.Decode(args.Session.Data, &d)

	room, seat := r.rooms.Create(p.Size, d.UserId)
	b := chanprtc.CreateResponse{
		RoomId: room,
		SeatId: seat,
	}
	if b.RoomId < 0 {
		return errors.New("Unable to create the room")
	}

	s, _ := encoder.Encode(b)
	*reply = protocol.RpcResponse{
		Response: s,
		Data: &map[string]interface{}{
			"roomId": s,
			"auth":   ((d.Auth & ^r.status.Conf.St.NoRoom) | r.status.Conf.St.InRoom),
		},
	}
	return nil
}

func (r *Chan) Join(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var p chanprtc.JoinRequest
	encoder.Decode(args.Request, &p)
	var d struct {
		Auth     uint32 `json:"auth"`
		Language string
		UserId   uint64 `json:"userId"`
	}
	encoder.Decode(args.Session.Data, &d)

	seat := r.rooms.Join(p.RoomId, d.UserId)
	b := chanprtc.JoinResponse{
		RoomId: p.RoomId,
		SeatId: seat,
	}
	if b.RoomId < 0 {
		return errors.New("Unable to join the room")
	}

	s, _ := encoder.Encode(b)
	*reply = protocol.RpcResponse{
		Response: s,
		Data: &map[string]interface{}{
			"roomId": s,
			"auth":   ((d.Auth & ^r.status.Conf.St.NoRoom) | r.status.Conf.St.InRoom),
		},
	}
	return nil
}
