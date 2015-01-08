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

func (b *Balance) Allot() int {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.i == b.size {
		b.i = 0
		return 0
	}
	i := b.i
	b.i++
	return i
}
