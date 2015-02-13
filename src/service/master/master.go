package main

import (
	"gooo"
)

type Master struct {
	conf *gooo.Configs
	hand *gooo.Handeln
}

func NewMaster(conf *gooo.Configs, hand *gooo.Handeln) *Master {
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
//    conf := gooo.GetServersConfig(path)
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
	defer gooo.Recover()
	conf := gooo.NewConfigsFrom("./conf")
	gooo.Name = conf.Master().Name
	gooo.Port = gooo.GetPort(conf.Master().Port)

	conf.StartServers()
	//conf.StartConnect()
	var b int
	m := gooo.InitRequest{
		Conf:  *conf,
		State: 1,
	}
	conf.Foreach("Status", "Init", m, &b)
	//time.Sleep(time.Second)
	//conf.Foreach("", "Init", m, &b)
	//conf.Foreach("", "Start", m, &b)

	h := gooo.NewHandeln()
	master := NewMaster(conf, h)
	h.Register(master)
	gooo.EchoPortInfo(gooo.Name, gooo.Port)
	h.Start(gooo.Port)
	//gooo.Prevent()
}
