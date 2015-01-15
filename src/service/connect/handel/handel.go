package handel

import (
	"fmt"
	"gooo/configs"
	"gooo/connser"
	"gooo/encoder"
	"gooo/helper"
	"gooo/protocol"
	"gooo/session"
	"log"
	"service/connect/route"
	"sync"
)

type Handel struct {
	helper.HandelInterface
	Server   *route.MethodServers
	session  map[uint]*session.Session
	sessLock sync.Mutex
}

func NewHandel(conf *configs.Configs) *Handel {
	return &Handel{
		Server:  route.NewMethodServers(conf),
		session: make(map[uint]*session.Session),
	}
}

func (h *Handel) Get(u uint) *session.Session {
	h.sessLock.Lock()
	defer h.sessLock.Unlock()
	return h.session[u]
}

func (h *Handel) Set(u uint, v *session.Session) {
	h.sessLock.Lock()
	defer h.sessLock.Unlock()
	h.session[u] = v
}

func (h *Handel) Del(u uint) {
	h.sessLock.Lock()
	defer h.sessLock.Unlock()
	delete(h.session, u)
}

func (h *Handel) Len() int {
	h.sessLock.Lock()
	defer h.sessLock.Unlock()
	return len(h.session)
}

func (h *Handel) Join(c *connser.Connect) {
	log.Printf("%v %v join\n", c.ToUint(), c.RemoteAddr())
	h.Set(c.ToUint(), session.NewSession(c))
}

func ErrorInfo(s string) []byte {
	return []byte(fmt.Sprintf("{\"e\":\"%s\"}", s))
}

func OkInfo() []byte {
	return []byte("{\"e\":\"\"}")
}

func (h *Handel) Mess(c *connser.Connect, msg []byte) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Client %v Mess error: %v", c.RemoteAddr(), err)
			helper.RecoverInfo()
		}
	}()
	id := c.ToUint()
	s := h.Get(id)
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
		if reply.Data != nil {
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
	h.Del(c.ToUint())
	log.Printf("%d  %v is quiting\n", h.Len(), c.RemoteAddr())
}
