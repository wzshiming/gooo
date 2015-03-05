package main

import (
	"rego/agent"
	"rego/cfg"
	"rego/defaul"

	"rego/server"
)

var ser *server.Server

func init() {
	cfg.TakeConf()
	ser = cfg.Self.Server()
}

func start() {

	ag := defaul.DefaulAgent()
	go agent.NewListener(cfg.Self.ClientPort, ag.Join).Start()
	ser.Start()
}

func main() {
	start()
}
