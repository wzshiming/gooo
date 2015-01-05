package handel

import (
	"gooo/connser"
	"log"
	"gooo/protocol"
	"gooo/configs"
	"gate/route"

)

type Handel struct {
	Server   *route.MethodServer
}

func NewHandel(conf *configs.Configs) *Handel {
	return &Handel{
		Server:  route.NewMethodServer("Connect",conf),
	}
}



func (h *Handel) Join(c *connser.Connect) {
	log.Printf("%v join\n", c.RemoteAddr())
}
func (h *Handel) Mess(c *connser.Connect, msg []byte) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Client %v Mess error: %v", c.RemoteAddr(), err)
		}
	}()

	var reply protocol.GateResponse
	h.Server.CallByName("Status.Join", protocol.GateRequest{
		Id:      c.ToUint(),
	}, &reply)
	c.Write(reply.Response)
	//c.Close()

}
func (h *Handel) Exit(c *connser.Connect) {
	log.Printf("%v is quiting\n", c.RemoteAddr())
}

func (h *Handel) Timeout(c *connser.Connect) {
	log.Printf("%v Timeout\n", c.RemoteAddr())
}

func (h *Handel) Recover(c *connser.Connect, err error) {
	log.Printf("%v error %v\n", c.RemoteAddr(), err)
}
