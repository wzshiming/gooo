package main

import (
	"github.com/wzshiming/base"
	"github.com/wzshiming/ffmt"
	"github.com/wzshiming/server"
	"github.com/wzshiming/server/cfg"
)

var ser *server.Server

func init() {
	cfg.TakeConf()
	ser = cfg.Self.Server()
}

func start() {

	//ag := defaul.DefaulAgent()
	//ser.Register(agent.NewConnect(ag))
	//go agent.NewListener(cfg.Self.ClientPort, ag.Join).Start()

	h := base.FilesHashIndexing("../static/web")
	base.SaveLocal("../static/versions.json", []byte(ffmt.Fmt(string(base.EnJson(h).Bytes()))))
	base.NOTICE("Start from port", cfg.Self.ClientPort)
	go filesServer("../static", cfg.Self.ClientPort)
	ser.Start()
}

func main() {
	start()
}
