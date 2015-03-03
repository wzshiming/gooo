package main

import (
	"flag"
	"rego/conf"
)

var (
	dir       = defaultDir(flag.Arg(0))
	selfconf  = conf.NewServerConfig(dir + "master.json")
	wholeconf = conf.NewWholeConfig(dir + "server.json")
)

func defaultDir(dir string) string {
	if dir == "" {
		return "./conf/"
	}
	return dir
}

type Master struct {
}

func NewMaster() *Master {
	return &Master{}
}

func (ma *Master) WholeConfig(args int, reply *conf.WholeConfig) error {
	*reply = *wholeconf
	return nil
}

func start() {
	ser := selfconf.Server()
	ser.Register(NewMaster())
	ser.Start()
}

func main() {
	start()
}
