package main

import (
	"gooo/session"
	"sync"
)

type Onlines struct {
	sync.Mutex
	Map map[uint64]*session.Unique
}

func NewOnlines(size uint64) *Onlines {
	return &Onlines{
		Map: make(map[uint64]*session.Unique, size),
	}
}

func (h *Onlines) Get(u uint64) *session.Unique {
	h.Lock()
	defer h.Unlock()
	return h.Map[u]
}

func (h *Onlines) Set(u uint64, v *session.Unique) {
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
