package gooo

import (
	"net/rpc"
)

type CallServer struct {
	Client []*rpc.Client
	Type   string
	bal    balance
}

func NewCallServer(tye string, conf *Configs) *CallServer {
	if tye == Type {
		return nil
	}
	fs := conf.Sc
	servers := fs[tye]
	s := CallServer{
		Type:   tye,
		Client: make([]*rpc.Client, len(servers)),
		bal:    newBalanceRoll(len(servers)),
	}
	for k, server := range servers {
		s.Client[k] = server.Conn()
		//log.Println(conf.Name, "Register", fmt.Sprintf("%v_%v", server.Name, k))
	}
	return &s
}

func (s *CallServer) Size() int {
	return len(s.Client)
}

func (s *CallServer) Call(method string, args, reply interface{}) (err error) {
	defer Recover()
	return s.Client[s.bal.allot()].Call(method, args, reply)
}

func (s *CallServer) CallBy(index int, method string, args, reply interface{}) (err error) {
	defer Recover()
	return s.Client[index].Call(method, args, reply)
}

func (s *CallServer) CallByUniq(uniq *Unique, method string, args, reply interface{}) (err error) {
	defer Recover()
	return s.Client[uniq.Server].Call(method, args, reply)
}

func (s *CallServer) CallBySession(sess *Session, method string, args, reply interface{}) (err error) {
	defer Recover()
	return s.Client[sess.Uniq.Server].Call(method, args, reply)
}

func (s *CallServer) CallByUniqs(uniqs []*Unique, method string, args interface{}) (err error) {
	for _, uniq := range uniqs {
		var t int
		s.CallByUniq(uniq, method, args, &t)
	}
	return
}

func (s *CallServer) CallBySessions(sesss []*Session, method string, args interface{}) (err error) {
	for _, sess := range sesss {
		var t int
		s.CallBySession(sess, method, args, &t)
	}
	return
}

type CallServers map[string]*CallServer

func NewCallServers(conf *Configs) *CallServers {
	fs := conf.Sc
	s := make(CallServers, len(fs))
	for k, _ := range fs {
		cs := NewCallServer(k, conf)
		if cs != nil {
			s[k] = cs
		}
	}
	return &s
}

func (s *CallServers) Call(typ, method string, args, reply interface{}) (err error) {
	defer Recover()
	return (*s)[typ].Call(method, args, reply)
}
