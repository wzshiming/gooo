package helper

import (
	"net/rpc"
	"net/rpc/jsonrpc"
	//"net/http"
	"log"
	"net"
	//"../configs"
)

type Handeln struct {
	Server *rpc.Server
}

func NewHandeln() *Handeln {
	return &Handeln{
		Server: rpc.NewServer(),
	}
}

func (h *Handeln) Register(i interface{}) {
	if err := h.Server.Register(i); err != nil {
		log.Fatalf("srv.RegisterName(): %v\n", err)
	}
}

func (h *Handeln) RegisterName(name string, i interface{}) {
	if err := h.Server.RegisterName(name, i); err != nil {
		log.Fatalf("srv.RegisterName(): %v\n", err)
	}
}

func (h *Handeln) Start(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("net.Listen(): %v\n", err)
	}
	defer lis.Close()
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go h.Server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
