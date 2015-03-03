package agent

import (
	"rego"
	"sync"
)

type Agent struct {
	maps map[uint]*rego.Session
	msg  func(Conn, *rego.Session, []byte)
}

func NewAgent(max int, msg func(Conn, *rego.Session, []byte)) *Agent {
	return &Agent{
		msg:  msg,
		maps: make(map[uint]*rego.Session, max),
	}
}

func (ag *Agent) join() uint {
	obj := rego.NewSession()
	ag.maps[obj.ToUint()] = obj
	return obj.Uniq
}

func (ag *Agent) Join(conn Conn) {
	uniq := ag.join()
	rego.NOTICE("Join ", conn.RemoteAddr())
	ag.loop(conn, uniq)
	ag.leave(uniq)
	rego.NOTICE("Leave ", conn.RemoteAddr())
}

func (ag *Agent) leave(uniq uint) {
	delete(ag.maps, uniq)
}

func (ag *Agent) get(uniq uint) *rego.Session {
	return ag.maps[uniq]
}

func (ag *Agent) loop(conn Conn, uniq uint) {
	var sy sync.Mutex
	for {
		if b, err := conn.ReadMsg(); err != nil {
			return
		} else if sess := ag.get(uniq); sess == nil {
			return
		} else {
			sy.Lock()
			sess.Refresh()
			ag.msg(conn, sess, b)
			sy.Unlock()
		}
	}
}
