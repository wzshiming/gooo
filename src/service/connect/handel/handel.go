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

type Sessions struct {
	sync.Mutex
	Map map[uint]*session.Session
}

func NewSessions(size int) *Sessions {
	return &Sessions{
		Map: make(map[uint]*session.Session, size),
	}
}

func (h *Sessions) Get(u uint) *session.Session {
	h.Lock()
	defer h.Unlock()
	return h.Map[u]
}

func (h *Sessions) Set(u uint, v *session.Session) {
	h.Lock()
	defer h.Unlock()
	h.Map[u] = v
}

func (h *Sessions) Del(u uint) {
	h.Lock()
	defer h.Unlock()
	delete(h.Map, u)
}

func (h *Sessions) Len() int {
	h.Lock()
	defer h.Unlock()
	return len(h.Map)
}

type Onlines struct {
	sync.Mutex
	Map map[int64]bool
}

func NewOnlines(size int) *Onlines {
	return &Onlines{
		Map: make(map[int64]bool, size),
	}
}

func (h *Onlines) Get(u int64) bool {
	h.Lock()
	defer h.Unlock()
	return h.Map[u]
}

func (h *Onlines) Set(u int64, v bool) {
	h.Lock()
	defer h.Unlock()
	h.Map[u] = v
}

func (h *Onlines) Del(u int64) {
	h.Lock()
	defer h.Unlock()
	delete(h.Map, u)
}

func (h *Onlines) Len() int {
	h.Lock()
	defer h.Unlock()
	return len(h.Map)
}

type Handel struct {
	helper.HandelInterface
	Server  *route.MethodServers
	Session *Sessions
	Online  *Onlines
}

func NewHandel(conf *configs.Configs) *Handel {
	return &Handel{
		Server:  route.NewMethodServers(conf),
		Session: NewSessions(1024),
		Online:  NewOnlines(1024),
	}
}

func (h *Handel) Join(c *connser.Connect) {
	log.Printf("%v %v join\n", c.ToUint(), c.RemoteAddr())
	h.Session.Set(c.ToUint(), session.NewSession(c))
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
	h.Session.Del(c.ToUint())
	log.Printf("%d  %v is quiting\n", h.Session.Len(), c.RemoteAddr())
}
