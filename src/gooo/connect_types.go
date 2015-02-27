package gooo

import (
	"io"
	"net"
)

type EventHandel interface {
	Join(c *Connect)
	Mess(c *Connect, msg []byte)
	Exit(c *Connect)
	Timeout(c *Connect)
	Recover(c *Connect, err error)
}

type IORanges interface {
	Reader(r io.Reader) (b []byte, err error)
	Writer(w io.Writer, b []byte) (err error)
}

type Conn interface {
	net.Conn
	ReadMessage() (b []byte, err error)
	WriteMessage(b []byte) (err error)
}
