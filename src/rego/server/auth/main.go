package main

import (
	"rego/cfg"
)

func start() {
	cfg.TakeConf()
	s := cfg.Self.Server()
	s.Start()
}

func main() {
	start()
}
