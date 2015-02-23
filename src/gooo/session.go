package gooo

import (
	"sync"
	"time"
)

type Session struct {
	sync.Mutex
	Data           []byte    `json:"d"`
	ConnectTime    time.Time `json:"c"`
	LastPacketTime time.Time `json:"l"`
	Dirtycount     uint      `json:"i"`
	Uniq           Unique    `json:"u"`
	ConnType       string    `json:"t"`
	conn           *Connect
}

func NewSession(conn *Connect) *Session {
	s := Session{
		Data:           []byte("{}"),
		ConnectTime:    time.Now(),
		LastPacketTime: time.Now(),
		Dirtycount:     0,
		ConnType:       conn.Type(),
		conn:           conn,
		Uniq:           NewUnique(Id, conn.ToUint64()),
	}

	return &s
}

func (s *Session) Conn() *Connect {
	return s.conn
}

func (s *Session) Refresh() {
	s.Lock()
	defer s.Unlock()
	s.LastPacketTime = time.Now()
}

func (s *Session) Bytes() (m []byte) {
	s.Lock()
	defer s.Unlock()
	m, _ = Encode(s)
	return
}
