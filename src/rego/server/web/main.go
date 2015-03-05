package main

import (
	"rego/cfg"
)

func init() {
	cfg.TakeConf()
}

func start() {
	s := cfg.Self.Server()
	s.Start()
}

func main() {
	start()
}
