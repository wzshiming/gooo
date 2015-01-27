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
	"service/connect/iorange"
	"service/connect/route"
	offlprot "service/offline/protocol"
)

type Handel struct {
	handeln.HandelInterface
	Server   *route.MethodServers
	Session  *Sessions
	Online   *Onlines
	Conf     *configs.Configs
	calloffl *router.CallServer
	dataInit []byte
}

func NewHandel(conf *configs.Configs, size uint64) *Handel {
	var t struct {
		Auth   uint32
		UserId uint64
	}
	t.Auth = conf.St.NoLogin
	t.UserId = 0
	b, _ := encoder.Encode(t)

	port := helper.GetPort(conf.Sc[configs.Type][configs.Id].ClientPort)
	helper.EchoPublicPortInfo(configs.Name, port)
	h := Handel{
		Server:   route.NewMethodServers(conf),
		Session:  NewSessions(size),
		Online:   NewOnlines(size),
		Conf:     conf,
		calloffl: router.NewCallServer("Offline", conf),
		dataInit: b,
	}
	ser := connser.NewServer(&h, iorange.NewIORange(1024))
	go ser.StartTCP(port)
	return &h
}

func (h *Handel) Join(c *connser.Connect) {
	//log.Printf("%v %v join\n", c.ToUint(), c.RemoteAddr())
	s := session.NewSession(c)
	s.Data = h.dataInit
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

	var t struct {
		Auth uint32
	}
	encoder.Decode(s.Data, &t)
	err := h.Server.Call(t.Auth, msg[:4], protocol.RpcRequest{
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
				t[k] = v
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
	var r offlprot.InterruptResponse
	if err := h.calloffl.Call("Offline.Interrupt", offlprot.InterruptRequest{
		Data: s.Data,
	}, &r); err != nil {
		return
	}
	h.Session.Del(c.ToUint64())
	h.Online.Del(r.UserId)
	return
}
