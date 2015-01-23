package handel

import (
	"sync"
)

type Onlines struct {
	sync.Mutex
	Map map[uint64]bool
}

func NewOnlines(size uint64) *Onlines {
	return &Onlines{
		Map: make(map[uint64]bool, size),
	}
}

func (h *Onlines) Get(u uint64) bool {
	h.Lock()
	defer h.Unlock()
	return h.Map[u]
}

func (h *Onlines) Set(u uint64, v bool) {
	h.Lock()
	defer h.Unlock()
	h.Map[u] = v
}

func (h *Onlines) Del(u uint64) {
	h.Lock()
	defer h.Unlock()
	delete(h.Map, u)
}

func (h *Onlines) Len() int {
	h.Lock()
	defer h.Unlock()
	return len(h.Map)
}
