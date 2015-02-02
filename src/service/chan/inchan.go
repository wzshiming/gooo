package main

import (
	"errors"
	"gooo/configs"
	"gooo/encoder"
	"gooo/protocol"
	chanprtc "service/chan/protocol"

	//connprtc "service/connect/protocol"
)

type InChan struct {
	conf   *configs.Configs
	status *Status
	rooms  *chanprtc.GameRooms
}

func NewInChan(m *Status) *InChan {
	r := InChan{
		status: m,
		conf:   m.Conf,
		rooms:  chanprtc.NewGameRooms(100),
	}
	return &r
}

func (r *InChan) Leave(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var d struct {
		Flag     uint32 `json:"flag"`
		Language string
		UserId   uint64 `json:"userId"`
		RoomId   int
		SeatId   int
	}
	encoder.Decode(args.Session.Data, &d)

	if (d.Flag & configs.FlagGame) != 0 {
		return errors.New("In the game")
	}

	r.rooms.Leave(d.RoomId, d.SeatId)

	*reply = protocol.RpcResponse{
		Data: &map[string]interface{}{
			"RoomId": nil,
			"SeatId": nil,
			"flag":   d.Flag & ^configs.FlagRoom,
		},
	}
	return nil
}
