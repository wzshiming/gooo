package gooo

import (
	"sync"
)

type ReBuf struct {
	sync.Mutex
	Buf []byte
}

func NewReBuf(size uint) *ReBuf {
	return &ReBuf{
		Buf: make([]byte, size),
	}
}
