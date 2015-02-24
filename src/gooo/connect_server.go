package gooo

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"net"
	"net/http"
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
	return self.start("tcp", port)
}

func (self *Server) WebsocketHandler() websocket.Handler {
	return websocket.Handler(func(conn *websocket.Conn) {
		self.Listen(conn)
	})
}

func (self *Server) StartWebsocket(port string) (err error) {
	return http.ListenAndServe(port, self.WebsocketHandler())
}

func (self *Server) start(name string, port string) (err error) {
	self.listener, err = net.Listen(name, port)
	if err != nil {
		log.Println(err)
	}
	for {
		if conn, err := self.listener.Accept(); err == nil {
			go self.Listen(conn)
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
	NewConnect(conn, self.ior, self.bc)
}
