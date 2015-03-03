package agent

import (
	"github.com/gorilla/websocket"
)

type ConnWeb struct {
	*websocket.Conn
}

func NewConnWeb(conn *websocket.Conn) *ConnWeb {
	return &ConnWeb{Conn: conn}
}

func (s *ConnWeb) ReadMsg() ([]byte, error) {
	_, b, err := s.Conn.ReadMessage()
	return b, err
}

func (s *ConnWeb) WriteMsg(b []byte) error {
	return s.Conn.WriteMessage(websocket.TextMessage, b)
}
