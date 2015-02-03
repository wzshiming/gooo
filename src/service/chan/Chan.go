package main

import (
	"errors"
	"gooo/configs"
	"gooo/encoder"
	"gooo/protocol"
	chanprtc "service/chan/protocol"
	"service/chan/room"

	//connprtc "service/connect/protocol"
)

type Chan struct {
	conf   *configs.Configs
	status *Status
	rooms  *room.GameRooms
}

func NewChan(m *Status) *Chan {
	r := Chan{
		status: m,
		conf:   m.Conf,
		rooms:  room.NewGameRooms(100),
	}

	return &r
}

func (r *Chan) Create(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var p chanprtc.CreateRequest
	encoder.Decode(args.Request, &p)
	var d struct {
		Flag     uint32 `json:"flag"`
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
			"RoomId": b.RoomId,
			"SeatId": b.SeatId,
			"flag":   d.Flag | configs.FlagRoom,
		},
	}
	return nil
}

func (r *Chan) Join(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var p chanprtc.JoinRequest
	encoder.Decode(args.Request, &p)
	var d struct {
		Flag     uint32 `json:"flag"`
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
			"RoomId": b.RoomId,
			"SeatId": b.SeatId,
			"flag":   d.Flag | configs.FlagRoom,
		},
	}
	return nil
}
