package gooo

import (
	"net"
)

func NewClientTCP(addr string, ior IORanges, bc EventHandel) *Connect {
	return newClient("tcp", addr, ior, bc)
}

func newClient(name string, addr string, ior IORanges, bc EventHandel) *Connect {
	conn, err := net.Dial(name, addr)
	if err != nil {
		panic(err)
	}
	return NewConnect(conn, ior, bc)
}
