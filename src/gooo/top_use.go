package gooo

import (
	"sync"
)

type TopUse struct {
	sync.RWMutex
	use  uint64
	pend chan bool
}

func NewTopUse() *TopUse {
	return &TopUse{
		pend: make(chan bool),
	}
}

func (t *TopUse) Start() {
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

func (t *TopUse) Add() {
	t.pend <- true
}

func (t *TopUse) Sub() {
	t.pend <- false
}

func (t *TopUse) Set(i uint64) {
	t.Lock()
	defer t.Unlock()
	t.use = i
}

func (t *TopUse) Get() uint64 {
	t.RLock()
	defer t.RUnlock()
	return t.use
}
