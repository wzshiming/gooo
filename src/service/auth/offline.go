package main

import (
	"errors"
	"gooo/configs"
	"gooo/encoder"
	"gooo/router"
	"service/auth/online"
	authprot "service/auth/protocol"
	chanprot "service/chan/protocol"
	"sync"
)

type Offline struct {
	status   *Status
	lock     sync.Mutex
	freeze   map[uint64][]byte
	online   *online.Onlines
	callchan *router.CallServer
}

func NewOffline(m *Status, size uint64) *Offline {
	return &Offline{
		status:   m,
		freeze:   make(map[uint64][]byte, size),
		online:   online.NewOnlines(size),
		callchan: router.NewCallServer("Chan", m.Conf),
	}
}

func (s *Offline) Interrupt(args authprot.InterruptRequest, reply *authprot.InterruptResponse) error {
	var d struct {
		UserId  uint64 `json:"userId"`
		RoomId  int
		SeatId  int
		UseChan int    `json:"useChan"`
		Flag    uint32 `json:"flag"`
	}
	encoder.Decode(args.Data, &d)
	if d.UserId == 0 {
		return errors.New("Interrupt index is 0")
	}

	s.online.Del(d.UserId)
	if 0 != (d.Flag & configs.FlagGame) {
		s.lock.Lock()
		defer s.lock.Unlock()
		s.freeze[d.UserId] = args.Data
	} else if 0 != (d.Flag & configs.FlagRoom) {
		var p chanprot.InterruptResponse
		s.callchan.CallBy(d.UseChan-1, "InChan.Interrupt", chanprot.InterruptRequest{
			RoomId: d.RoomId,
			SeatId: d.SeatId,
		}, &p)
	}
	return nil
}

func (s *Offline) GetOnline(args authprot.GetOnlineRequest, reply *authprot.GetOnlineResponse) error {

	b := s.online.Get(args.UserId)
	if b != nil {
		*reply = authprot.GetOnlineResponse{
			Online: false,
		}
	} else {
		*reply = authprot.GetOnlineResponse{
			Online: true,
			Unique: *b,
		}
	}
	return nil
}

func (s *Offline) Reconnection(args authprot.ReconnectionRequest, reply *authprot.ReconnectionResponse) error {

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

	*reply = authprot.ReconnectionResponse{
		Data: b,
	}
	return nil
}

func (s *Offline) OutTime(args authprot.OutTimeRequest, reply *int) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.freeze, args.UserId)
	return nil
}
