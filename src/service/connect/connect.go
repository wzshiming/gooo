package main

import (
	"fmt"
	"gooo/configs"
	"gooo/connser"
	"gooo/helper"
	"gooo/protocol"
	"log"
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

		ser := connser.NewServer(s.auth, iorange.NewIORange(1024))
		go ser.StartTCP(port)

		s.caddr = fmt.Sprintf("%s%s", configs.IP, port)
	}
	return nil
}

func (s *Connect) Join(args protocol.GateRequest, reply *protocol.GateResponse) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Join error: ", err)
		}
	}()
	id := []byte(fmt.Sprintf("%d", args.Id))
	reg := []byte(fmt.Sprintf("%s %d", s.caddr, args.Id))
	*reply = protocol.GateResponse{
		Sum:      s.hand.Len(),
		Response: reg,
	}

	s.auth.Register(id)
	return nil
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	h := helper.NewHandeln()
	c := NewConnect()
	h.Register(c)
	h.RegisterName("Status", c)
	helper.EchoPortInfo(configs.Name, configs.Port)
	h.Start(configs.Port)
}
