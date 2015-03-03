package server

import (
	"testing"
)

func Test_cs(t *testing.T) {
	go server(t)
	client(t)
}

type Te struct{}

func (t *Te) Echo(a int, r *int) error {
	*r = a
	return nil
}

func server(t *testing.T) {
	s := NewServer(7799)
	s.Register(&Te{})
	s.Start()
}

func client(t *testing.T) {
	c := NewClient("127.0.0.1:7799")
	var a int
	c.Call("Te.Echo", 1000, &a)
	if a != 1000 {
		t.Fail()
	}
}
