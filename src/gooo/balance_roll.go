package gooo

import (
	"sync"
)

type balanceRoll struct {
	lock sync.Mutex
	size int
	i    int
}

func newBalanceRoll(size int) *balanceRoll {
	return &balanceRoll{
		size: size,
		i:    0,
	}
}

func (b *balanceRoll) allot() (i int) {
	b.lock.Lock()
	defer b.lock.Unlock()
	i = 0
	if b.i == b.size {
		b.i = 0
	}
	i = b.i
	b.i++
	return
}
