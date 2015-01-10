package handel

import (
	"gooo/configs"
	"gooo/connser"
	"gooo/helper"
	"gooo/router"
	"log"
	connprot "service/connect/protocol"
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

	var reply connprot.JoinResponse
	h.Server.Call("Connect.Join", connprot.JoinRequest{
		Id: c.ToUint(),
	}, &reply)
	c.Write(reply.Response)
	c.Close()
}
