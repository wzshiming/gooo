package main

import (
	"gooo/configs"
	connprot "service/connect/protocol"
)

type Connect struct {
	conf   *configs.Configs
	status *Status
	//auth  *handel.Auth
}

func NewConnect(m *Status) *Connect {
	s := Connect{
		status: m,
		conf:   m.Conf,
	}
	return &s
}

func (s *Connect) Send(args connprot.SendRequest, reply *connprot.SendResponse) error {
	for _, v := range args.Clients {
		c := s.status.CilentHandeln.Session.Get(v)
		if c != nil {
			c.Conn().Write(append([]byte{255, 255, 255, 255}, args.Data...))
		}
	}
	return nil
}

func (s *Connect) SetSession(args connprot.SetSessionRequest, reply *connprot.SetSessionResponse) error {
	obj := s.status.CilentHandeln.Session.Get(args.Id)
	obj.Data = args.Data
	obj.Unlock()
	return nil
}

func (s *Connect) GetSession(args connprot.GetSessionRequest, reply *connprot.GetSessionResponse) error {
	obj := s.status.CilentHandeln.Session.Get(args.Id)
	obj.Lock()
	*reply = connprot.GetSessionResponse{
		Data: obj.Data,
	}
	return nil
}
