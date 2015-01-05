package session

import (
	//"net"
	"gooo/encoder"
	"time"
	//"encoding/json"
	"gooo/connser"
	"gooo/uniq"
	"sync"
)

var uni = uniq.NewUniq()

type Session struct {
	sync.Mutex
	Data           map[string]interface{} `json: d`
	ConnectTime    time.Time              `json: c`
	LastPacketTime time.Time              `json: l`
	Dirtycount     uint                   `json: i`
	ServerId       uint                   `json: s`
	Uniq           uint                   `json: u`
	conn           *connser.Connect
}

func NewSession(conn *connser.Connect) *Session {
	return &Session{
		Data:           make(map[string]interface{}, 1024),
		ConnectTime:    time.Now(),
		LastPacketTime: time.Now(),
		Dirtycount:     0,
		conn:           conn,
		Uniq:           uni(),
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
