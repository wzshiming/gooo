package main

import (
	"gooo"
	"gooo/protocol"
)

type Connect struct {
	conf   *gooo.Configs
	status *Status
}

func NewConnect(m *Status) *Connect {
	s := Connect{
		status: m,
		conf:   m.Conf,
	}
	return &s
}

func (s *Connect) Send(args protocol.SendRequest, reply *protocol.SendResponse) error {
	for _, v := range args.Clients {
		c := s.status.CilentHandeln.Session.Get(v)
		if c != nil {
			c.Conn().Write(append([]byte{255, 255, 255, 255}, args.Data...))
		}
	}
	return nil
}

func (s *Connect) SetSession(args protocol.SetSessionRequest, reply *protocol.SetSessionResponse) error {
	obj := s.status.CilentHandeln.Session.Get(args.Id)
	obj.Data = args.Data
	obj.Unlock()
	return nil
}

func (s *Connect) GetSession(args protocol.GetSessionRequest, reply *protocol.GetSessionResponse) error {
	obj := s.status.CilentHandeln.Session.Get(args.Id)
	obj.Lock()
	*reply = protocol.GetSessionResponse{
		Data: obj.Data,
	}
	return nil
}
