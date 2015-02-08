package gooo

import (
	"io"
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
