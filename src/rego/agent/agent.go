package agent

import (
	"rego"
	"sync"
)

type User struct {
	Session
	Conn
}

type Agent struct {
	maps map[uint]*User
	msg  func(*User, []byte) error
}

func NewAgent(max int, msg func(*User, []byte) error) *Agent {
	return &Agent{
		msg:  msg,
		maps: make(map[uint]*User, max),
	}
}

func (ag *Agent) Join(conn Conn) {
	obj := NewSession()
	uniq := obj.ToUint()
	user := User{
		Session: *obj,
		Conn:    conn,
	}
	ag.maps[uniq] = &user
	rego.NOTICE("Join ", conn.RemoteAddr())
	ag.loop(&user)
	ag.leave(uniq)
	rego.NOTICE("Leave ", conn.RemoteAddr())
}

func (ag *Agent) leave(uniq uint) {
	delete(ag.maps, uniq)
}

func (ag *Agent) Get(uniq uint) *User {
	return ag.maps[uniq]
}

func (ag *Agent) loop(user *User) {
	var sy sync.Mutex
	for {
		if b, err := user.Conn.ReadMsg(); err != nil {
			return
		} else {
			sy.Lock()
			err = ag.msg(user, b)
			sy.Unlock()
			if err != nil {
				return
			}
		}
	}
}
