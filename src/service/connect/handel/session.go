package handel

import (
	"gooo/session"
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
