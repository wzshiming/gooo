package agent

import (
	"net"
	"time"
)

type Conn interface {
	ReadMsg() ([]byte, error)
	WriteMsg([]byte) error
	Close() error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	SetDeadline(t time.Time) error
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
}
