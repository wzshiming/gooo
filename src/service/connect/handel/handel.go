package handel

import (
	"gooo/configs"
	"gooo/connser"
	"gooo/encoder"
	"gooo/handeln"
	"gooo/helper"
	"gooo/protocol"
	"gooo/router"
	"gooo/session"
	"log"
	authprot "service/auth/protocol"
	"service/connect/iorange"
	"service/connect/route"
)

type Handel struct {
	handeln.HandelInterface
	Server   *route.MethodServers
	Session  *Sessions
	Conf     *configs.Configs
	callauth *router.CallServer
}

func NewHandel(conf *configs.Configs, size uint64) *Handel {

	port := helper.GetPort(conf.Sc[configs.Type][configs.Id].ClientPort)
	helper.EchoPublicPortInfo(configs.Name, port)
	h := Handel{
		Server:   route.NewMethodServers(conf),
		Session:  NewSessions(size),
		Conf:     conf,
		callauth: router.NewCallServer("Auth", conf),
	}
	ser := connser.NewServer(&h, iorange.NewIORange(1024))
	go ser.StartTCP(port)
	return &h
}

func (h *Handel) Join(c *connser.Connect) {
	//log.Printf("%v %v join\n", c.ToUint(), c.RemoteAddr())
	s := session.NewSession(c)
	h.Session.Set(c.ToUint64(), s)
}

func (h *Handel) Mess(c *connser.Connect, msg []byte) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Client %v Mess error: %v", c.RemoteAddr(), err)
			helper.RecoverInfo()
		}
	}()
	id := c.ToUint64()
	s := h.Session.Get(id)
	s.Lock()
	defer s.Unlock()
	//log.Printf("%v msg  %v\n", c.RemoteAddr(),msg)
	var reply protocol.RpcResponse

	err := h.Server.Call(msg[:4], protocol.RpcRequest{
		Request: msg[4:],
		Session: s,
	}, &reply)

	if err != nil {
		c.Write(append(msg[:4], ErrorInfo(err.Error())...))
	} else {
		if reply.Error != "" {
			c.Write(append(msg[:4], ErrorInfo(reply.Error)...))
		} else if len(reply.Response) != 0 {
			c.Write(append(msg[:4], reply.Response...))
		} else {
			c.Write(append(msg[:4], OkInfo()...))
		}
		if reply.Coverage != nil {
			s.Data = reply.Coverage
		} else if reply.Data != nil {
			var t map[string]interface{}
			encoder.Decode(s.Data, &t)
			for k, v := range *reply.Data {
				if v == nil {
					delete(t, k)
				} else {
					t[k] = v
				}
			}
			if b, err := encoder.Encode(t); err == nil {
				s.Data = b
			}
		}
	}
	return
}
func (h *Handel) Exit(c *connser.Connect) {
	id := c.ToUint64()
	s := h.Session.Get(id)
	h.Session.Del(c.ToUint64())
	var r authprot.InterruptResponse
	if err := h.callauth.Call("Offline.Interrupt", authprot.InterruptRequest{
		Data: s.Data,
	}, &r); err != nil {
		return
	}
	return
}
