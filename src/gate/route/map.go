package route

import (
	"gooo/configs"
	"fmt"
	"log"
	"net/rpc"
    "gooo/balance"
	//"sync"
)

type MethodServer struct {
	Client  []*rpc.Client
	Name    string
    bal     balance.Balances
}

func NewMethodServer(name string,conf *configs.Configs)* MethodServer{
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("Servers Configs file error:", err)
		}
	}()
	fs := conf.Sc.Devel
	serverHosts := fs[name]
	m := MethodServer{
        Name:    name,
        Client:  make([]*rpc.Client, len(serverHosts)),
        bal:     balance.NewBalance(len(serverHosts)),
    }
	for i2, serverAddr := range serverHosts {
		m.Client[i2] = serverAddr.Conn
		log.Println(conf.Name, "Register", fmt.Sprintf("%v_%v", name, i2))
	}
    return &m
}

func (self *MethodServer) CallByName(name string,args, reply interface{}) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Call error: ", err)
		}
	}()
	return self.Client[self.bal.Allot()].Call(name, args, reply)
}
