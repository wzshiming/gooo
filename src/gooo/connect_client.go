package gooo

import (
	"net"
)

func NewClientTCP(addr string, ior IORanges, bc EventHandel) *Connect {
	return NewClient("tcp", addr, ior, bc)
}

func NewClient(name string, addr string, ior IORanges, bc EventHandel) *Connect {
	conn, err := net.Dial(name, addr)
	if err != nil {
		panic(err)
	}
	return NewConnect(conn, ior, bc)
}
