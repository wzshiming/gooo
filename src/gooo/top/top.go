package top

import (
	"sync"
)

type Top struct {
	sync.RWMutex
	use  uint64
	pend chan bool
}

func NewTop() *Top {
	return &Top{
		pend: make(chan bool),
	}
}

func (t *Top) Start() {
	for b := range t.pend {
		t.Lock()
		if b {
			t.use++
		} else if t.use != 0 {
			t.use--
		}
		t.Unlock()
	}
}

func (t *Top) Add() {
	t.pend <- true
}

func (t *Top) Sub() {
	t.pend <- false
}

func (t *Top) Set(i uint64) {
	t.Lock()
	defer t.Unlock()
	t.use = i
}

func (t *Top) Get() uint64 {
	t.RLock()
	defer t.RUnlock()
	return t.use
}
