package main

import (
	"gooo/configs"
	//"log"
	//"reflect"
	//"fmt"
	"gooo/handeln"
	"gooo/helper"
	"gooo/protocol"
	//"time"
)

type Master struct {
	conf *configs.Configs
	hand *handeln.Handeln
}

func NewMaster(conf *configs.Configs, hand *handeln.Handeln) *Master {
	return &Master{
		conf: conf,
		hand: hand,
	}
}

func (s *Master) Stop(args int, reply *int) error {
	if args == 222 {
		var b int
		s.conf.Foreach("Status", "Stop", 222, &b)
		s.hand.Stop()
	}
	return nil
}

//func Build(path string){
//    conf := configs.GetServersConfig(path)
//    for k1,_ := range conf {
//        b := fmt.Sprintf("%s.go",k1)
//        args := []string{"build", b}
//        cmd := exec.Command("go", args...)
//        cmd.Stdout = os.Stdout
//        cmd.Stderr = os.Stderr
//        cmd.Run()
//    }
//}

func main() {
	defer helper.Recover()
	conf := configs.NewConfigsFrom("./conf")
	configs.Name = conf.Mc.Name
	configs.Port = helper.GetPort(conf.Mc.Port)

	conf.StartServers()
	//conf.StartConnect()
	var b int
	m := protocol.InitRequest{
		Conf:  *conf,
		State: 1,
	}
	conf.Foreach("Status", "Init", m, &b)
	//time.Sleep(time.Second)
	conf.Foreach("", "Init", m, &b)
	conf.Foreach("", "Start", m, &b)

	h := handeln.NewHandeln()
	master := NewMaster(conf, h)
	h.Register(master)
	helper.EchoPortInfo(configs.Name, configs.Port)
	h.Start(configs.Port)
}
