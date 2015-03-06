package agent

import (
	"errors"
	"rego"
	"rego/cfg"
	"time"
	"unsafe"
)

type Session struct {
	Data           *rego.EncodeBytes
	ConnectTime    time.Time
	LastPacketTime time.Time
	Dirtycount     uint
	Uniq           uint
	SerId          int
}

func NewSession() *Session {
	s := Session{
		ConnectTime:    time.Now(),
		LastPacketTime: time.Now(),
		Dirtycount:     0,
		SerId:          cfg.SelfId,
	}
	s.Uniq = s.ToUint()
	s.Data = rego.EnJson(map[string]uint{
		"_UID": s.Uniq,
	})
	return &s
}

func (s *Session) ToUint() uint {
	return (uint)((uintptr)(unsafe.Pointer(s)))
}

func (s *Session) Refresh() {
	s.LastPacketTime = time.Now()
}

func (s *Session) Send(b []byte) (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = errors.New("Session.Send: index out of range")
		}
	}()
	return cfg.ConfForId[s.SerId].Client().Send("Push.Send", PushRequest{
		Uniq: s.Uniq,
		Msg:  b,
	})
}
