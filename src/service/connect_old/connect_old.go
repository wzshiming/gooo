package main

import (
	"fmt"
	"gooo/configs"
	"gooo/connser"
	"gooo/helper"
	"gooo/protocol"
	"service/connect/handel"
	"service/connect/iorange"
)

type Connect struct {
	hand  *handel.Handel
	auth  *handel.Auth
	caddr string
}

func NewConnect() *Connect {
	return &Connect{}
}

func (s *Connect) Send(args protocol.SendRequest, reply *protocol.SendResponse) error {
	for _, v := range args.Clients {

		c := s.hand.Get(v)
		if c != nil {
			c.Conn().Write(args.Data)
		}
	}
	*reply = protocol.SendResponse{
		Error: 0,
	}
	return nil
}

func (s *Connect) Init(args protocol.InitRequest, reply *int) error {
	if args.State == 1 {
		conf := configs.NewConfigs(&args.Conf)
		conf.StartConnect()
		port := helper.GetPort(conf.Sc.Devel[configs.Type][configs.Id].ClientPort)
		helper.EchoPublicPortInfo(configs.Name, port)
		s.hand = handel.NewHandel(conf)
		s.auth = handel.NewAuth(s.hand)

		ser := connser.NewServer(s.hand, iorange.NewIORange(1024))
		go ser.StartTCP(port)

		s.caddr = fmt.Sprintf("127.0.0.1%s", port)
	}
	return nil
}

func main() {
	defer helper.Recover()

	h := helper.NewHandeln()
	c := NewConnect()
	h.Register(c)
	h.RegisterName("Status", c)
	helper.EchoPortInfo(configs.Name, configs.Port)
	h.Start(configs.Port)
}
