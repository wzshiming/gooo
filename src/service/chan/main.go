package main

import (
	"rego/cfg"
	"rego/server"
)

var ser *server.Server

func init() {
	cfg.TakeConf()
	dbconn(cfg.Whole.Dbs["Decks"])
	ser = cfg.Self.Server()
}

func main() {
	ser.Register(NewRoom())
	ser.Start()
}
