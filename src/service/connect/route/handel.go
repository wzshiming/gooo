package route

import (
	"gooo"
	"gooo/protocol"
	"log"
)

type Handel struct {
	gooo.HandelInterface
	Server   *MethodServers
	Session  *Sessions
	Conf     *gooo.Configs
	callauth *gooo.CallServer
}

func NewHandel(conf *gooo.Configs, size uint64) *Handel {
	port := gooo.GetPort(conf.Self().ClientPort)
	gooo.EchoPublicPortInfo(gooo.Name, port)
	h := Handel{
		Server:   NewMethodServers(conf),
		Session:  NewSessions(size),
		Conf:     conf,
		callauth: gooo.NewCallServer("Auth", conf),
	}
	ser := gooo.NewServer(&h, NewIORange(1024))
	go ser.StartTCP(port)
	return &h
}
func (h *Handel) Recover(c *gooo.Connect, err error) {

}

func (h *Handel) Join(c *gooo.Connect) {
	//log.Printf("%v %v join\n", c.ToUint(), c.RemoteAddr())
	s := gooo.NewSession(c)
	h.Session.Set(c.ToUint64(), s)
}

func (h *Handel) Mess(c *gooo.Connect, msg []byte) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Client %v Mess error: %v", c.RemoteAddr(), err)
			gooo.RecoverInfo()
		}
	}()
	id := c.ToUint64()
	s := h.Session.Get(id)
	s.Lock()
	defer s.Unlock()
	//log.Printf("%v msg  %v\n", c.RemoteAddr(),msg)
	var reply gooo.RpcResponse

	err := h.Server.Call(msg[:4], gooo.RpcRequest{
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
			gooo.Decode(s.Data, &t)
			for k, v := range *reply.Data {
				if v == nil {
					delete(t, k)
				} else {
					t[k] = v
				}
			}
			if b, err := gooo.Encode(t); err == nil {
				s.Data = b
			}
		}
	}
	return
}

func (h *Handel) Exit(c *gooo.Connect) {
	id := c.ToUint64()
	s := h.Session.Get(id)
	h.Session.Del(c.ToUint64())
	var r protocol.InterruptResponse
	h.callauth.Call("Offline.Interrupt", protocol.InterruptRequest{
		Data: s.Data,
	}, &r)
	return
}
