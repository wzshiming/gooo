package server

import (
	"fmt"
	"net"
	"net/rpc"
	"rego"
)

type Server struct {
	rpc.Server
	listen net.Listener
	port   int
}

func NewServer(port int) *Server {
	return &Server{
		Server: *rpc.NewServer(),
		port:   port,
	}
}

func (se *Server) Start() (err error) {
	rego.NOTICE("Server start from port", se.port)
	se.listen, err = net.Listen("tcp", fmt.Sprintf(":%d", se.port))
	if err != nil {
		rego.ERR(err)
		return
	}
	for {
		if conn, err := se.listen.Accept(); err == nil {
			go se.ServeConn(conn)
			//go h.Server.ServeCodec(jsonrpc.NewServerCodec(conn))
		} else {
			rego.ERR(err)
			return err
		}
	}
	rego.NOTICE("Server stop")
	return nil
}

func (se *Server) Stop() {
	se.listen.Close()
}
