package main

import (
	"gooo/configs"
	"gooo/connser"
	"gooo/helper"
	"service/connect/handel"
	"service/connect/iorange"
	connprot "service/connect/protocol"
)

type Connect struct {
	hand   *handel.Handel
	conf   *configs.Configs
	status *Status
	//auth  *handel.Auth
}

func NewConnect(m *Status) *Connect {
	s := Connect{
		status: m,
		conf:   m.Conf,
		hand:   handel.NewHandel(m.Conf, 1024),
	}
	port := helper.GetPort(s.conf.Sc[configs.Type][configs.Id].ClientPort)
	helper.EchoPublicPortInfo(configs.Name, port)
	ser := connser.NewServer(s.hand, iorange.NewIORange(1024))
	go ser.StartTCP(port)
	return &s
}

func (s *Connect) Send(args connprot.SendRequest, reply *connprot.SendResponse) error {
	for _, v := range args.Clients {
		c := s.hand.Session.Get(v)
		if c != nil {
			c.Conn().Write(append([]byte{255, 255, 255, 255}, args.Data...))
		}
	}
	return nil
}

func (s *Connect) SetOnline(args connprot.SetOnlineRequest, reply *connprot.SetOnlineResponse) error {
	if args.Online {
		s.hand.Online.Set(args.UserId, true)
	} else {
		s.hand.Online.Del(args.UserId)
	}
	return nil
}

func (s *Connect) GetOnline(args connprot.GetOnlineRequest, reply *connprot.GetOnlineResponse) error {
	*reply = connprot.GetOnlineResponse{
		Online: s.hand.Online.Get(args.UserId),
	}
	return nil
}

func (s *Connect) SetSession(args connprot.SetSessionRequest, reply *connprot.SetSessionResponse) error {
	obj := s.hand.Session.Get(args.Id)
	obj.Data = args.Data
	obj.Unlock()
	return nil
}

func (s *Connect) GetSession(args connprot.GetSessionRequest, reply *connprot.GetSessionResponse) error {
	obj := s.hand.Session.Get(args.Id)
	obj.Lock()
	*reply = connprot.GetSessionResponse{
		Data: obj.Data,
	}
	return nil
}

//func (s *Connect) Join(args connprot.JoinRequest, reply *connprot.JoinResponse) error {
//	defer helper.Recover()
//	id := fmt.Sprintln(s.uniq())
//	reg := []byte(fmt.Sprintf("%s %s", s.caddr, id))
//	s.auth.Register(id)
//	*reply = connprot.JoinResponse{
//		Sum:      s.hand.Session.Len(),
//		Response: reg,
//	}
//	return nil
//}
