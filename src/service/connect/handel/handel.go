package handel

import (
	"gooo/configs"
	"gooo/connser"
	"gooo/protocol"
	"gooo/session"
	"log"
	"service/connect/route"
	"sync"
)

type Handel struct {
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

func (h *Handel) Mess(c *connser.Connect, msg []byte) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Client %v Mess error: %v", c.RemoteAddr(), err)
		}
	}()
	id := c.ToUint()
	s := h.Get(id)
	//log.Printf("%v msg  %v\n", c.RemoteAddr(),msg)
	var reply protocol.RpcResponse

	err := h.Server.Call(msg[:4], protocol.RpcRequest{
		Request: msg[4:],
		Session: s,
	}, &reply)

	if err != nil {
		log.Println(err)
		return
	}
	if reply.Error == 0 && len(reply.Response) != 0 {
		c.Write(reply.Response)
	}
	if reply.Data != nil {
		s := h.Get(id)
		s.Lock()
		defer s.Unlock()
		//log.Println(reply.Data)
		for k, v := range *reply.Data {
			s.Data[k] = v
		}
	}

	//log.Printf("%s\n", reply)
}
func (h *Handel) Exit(c *connser.Connect) {
	h.Del(c.ToUint())
	log.Printf("%d  %v is quiting\n", h.Len(), c.RemoteAddr())
}
func (h *Handel) Timeout(c *connser.Connect) {
	log.Printf("%v Timeout\n", c.RemoteAddr())
}

func (h *Handel) Recover(c *connser.Connect, err error) {
	log.Printf("%v error %v\n", c.RemoteAddr(), err)
}
