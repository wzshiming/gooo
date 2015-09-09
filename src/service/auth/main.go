package main

import (
	"github.com/wzshiming/server"
	"github.com/wzshiming/server/cfg"
)

var ser *server.Server

func init() {
	cfg.TakeConf()
	dbconn(cfg.Whole.Dbs["Users"])
	ser = cfg.Self.Server()
}

func main() {
	ser.Register(NewUsers())
	ser.Start()
}
