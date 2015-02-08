package route

import (
	"errors"
	"fmt"
	"gooo"
)

type CallName struct {
	Name    string
	Allow   uint32
	Unallow uint32
}

type MethodServer struct {
	caller  gooo.CallServer
	methods [][]CallName
	name    string
	nameuse string
}

func NewMethodServer(index int, conf *gooo.Configs) *MethodServer {
	fr := conf.Rc
	fri := fr[index]
	s := MethodServer{
		caller:  *gooo.NewCallServer(fri.Name, conf),
		methods: make([][]CallName, len(fri.Map)),
		name:    fri.Name,
		nameuse: fmt.Sprintf("use%s", fri.Name),
	}

	for k1, v1 := range fri.Map {
		s.methods[k1] = make([]CallName, len(v1.Map))
		for k2, v2 := range v1.Map {
			s.methods[k1][k2] = CallName{
				Name:    fmt.Sprintf("%v.%v", v1.Name, v2),
				Allow:   fri.Allow | v1.Allow,
				Unallow: fri.Unallow | v1.Unallow,
			}
		}
	}
	return &s
}

func (s *MethodServer) Call(c2 uint8, c3 uint8, args gooo.RpcRequest, reply *gooo.RpcResponse) error {
	if c2 >= uint8(len(s.methods)) {
		return errors.New("MethodServer Call index c2 error")
	}
	m := s.methods[c2]
	if c3 >= uint8(len(m)) {
		return errors.New("MethodServer Call index c3 error")
	}
	var t map[string]uint32
	gooo.Decode(args.Session.Data, &t)
	if a := m[c3].Allow; a != (a & t["flag"]) {
		return errors.New("Permission denied")
	}

	if a := m[c3].Unallow; 0 != (a & t["flag"]) {
		return errors.New("Permission denied")
	}

	if a := t[s.nameuse]; a != 0 {
		return s.caller.CallBy(int(a-1), m[c3].Name, args, reply)
	}
	return s.caller.Call(m[c3].Name, args, reply)
}

func (s *MethodServer) Type() string {
	return s.name
}

func (s *MethodServer) Methods() [][]CallName {
	return s.methods
}

type MethodServers []MethodServer

func NewMethodServers(conf *gooo.Configs) *MethodServers {
	fr := conf.Rc
	s := make(MethodServers, len(fr))
	for k, _ := range fr {
		s[k] = *NewMethodServer(k, conf)
	}
	return &s
}

func (s *MethodServers) Call(msg []byte, args gooo.RpcRequest, reply *gooo.RpcResponse) error {
	c1 := msg[1]
	c2 := msg[2]
	c3 := msg[3]
	//c3 := binary.BigEndian.Uint16(msg[2:4])

	if c1 >= byte(len(*s)) {
		return errors.New("MethodServers Call index c1 error")
	}
	ss := (*s)[c1]
	return ss.Call(c2, c3, args, reply)
}
