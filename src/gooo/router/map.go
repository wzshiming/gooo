package router

import (
	"gooo/balance"
	"gooo/configs"
	"gooo/helper"
	"net/rpc"
)

type CallServer struct {
	Client []*rpc.Client
	Type   string
	bal    balance.Balances
}

func NewCallServer(tye string, conf *configs.Configs) *CallServer {
	if tye == configs.Type {
		return nil
	}
	fs := conf.Sc
	servers := fs[tye]
	s := CallServer{
		Type:   tye,
		Client: make([]*rpc.Client, len(servers)),
		bal:    balance.NewBalance(len(servers)),
	}
	for k, server := range servers {
		s.Client[k] = server.Conn()
		//log.Println(conf.Name, "Register", fmt.Sprintf("%v_%v", server.Name, k))
	}
	return &s
}

func (s *CallServer) Call(method string, args, reply interface{}) (err error) {
	defer helper.Recover()
	return s.Client[s.bal.Allot()].Call(method, args, reply)
}

func (s *CallServer) CallBy(index int, method string, args, reply interface{}) (err error) {
	defer helper.Recover()
	return s.Client[index].Call(method, args, reply)
}

type CallServers map[string]*CallServer

func NewCallServers(conf *configs.Configs) *CallServers {
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
	return (*s)[typ].Call(method, args, reply)
}
