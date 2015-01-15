package session

import (
	//"net"
	"gooo/encoder"
	"time"
	//"encoding/json"
	"gooo/configs"
	"gooo/connser"
	"sync"
)

type Session struct {
	sync.Mutex
	Data           []byte    `json:"d"`
	ConnectTime    time.Time `json:"c"`
	LastPacketTime time.Time `json:"l"`
	Dirtycount     uint      `json:"i"`
	Uniq           Unique    `json:"u"`
	conn           *connser.Connect
}

func NewSession(conn *connser.Connect) *Session {
	return &Session{
		Data:           []byte("{}"),
		ConnectTime:    time.Now(),
		LastPacketTime: time.Now(),
		Dirtycount:     0,
		conn:           conn,
		Uniq:           NewUnique(uint(configs.Id), conn.ToUint()),
	}
}

func (s *Session) Conn() *connser.Connect {
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
	m, _ = encoder.Encode(s)
	return
}
