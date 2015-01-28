package used

import (
	"sync"
)

type Used struct {
	sync.Mutex
	use  []bool
	size int
}

func NewUsed(size int) *Used {
	return &Used{
		use:  make([]bool, size),
		size: 0,
	}
}

func (s *Used) Cap() int {
	return cap(s.use)
}

func (s *Used) IsUse(seat int) bool {
	return s.use[seat]
}

func (s *Used) List() []int {
	s.Lock()
	defer s.Unlock()
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

func (s *Used) Usable() int {
	s.Lock()
	defer s.Unlock()
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
	if s.use[seat] {
		s.size--
		s.use[seat] = false
		return 0
	}
	return -1
}
