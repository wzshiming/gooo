package handel

import (
	"gooo/session"
	"sync"
)

type Sessions struct {
	sync.Mutex
	Map map[uint64]*session.Session
}

func NewSessions(size uint64) *Sessions {
	return &Sessions{
		Map: make(map[uint64]*session.Session, size),
	}
}

func (h *Sessions) Get(u uint64) *session.Session {
	h.Lock()
	defer h.Unlock()
	return h.Map[u]
}

func (h *Sessions) Set(u uint64, v *session.Session) {
	h.Lock()
	defer h.Unlock()
	h.Map[u] = v
}

func (h *Sessions) Del(u uint64) {
	h.Lock()
	defer h.Unlock()
	delete(h.Map, u)
}

func (h *Sessions) Len() int {
	h.Lock()
	defer h.Unlock()
	return len(h.Map)
}
