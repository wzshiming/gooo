package main

import (
	"rego/agent"
	"rego/agent/defaul"
	"rego/cfg"
	"rego/server"
)

var ser *server.Server

func init() {
	cfg.TakeConf()
	ser = cfg.Self.Server()
}

func start() {

	ag := defaul.DefaulAgent()
	ser.Register(agent.NewConnect(ag))
	go agent.NewListener(cfg.Self.ClientPort, ag.Join).Start()
	ser.Start()
}

func main() {
	start()
}
