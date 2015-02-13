package main

import (
	"errors"
	"gooo"
	"gooo/protocol"
)

type Chan struct {
	conf   *gooo.Configs
	status *Status
	rooms  *gooo.GameRooms
}

func NewChan(m *Status, rooms *gooo.GameRooms) *Chan {
	r := Chan{
		status: m,
		conf:   m.Conf,
		rooms:  rooms,
	}

	return &r
}

func (r *Chan) Create(args gooo.RpcRequest, reply *gooo.RpcResponse) error {
	var p protocol.CreateRequest
	gooo.Decode(args.Request, &p)
	var d struct {
		Flag     uint32 `json:"flag"`
		Language string
		UserId   uint64 `json:"userId"`
	}
	gooo.Decode(args.Session.Data, &d)

	room, seat := r.rooms.Create(p.Size, d.UserId)

	if seat < 0 {
		return errors.New("Unable to create the room")
	}

	b := protocol.CreateResponse{
		RoomId: room,
		SeatId: seat,
	}

	s, _ := gooo.Encode(b)
	*reply = gooo.RpcResponse{
		Response: s,
		Data: &map[string]interface{}{
			"RoomId": b.RoomId,
			"SeatId": b.SeatId,
			"flag":   d.Flag | gooo.FlagRoom,
		},
	}
	return nil
}

func (r *Chan) Join(args gooo.RpcRequest, reply *gooo.RpcResponse) error {
	var p protocol.JoinRequest
	gooo.Decode(args.Request, &p)
	var d struct {
		Flag     uint32 `json:"flag"`
		Language string
		UserId   uint64 `json:"userId"`
	}
	gooo.Decode(args.Session.Data, &d)

	seat := r.rooms.Join(p.RoomId, d.UserId)

	if seat < 0 {
		return errors.New("Unable to join the room")
	}

	b := protocol.JoinResponse{
		RoomId: p.RoomId,
		SeatId: seat,
	}

	s, _ := gooo.Encode(b)
	*reply = gooo.RpcResponse{
		Response: s,
		Data: &map[string]interface{}{
			"RoomId": b.RoomId,
			"SeatId": b.SeatId,
			"flag":   d.Flag | gooo.FlagRoom,
		},
	}
	return nil
}
