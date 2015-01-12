package main

import (
	"gooo/configs"
	"gooo/connser"
	"gooo/helper"
	"gooo/protocol"
	"service/connect/iorange"
	"service/gate/handel"
)

type Gate struct {
	handel *handel.Handel
}

func NewGate() *Gate {
	return &Gate{}
}

func (s *Gate) Init(args protocol.InitRequest, reply *int) error {
	if args.State == 1 {
		conf := &args.Conf
		conf.StartConnect()
		port := helper.GetPort(conf.Sc[configs.Type][configs.Id].ClientPort)
		helper.EchoPublicPortInfo(configs.Name, port)
		s.handel = handel.NewHandel(conf)
		ser := connser.NewServer(s.handel, iorange.NewIORange(1024))
		go ser.StartTCP(port)
	}
	return nil
}

func main() {
	defer helper.Recover()

	h := helper.NewHandeln()
	c := NewGate()
	h.Register(c)
	h.Register(helper.NewStatus(h))
	helper.EchoPortInfo(configs.Name, configs.Port)
	h.Start(configs.Port)
}
