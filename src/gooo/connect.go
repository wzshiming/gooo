package gooo

import (
	"code.google.com/p/go.net/websocket"
	"net"
	"unsafe"
)

type Connect struct {
	net.Conn
	Dirty uint
	ior   IORanges
	Bc    EventHandel
}

func NewConnect(conn net.Conn, ior IORanges, bc EventHandel) *Connect {
	connect := &Connect{
		Conn: conn,
		ior:  ior,
		Bc:   bc,
	}
	connect.listen()
	return connect
}

func ToConnect(u uint) *Connect {
	return (*Connect)(unsafe.Pointer((uintptr)(u)))
}

func (self *Connect) ToUint() uint {
	return (uint)((uintptr)(unsafe.Pointer(self)))
}

func (self *Connect) ToUint64() uint64 {
	return (uint64)((uintptr)(unsafe.Pointer(self)))
}

func (self *Connect) RemoteAddr() string {
	if b := self.Websocket(); b != nil {
		return b.Request().RemoteAddr
	}
	return self.Conn.RemoteAddr().String()
}

func (self *Connect) LocalAddr() string {
	return self.Conn.LocalAddr().String()
}

func (self *Connect) Type() string {
	if _, ok := self.Conn.(*websocket.Conn); ok {
		return "websocket"
	}
	return "tcp"
}

func (self *Connect) Websocket() *websocket.Conn {
	if t, ok := self.Conn.(*websocket.Conn); ok {
		return t
	}
	return nil
}

func (self *Connect) Write(message []byte) {
	err := self.ior.Writer(self.Conn, message)
	if err != nil {
		self.Bc.Recover(self, err)
	}
	return
}

func (self *Connect) listen() {
	defer func() {
		if err := recover(); err != nil {
			self.Bc.Recover(self, err.(error))
		}
		self.Close()
	}()
	for self.Bc.Join(self); self.Dirty != 255; {
		msg, err := self.ior.Reader(self.Conn)
		if err != nil {
			self.Bc.Recover(self, err)
			self.Dirty += 1
			continue
		}
		self.Bc.Mess(self, msg)
	}
}

func (self *Connect) Close() error {
	self.Bc.Exit(self)
	return self.Conn.Close()
}
