package main

import (
	"rego/cfg"
	"rego/server"
)

var ser *server.Server

func init() {
	cfg.TakeConf()
	ser = cfg.Self.Server()
}

func main() {
	ser.Register(NewRoom())
	ser.Start()
}
