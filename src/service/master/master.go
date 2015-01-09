package main

import (
	"gooo/configs"
	"log"
	//"reflect"
	"fmt"
	"gooo/helper"
	"gooo/protocol"
	//"log"
	//"time"
)

type Master struct {
}

func NewMaster() *Master {
	return &Master{}
}

func (m *Master) None(args int, reply *int) error {
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
	//Build("./configs/Servers.json")
	mm := map[string][]byte{
		"master":  helper.OpenFile("./configs/master.json"),
		"servers": helper.OpenFile("./configs/servers.json"),
		"route":   helper.OpenFile("./configs/route.json"),
	}

	h := helper.NewHandeln()

	conf := configs.NewConfigs(&mm)
	configs.Name = conf.Mc.Name
	configs.Port = helper.GetPort(conf.Mc.Port)

	conf.StartServers()
	conf.StartConnect()

	var b int
	m := protocol.InitRequest{
		Conf:  *conf,
		State: 1,
	}
	for k1, v1 := range conf.Sc {
		for k2, v2 := range v1 {
			if v2.Conn == nil {
				log.Printf("%v_%v Not connect\n", k1, k2)
				continue
			}
			err := v2.Conn.Call(fmt.Sprintf("%v.Init", k1), m, &b)
			if err != nil {
				log.Println(err)
			}
		}

	}
	//for _, v := range conf.AllConnect() {
	//	err := v.Call("Status.Init", m, &b)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}

	master := NewMaster()
	h.Register(master)
	helper.EchoPortInfo(configs.Name, configs.Port)
	h.Start(configs.Port)
}
