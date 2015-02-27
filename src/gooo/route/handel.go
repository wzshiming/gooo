package route

import (
	//"encoding/base64"
	"fmt"
	"gooo"
	"gooo/protocol"
	"log"
	//"net/http"
)

type Handel struct {
	gooo.HandelInterface
	Server   *MethodServers
	Session  *Sessions
	Conf     *gooo.Configs
	callauth *gooo.CallServer
}

func NewHandel(conf *gooo.Configs, size uint64, names ...string) *Handel {
	h := Handel{
		Server:   NewMethodServers(conf, names...),
		Session:  NewSessions(size),
		Conf:     conf,
		callauth: gooo.NewCallServer("Auth", conf),
	}

	h.Server.Map().WriteFile(fmt.Sprintf("./conf/%s_map.json", gooo.Name))

	return &h
}

func (h *Handel) Start(ior gooo.IORanges) {
	port := gooo.GetPort(h.Conf.Self().ClientPort)
	gooo.EchoPublicPortInfo(gooo.Name, port)
	ser := gooo.NewServer(h, ior)
	go ser.StartTCP(port)

}

func (h *Handel) Recover(c *gooo.Connect, err error) {
	//log.Printf("%v %v\n", c.RemoteAddr(), err)
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
	//log.Printf("%v msg  %v\n", c.RemoteAddr(), string(msg))

	var reply gooo.RpcResponse

	err := h.Server.Call(msg[:4], gooo.RpcRequest{
		Request: msg[4:],
		Session: s,
	}, &reply)

	var ret []byte
	if err != nil {
		ret = ErrorInfo(err.Error())
	} else {
		if reply.Error != "" {
			ret = ErrorInfo(reply.Error)
		} else if len(reply.Response) != 0 {
			ret = reply.Response
		} else {
			ret = OkInfo()
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

	c.Write(append(msg[:4], ret[:]...))
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
