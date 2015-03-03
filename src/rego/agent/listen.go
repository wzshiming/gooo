package agent

import (
	"fmt"
	"net"
	"rego"
)

type Listener struct {
	listener   net.Listener
	listenfunc func(Conn)
	port       int
}

func NewListener(port int, listenfunc func(Conn)) *Listener {
	return &Listener{
		listenfunc: listenfunc,
		port:       port,
	}
}

func (se *Listener) Start() error {
	var err error
	rego.NOTICE("Listen start from port", se.port)
	se.listener, err = net.Listen("tcp", fmt.Sprintf(":%d", se.port))
	if err != nil {
		rego.ERR(err)
		return err
	}
	for {
		if conn, err := se.listener.Accept(); err == nil {
			go se.Listen(NewConnNet(conn))
		} else {
			rego.ERR(err)
			return err
		}
	}
	rego.NOTICE("Listen stop")
	return nil
}

func (se *Listener) Stop() {
	se.listener.Close()
}

func (se *Listener) Listen(conn Conn) {
	se.listenfunc(conn)
	conn.Close()
}
