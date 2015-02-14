package main

import (
	"errors"
	"gooo"
	"gooo/protocol"
	"sync"
)

type Offline struct {
	lock     sync.Mutex
	freeze   map[uint64][]byte
	online   *gooo.Onlines
	callchan *gooo.CallServer
}

func NewOffline(m *Status, size uint64) *Offline {
	r := Offline{
		freeze:   make(map[uint64][]byte, size),
		online:   gooo.NewOnlines(size),
		callchan: gooo.NewCallServer("Chan", m.Conf),
	}

	return &r
}

func (s *Offline) Interrupt(args protocol.InterruptRequest, reply *protocol.InterruptResponse) error {
	var d struct {
		UserId  uint64 `json:"userId"`
		RoomId  int
		SeatId  int
		UseChan int    `json:"useChan"`
		Flag    uint32 `json:"flag"`
	}

	gooo.Decode(args.Data, &d)
	if d.UserId == 0 {
		return errors.New("Interrupt index is 0")
	}

	s.online.Del(d.UserId)
	if 0 != (d.Flag & gooo.FlagGame) {
		s.lock.Lock()
		defer s.lock.Unlock()
		s.freeze[d.UserId] = args.Data
	} else if 0 != (d.Flag&gooo.FlagChan) && 0 != (d.Flag&gooo.FlagRoom) {
		var p protocol.InterruptLeaveResponse
		s.callchan.CallBy(d.UseChan, "InChan.InterruptLeave", protocol.InterruptLeaveRequest{
			RoomId: d.RoomId,
			SeatId: d.SeatId,
		}, &p)
	}
	return nil
}

func (s *Offline) GetOnline(args protocol.GetOnlineRequest, reply *protocol.GetOnlineResponse) error {

	b := s.online.Get(args.UserId)
	if b != nil {
		*reply = protocol.GetOnlineResponse{
			Online: false,
		}
	} else {
		*reply = protocol.GetOnlineResponse{
			Online: true,
			Unique: *b,
		}
	}
	return nil
}

func (s *Offline) Reconnection(args protocol.ReconnectionRequest, reply *protocol.ReconnectionResponse) error {

	if s.online.Get(args.UserId) != nil {
		return errors.New("It has been used")
	}
	s.online.Set(args.UserId, &args.Unique)

	s.lock.Lock()
	defer s.lock.Unlock()
	b := s.freeze[args.UserId]
	if b != nil {
		delete(s.freeze, args.UserId)
	}

	*reply = protocol.ReconnectionResponse{
		Data: b,
	}
	return nil
}

func (s *Offline) OutTime(args protocol.OutTimeRequest, reply *int) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.freeze, args.UserId)
	return nil
}
