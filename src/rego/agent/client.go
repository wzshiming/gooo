package agent

import (
	"net"
	"rego"
)

func NewConn(addr string) Conn {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		rego.ERR(err)
		return nil
	}
	return NewConnNet(conn)
}
