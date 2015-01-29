package main

import (
	"errors"
	"gooo/encoder"
	authprot "service/auth/protocol"
	"sync"
)

type Offline struct {
	lock   sync.Mutex
	freeze map[uint64][]byte
	online *Onlines
}

func NewOffline(size uint64) *Offline {
	return &Offline{
		freeze: make(map[uint64][]byte, size),
		online: NewOnlines(size),
	}
}

func (s *Offline) Interrupt(args authprot.InterruptRequest, reply *authprot.InterruptResponse) error {
	var d struct {
		UserId uint64
	}
	encoder.Decode(args.Data, &d)
	if d.UserId == 0 {
		return errors.New("Interrupt index is 0")
	}
	*reply = authprot.InterruptResponse{
		UserId: d.UserId,
	}
	s.online.Del(d.UserId)
	s.lock.Lock()
	defer s.lock.Unlock()
	s.freeze[d.UserId] = args.Data
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
