package gooo

import (
	"sync"
)

type Used struct {
	sync.RWMutex
	use  []bool
	size int
}

func NewUsed(size int) *Used {
	return &Used{
		use:  make([]bool, size),
		size: 0,
	}
}

func (s *Used) Size() int {
	s.RLock()
	defer s.RUnlock()
	return s.size
}

func (s *Used) Cap() int {
	return cap(s.use)
}

func (s *Used) IsUse(seat int) bool {
	s.RLock()
	defer s.RUnlock()
	if seat >= s.Cap() {
		return false
	}
	return s.use[seat]
}

func (s *Used) List() []int {
	s.RLock()
	defer s.RUnlock()
	r := make([]int, s.size)
	i := 0
	for k, v := range s.use {
		if v {
			r[i] = k
			i++
			if i == s.size {
				return r
			}
		}
	}
	return r
}

func (s *Used) Unusable() int {
	s.RLock()
	defer s.RUnlock()
	for k, v := range s.use {
		if v {
			return k
		}
	}
	return -1
}

func (s *Used) Usable() int {
	s.RLock()
	defer s.RUnlock()
	for k, v := range s.use {
		if !v {
			return k
		}
	}
	return -1
}

func (s *Used) Join() int {
	s.Lock()
	defer s.Unlock()
	if s.size < s.Cap() {
		for k, v := range s.use {
			if !v {
				s.use[k] = true
				s.size++
				return k
			}
		}
	}
	return -1
}

func (s *Used) Leave(seat int) int {
	s.Lock()
	defer s.Unlock()
	if seat >= s.Cap() {
		return -1
	}
	if s.use[seat] {
		s.size--
		s.use[seat] = false
		return 0
	}
	return -1
}
