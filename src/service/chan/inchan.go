package main

import (
	"errors"
	"gooo"
	"gooo/protocol"
)

type InChan struct {
	conf   *gooo.Configs
	status *Status
	rooms  *gooo.GameRooms
}

func NewInChan(m *Status, rooms *gooo.GameRooms) *InChan {
	r := InChan{
		status: m,
		conf:   m.Conf,
		rooms:  rooms,
	}
	return &r
}

func (r *InChan) InterruptLeave(args protocol.InterruptLeaveRequest, reply *protocol.InterruptLeaveResponse) error {
	r.rooms.Leave(args.RoomId, args.SeatId)
	return nil
}

func (r *InChan) Leave(args gooo.RpcRequest, reply *gooo.RpcResponse) error {
	var d struct {
		Flag     uint32 `json:"flag"`
		Language string
		UserId   uint64 `json:"userId"`
		RoomId   int
		SeatId   int
	}
	gooo.Decode(args.Session.Data, &d)

	if (d.Flag & gooo.FlagGame) != 0 {
		return errors.New("In the game")
	}

	r.rooms.Leave(d.RoomId, d.SeatId)

	*reply = gooo.RpcResponse{
		Data: &map[string]interface{}{
			"RoomId": nil,
			"SeatId": nil,
			"flag":   d.Flag & ^gooo.FlagRoom,
		},
	}
	return nil
}

func (r *InChan) Play(args gooo.RpcRequest, reply *gooo.RpcResponse) error {
	var d struct {
		Flag     uint32 `json:"flag"`
		Language string
		UserId   uint64 `json:"userId"`
		RoomId   int
		SeatId   int
	}
	gooo.Decode(args.Session.Data, &d)

	if r.rooms.Room(d.RoomId).Master != d.SeatId {
		return errors.New("You are not master")
	}

	r.rooms.Play(d.RoomId)
	*reply = gooo.RpcResponse{
		Data: &map[string]interface{}{
			"flag": d.Flag | gooo.FlagGame,
		},
	}
	return nil
}
