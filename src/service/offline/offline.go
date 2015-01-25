package main

import (
	"errors"
	"gooo/encoder"
	offlprot "service/offline/protocol"
	"sync"
)

type Offline struct {
	lock   sync.Mutex
	freeze map[uint64][]byte
}

func NewOffline(size uint64) *Offline {
	return &Offline{
		freeze: make(map[uint64][]byte, size),
	}
}

func (s *Offline) Interrupt(args offlprot.InterruptRequest, reply *offlprot.InterruptResponse) error {
	var d struct {
		UserId uint64
	}
	encoder.Decode(args.Data, &d)
	if d.UserId == 0 {
		return errors.New("Interrupt index is 0")
	}
	*reply = offlprot.InterruptResponse{
		UserId: d.UserId,
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	s.freeze[d.UserId] = args.Data
	return nil
}

func (s *Offline) Reconnection(args offlprot.ReconnectionRequest, reply *offlprot.ReconnectionResponse) error {
	if args.UserId == 0 {
		return errors.New("Reconnection index is 0")
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	b := s.freeze[args.UserId]
	if b != nil {
		delete(s.freeze, args.UserId)
	} else {
		return errors.New("Not Reconnection")
	}
	*reply = offlprot.ReconnectionResponse{
		Data: b,
	}
	return nil
}

func (s *Offline) OutTime(args offlprot.OutTimeRequest, reply *int) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.freeze, args.UserId)
	return nil
}
