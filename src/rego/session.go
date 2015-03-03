package rego

import (
	"time"
	"unsafe"
)

type Session struct {
	Data           *EncodeStream `json:"d"`
	ConnectTime    time.Time     `json:"c"`
	LastPacketTime time.Time     `json:"l"`
	Dirtycount     uint          `json:"i"`
	Uniq           uint          `json:"u"`
}

func NewSession() *Session {
	s := Session{
		Data:           NewEncodeStream(),
		ConnectTime:    time.Now(),
		LastPacketTime: time.Now(),
		Dirtycount:     0,
	}
	s.Uniq = s.ToUint()
	return &s
}

func (s *Session) ToUint() uint {
	return (uint)((uintptr)(unsafe.Pointer(s)))
}

func (s *Session) Refresh() {
	s.LastPacketTime = time.Now()
}
