package route

import (
	"encoding/binary"
	"errors"
	"fmt"
	//"gooo/balance"
	"gooo/configs"
	//"gooo/protocol"
	"gooo/router"
	//"gooo/session"
	"log"
	//"net/rpc"
	//"sync"
)

type MethodServer struct {
	caller  router.CallServer
	methods [][]string
}

func NewMethodServer(typ string, conf *configs.Configs) *MethodServer {
	fr := conf.Rc.Devel
	s := MethodServer{
		caller:  *router.NewCallServer(typ, conf),
		methods: make([][]string, len(fr)),
	}

	i := 0
	for k1, v1 := range fr[typ] {
		s.methods[i] = make([]string, len(v1))
		for k2, v2 := range v1 {
			s.methods[i][k2] = fmt.Sprintf("%v.%v", k1, v2)
		}
		i++
	}
	return &s
}

func (s *MethodServer) Call(c2 uint8, c3 uint16, args, reply interface{}) error {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("MethodServer Call:", err)
		}
	}()
	return s.caller.Call(s.methods[c2][c3], args, reply)
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
	fr := conf.Rc.Devel
	s := make(MethodServers, len(fr))
	i := 0
	for k, _ := range fr {
		s[i] = *NewMethodServer(k, conf)
		i++
	}
	return &s
}

func (s *MethodServers) Call(msg []byte, args, reply interface{}) error {
	c1 := msg[0]
	c2 := msg[1]
	c3 := binary.BigEndian.Uint16(msg[2:4])
	ss := (*s)[c1]
	if c1 >= byte(len(*s)) {
		return errors.New("call index error")
	}
	return ss.Call(c2, c3, args, reply)
}
