package handel

import (
	"gooo/configs"
	"gooo/connser"
	"gooo/helper"
	"gooo/protocol"
	"gooo/router"
	"log"
)

type Handel struct {
	helper.HandelInterface
	Server *router.CallServer
}

func NewHandel(conf *configs.Configs) *Handel {
	return &Handel{
		Server: router.NewCallServer("Connect", conf),
	}
}

func (h *Handel) Mess(c *connser.Connect, msg []byte) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Client %v Mess error: %v", c.RemoteAddr(), err)
			helper.RecoverInfo()
		}
	}()

	var reply protocol.GateResponse
	h.Server.Call("Status.Join", protocol.GateRequest{
		Id: c.ToUint(),
	}, &reply)
	c.Write(reply.Response)
	c.Close()
}
