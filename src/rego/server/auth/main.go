package main

import (
	"rego/cfg"
	"rego/server"
)

var ser *server.Server

func init() {
	cfg.TakeConf()
	dbconn(cfg.Whole.Dbs["Users"])
	ser = cfg.Self.Server()
}

func main() {
	ser.Register(NewAuth())
	ser.Register(NewUsers())
	ser.Start()
}
