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

func (conn ConnWeb) ReadMsg() ([]byte, error) {
	_, b, err := conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.DecodeString(string(b))
}

func (conn ConnWeb) WriteMsg(b []byte) error {

	return conn.WriteMessage(websocket.TextMessage, []byte(base64.StdEncoding.EncodeToString(b)))
}

func (conn ConnWeb) SetDeadline(t time.Time) error {
	return conn.UnderlyingConn().SetDeadline(t)
}

func (conn ConnWeb) LocalAddr() string {
	return conn.Conn.LocalAddr().String()
}
func (conn ConnWeb) RemoteAddr() string {
	return conn.Conn.RemoteAddr().String()
}
