package route

import (
	//"encoding/binary"
	"errors"
	"fmt"
	//"gooo/balance"
	"gooo/configs"
	//"gooo/protocol"
	//"gooo/helper"
	"gooo/router"
	//"gooo/session"
	//"log"
	//"net/rpc"
	//"sync"
)

type CallName struct {
	Name string
	Auth uint32
}

type MethodServer struct {
	caller  router.CallServer
	methods [][]CallName
	name    string
}

func NewMethodServer(index int, conf *configs.Configs) *MethodServer {
	fr := conf.Rc
	fri := fr[index]
	s := MethodServer{
		caller:  *router.NewCallServer(fri.Name, conf),
		methods: make([][]CallName, len(fr)),
		name:    fri.Name,
	}

	for k1, v1 := range fri.Map {
		s.methods[k1] = make([]CallName, len(v1.Map))
		for k2, v2 := range v1.Map {
			s.methods[k1][k2] = CallName{
				Name: fmt.Sprintf("%v.%v", v1.Name, v2),
				Auth: fri.Auth | v1.Auth,
			}
		}
	}
	return &s
}

func (s *MethodServer) Call(auth uint32, c2 uint8, c3 uint8, args, reply interface{}) error {
	if c2 >= uint8(len(s.methods)) {
		return errors.New("MethodServer Call index c2 error")
	}
	m := s.methods[c2]
	if c3 >= uint8(len(m)) {
		return errors.New("MethodServer Call index c3 error")
	}

	if a := m[c3].Auth; a != (a & auth) {
		return errors.New("Permission denied")
	}

	return s.caller.Call(m[c3].Name, args, reply)
}

func (s *MethodServer) Type() string {
	return s.name
}

func (s *MethodServer) Methods() [][]CallName {
	return s.methods
}

/*
func (self *MethodServer) CallsBy(clients []session.Unique, name string, args []byte, reply interface{}) (err error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("CallsBy error: ", err)
		}
	}()
	size := len(self.Client)
	cs := make([][]uint, size)
	var sr protocol.SendRequest
	sr.Data = args
	for _, v := range clients {
		if v.Server < uint(size) {
			cs[v.Server] = append(cs[v.Server], v.Id)
		}
	}

	for i := 0; i != size; i++ {
		go self.Client[i].Call(name, protocol.SendRequest{
			Clients: cs[i],
			Data:    args,
		}, reply)
	}
	return
}
*/
type MethodServers []MethodServer

func NewMethodServers(conf *configs.Configs) *MethodServers {
	fr := conf.Rc
	s := make(MethodServers, len(fr))
	for k, _ := range fr {
		s[k] = *NewMethodServer(k, conf)
	}
	return &s
}

func (s *MethodServers) Call(auth uint32, msg []byte, args, reply interface{}) error {
	c1 := msg[1]
	c2 := msg[2]
	c3 := msg[3]
	//c3 := binary.BigEndian.Uint16(msg[2:4])

	if c1 >= byte(len(*s)) {
		return errors.New("MethodServers Call index c1 error")
	}
	ss := (*s)[c1]
	return ss.Call(auth, c2, c3, args, reply)
}
