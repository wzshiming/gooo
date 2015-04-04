package agent

import (
	"errors"
	"fmt"
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
	readlist       chan *rego.EncodeBytes
}

func NewSession() *Session {
	s := Session{
		ConnectTime:    time.Now(),
		LastPacketTime: time.Now(),
		Dirtycount:     0,
		SerId:          cfg.SelfId,
	}
	s.Uniq = s.toUint()
	s.Data = rego.EnJson(map[string]uint{
		"__Uniq__": s.Uniq,
	})
	return &s
}

func (s *Session) Readlist() chan *rego.EncodeBytes {
	if s.readlist == nil {
		s.readlist = make(chan *rego.EncodeBytes, 10)
	}
	return s.readlist
}

func (s *Session) toUint() uint {
	return (uint)((uintptr)(unsafe.Pointer(s)))
}
func (s *Session) ToUint() uint {
	return s.Uniq
}

func (s *Session) Refresh() {
	s.LastPacketTime = time.Now()
}

func (s *Session) Push(reply interface{}) (err error) {
	return s.Send(&Response{
		Response: rego.EnJson(reply),
	})
}

func (s *Session) Send(reply *Response) (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = errors.New("Session.Send: " + fmt.Sprintln(x))
		}
	}()
	return cfg.ConfForId[s.SerId].Client().Send("Connect.Push", PushRequest{
		Uniq:  s.Uniq,
		Reply: reply,
	})
}

func (s *Session) SyncSession() (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = errors.New("Session.SyncSession: " + fmt.Sprintln(x))
		}
	}()
	var reply LockResponse
	err = cfg.ConfForId[s.SerId].Client().Call("Connect.Sync", LockRequest{
		Uniq: s.Uniq,
	}, &reply)
	if err != nil {
		return err
	}
	*s = *reply.Session
	return nil
}

func (s *Session) LockSession() (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = errors.New("Session.LockSession: " + fmt.Sprintln(x))
		}
	}()
	var reply LockResponse
	err = cfg.ConfForId[s.SerId].Client().Call("Connect.Lock", LockRequest{
		Uniq: s.Uniq,
	}, &reply)
	if err != nil {
		return err
	}
	*s = *reply.Session
	return nil
}

func (s *Session) UnlockSession(reply *Response) error {
	return cfg.ConfForId[s.SerId].Client().Send("Connect.Unlock", UnlockRequest{
		Uniq:  s.Uniq,
		Reply: reply,
	})
}
