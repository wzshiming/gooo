package main

import (
	"connect/iorange"
	"gate/handel"
	"gooo/configs"
	"gooo/connser"
	"gooo/helper"
	"gooo/protocol"
	"log"
)

type Gate struct {
	handel *handel.Handel
}

func NewGate() *Gate {
	return &Gate{}
}

func (s *Gate) Init(args protocol.InitRequest, reply *int) error {
	if args.State == 1 {
		conf := configs.NewConfigs(&args.Conf)
		conf.StartConnect()
		port := helper.GetPort(conf.Sc.Devel[configs.Type][configs.Id].ClientPort)
		helper.EchoPublicPortInfo(configs.Name, port)
		s.handel = handel.NewHandel(conf)
		ser := connser.NewServer(s.handel, iorange.NewIORange(1024))
		go ser.StartTCP(port)
	}
	return nil
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	h := helper.NewHandeln()
	c := NewGate()
	h.Register(c)
	h.RegisterName("Status", c)
	helper.EchoPortInfo(configs.Name, configs.Port)
	h.Start(configs.Port)
}
