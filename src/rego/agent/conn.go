package agent

import (
	"net"
)

type Conn interface {
	net.Conn
	ReadMsg() ([]byte, error)
	WriteMsg([]byte) error
}
