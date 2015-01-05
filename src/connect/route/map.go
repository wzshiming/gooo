package route

import (
	"gooo/configs"
	"encoding/binary"
	"fmt"
	"log"
	"net/rpc"
	"errors"
    "gooo/balance"
	//"sync"
)

type MethodServer struct {
	Client  []*rpc.Client
	Name    string
	Methods [][]string
    bal     balance.Balances
}

func NewMethodServer(name string, cs, ms int)* MethodServer{
    return &MethodServer{
        Name:    name,
        Client:  make([]*rpc.Client, cs),
        Methods: make([][]string, ms),
        bal:     balance.NewBalance(cs),
    }
}

func (self *MethodServer) Call(name string,args, reply interface{}) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Call error: ", err)
		}
	}()
	return self.Client[self.bal.Allot()].Call(name, args, reply)
}

type MapMethodServer []MethodServer

func (self *MapMethodServer) Call(msg []byte, args, reply interface{}) error {
	c1 := msg[0]
	c2 := msg[1]
	c3 := binary.BigEndian.Uint16(msg[2:4])
	s := (*self)[c1]
	if c1 >= byte(len(*self)) || c2 >= byte(len(s.Methods)) || c3 >= uint16(len(s.Methods[c2])) {
		return errors.New("call index error")
	}
	return s.Call(s.Methods[c2][c3], args, reply)
}

func NewMapMethodServer(conf *configs.Configs) *MapMethodServer {
	fr := conf.Rc.Devel
	fs := conf.Sc.Devel
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("Servers Configs file error:", err)
		}
	}()
	s := make(MapMethodServer, len(fr))
	for i1, server := range fr {
		serverName := server.Name
		serverHosts := fs[serverName]
        //NewMethodServer(serverName,len(serverHosts),len(server.Class))
		s[i1] = *NewMethodServer(serverName,len(serverHosts),len(server.Class))
		for i2, serverAddr := range serverHosts {
			s[i1].Client[i2] = serverAddr.Conn
			log.Println(conf.Name, "Register", fmt.Sprintf("%v_%v", serverName, i2))
		}
		for i2, class := range server.Class {
			className := class.Name
			s[i1].Methods[i2] = make([]string, len(class.Methods))
			for i3, methodName := range class.Methods {
				s[i1].Methods[i2][i3] = fmt.Sprintf("%v.%v", className, methodName)
			}
		}
	}
	return &s
}
