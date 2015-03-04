package main

import (
	"rego/cfg"
)

func init() {
	cfg.Whole = cfg.NewWholeConfig(cfg.DirConf + "server.json")
	cfg.Self = cfg.Master
}

func main() {
	cfg.Whole.Shutdown()
	cfg.Master.Client().ShutdownNow()
}
