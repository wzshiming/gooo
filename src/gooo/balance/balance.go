package balance

import (
	"sync"
)

type Balance struct {
	lock sync.Mutex
	size int
	i    int
}

func NewBalance(size int) *Balance {
	return &Balance{
		size: size,
		i:    0,
	}
}

func (b *Balance) Allot() (i int) {
	b.lock.Lock()
	defer b.lock.Unlock()
	i = 0
	if b.i != b.size {
		i = b.i
		b.i++
	} else {
		b.i = 0
	}
	return
}
