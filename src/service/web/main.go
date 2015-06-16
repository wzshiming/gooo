package main

import (
	"github.com/wzshiming/rego/agent"
	"github.com/wzshiming/rego/agent/defaul"
	"github.com/wzshiming/rego/cfg"
	"github.com/wzshiming/rego/server"
)

var ser *server.Server

func init() {
	cfg.TakeConf()
	ser = cfg.Self.Server()
}

func start() {

	ag := defaul.DefaulAgent()
	ser.Register(agent.NewConnect(ag))
	go run(ag)
	//go agent.NewListener(cfg.Self.ClientPort, ag.Join).Start()
	ser.Start()
}

func main() {
	start()
}
