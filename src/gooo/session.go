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
	conn           *Connect
}

func NewSession(conn *Connect) *Session {
	return &Session{
		Data:           []byte("{}"),
		ConnectTime:    time.Now(),
		LastPacketTime: time.Now(),
		Dirtycount:     0,
		conn:           conn,
		Uniq:           NewUnique(Id, conn.ToUint64()),
	}
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
