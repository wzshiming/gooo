package handel

import (
	"sync"
)

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
