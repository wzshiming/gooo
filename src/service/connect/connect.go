package main

import (
	"fmt"
	"gooo/configs"
	"gooo/connser"
	"gooo/helper"
	"gooo/protocol"
	"gooo/session"
	"service/connect/handel"
	"service/connect/iorange"
	connprot "service/connect/protocol"
)

type Connect struct {
	hand  *handel.Handel
	auth  *handel.Auth
	caddr string
	uniq  func() uint
}

func NewConnect() *Connect {
	return &Connect{
		uniq: session.NewUniqUint(),
	}
}

func (s *Connect) Init(args protocol.InitRequest, reply *int) error {
	if args.State == 1 {
		conf := &args.Conf
		conf.StartConnect()
		port := helper.GetPort(conf.Sc[configs.Type][configs.Id].ClientPort)
		helper.EchoPublicPortInfo(configs.Name, port)
		s.hand = handel.NewHandel(conf)
		s.auth = handel.NewAuth(s.hand)

		ser := connser.NewServer(s.auth, iorange.NewIORange(1024))
		go ser.StartTCP(port)

		s.caddr = fmt.Sprintf("%s%s", configs.IP, port)
	}
	return nil
}

func (s *Connect) Send(args connprot.SendRequest, reply *connprot.SendResponse) error {
	for _, v := range args.Clients {
		c := s.hand.Get(v)
		if c != nil {
			c.Conn().Write(args.Data)
		}
	}
	*reply = connprot.SendResponse{
		Error: 0,
	}
	return nil
}

func (s *Connect) Join(args connprot.JoinRequest, reply *connprot.JoinResponse) error {
	defer helper.Recover()
	id := fmt.Sprintln(s.uniq())
	reg := []byte(fmt.Sprintf("%s %s", s.caddr, id))
	s.auth.Register(id)
	*reply = connprot.JoinResponse{
		Sum:      s.hand.Len(),
		Response: reg,
	}
	return nil
}

func main() {
	defer helper.Recover()

	h := helper.NewHandeln()
	c := NewConnect()
	h.Register(c)

	helper.EchoPortInfo(configs.Name, configs.Port)
	h.Start(configs.Port)
}
