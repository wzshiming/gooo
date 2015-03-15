package agent

import (
	"encoding/base64"
	"github.com/gorilla/websocket"
	"time"
)

type ConnWeb struct {
	*websocket.Conn
}

func NewConnWeb(conn *websocket.Conn) ConnWeb {
	return ConnWeb{Conn: conn}
}

func (s ConnWeb) ReadMsg() ([]byte, error) {
	_, b, err := s.ReadMessage()
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.DecodeString(string(b))
}

func (s ConnWeb) WriteMsg(b []byte) error {

	return s.WriteMessage(websocket.TextMessage, []byte(base64.StdEncoding.EncodeToString(b)))
}

func (s ConnWeb) SetDeadline(t time.Time) error {
	return s.UnderlyingConn().SetDeadline(t)
}
