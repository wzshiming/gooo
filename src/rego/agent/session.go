package agent

import (
	"errors"
	"fmt"
	"rego"
	"rego/cfg"
	"time"
)

type Session struct {
	rego.Unique
	Data           *rego.EncodeBytes
	ConnectTime    time.Time
	LastPacketTime time.Time
	Dirtycount     uint
	SerId          int
}

func NewSession() *Session {
	s := Session{
		ConnectTime:    time.Now(),
		LastPacketTime: time.Now(),
		Dirtycount:     0,
		SerId:          cfg.SelfId,
	}
	s.InitUint()
	s.Data = rego.EnJson(map[string]uint{
		"none": 0,
	})
	return &s
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
		Uniq:  s.ToUint(),
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
		Uniq: s.ToUint(),
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
		Uniq: s.ToUint(),
	}, &reply)
	if err != nil {
		return err
	}
	*s = *reply.Session
	return nil
}

func (s *Session) UnlockSession(reply *Response) error {
	return cfg.ConfForId[s.SerId].Client().Send("Connect.Unlock", UnlockRequest{
		Uniq:  s.ToUint(),
		Reply: reply,
	})
}

func (s *Session) Change(i interface{}) error {
	return cfg.ConfForId[s.SerId].Client().Send("Connect.Change", ChangeRequest{
		Uniq: s.ToUint(),
		Data: rego.EnJson(i),
	})
}
