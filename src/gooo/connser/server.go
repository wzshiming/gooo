package connser

import (
	"log"
	"net"
)

const (
//TIMEOUT = time.Second * 60 * 1
)

type Server struct {
	listener net.Listener
	pending  chan net.Conn
	bc       EventHandel
	ior      IORanges
}

func NewServer(bc EventHandel, ior IORanges) *Server {
	server := &Server{
		pending: make(chan net.Conn, 4),
		ior:     ior,
		bc:      bc,
	}
	return server
}

func (self *Server) StartTCP(port string) (err error) {
	return self.Start("tcp", port)
}

func (self *Server) Start(name string, port string) (err error) {
	self.listener, err = net.Listen(name, port)
	if err != nil {
		log.Println(err)
	}
	go self.StartListen()
	for {
		if conn, err := self.listener.Accept(); err == nil {
			self.Listen(conn)
		} else {
			break
		}
	}
	return
}

func (self *Server) Stop() {
	log.Println("Server stop")
	self.listener.Close()
	close(self.pending)
}

func (self *Server) Listen(conn net.Conn) {
	self.pending <- conn
}

func (self *Server) StartListen() {
	for conn := range self.pending {
		NewConnect(conn, self.ior, self.bc)
	}
}
