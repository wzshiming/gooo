package main

import (
	"github.com/wzshiming/server"
	"github.com/wzshiming/server/cfg"
)

var ser *server.Server

func init() {
	cfg.TakeConf()
	dbconn(cfg.Whole.Dbs[cfg.Self.DB])
	ser = cfg.Self.Server()
}

func main() {
	ser.Register(NewRoom())
	ser.Start()
}
