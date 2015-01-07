package main

import (

	//"log"
	"gooo/configs"
	//"reflect"
	"gooo/helper"
	"gooo/protocol"
	"fmt"
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
//    conf := configs.GetServersConfig(path).Devel
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
	//Build("./configs/Servers.json")
	mm := map[string][]byte{
		"master":  helper.OpenFile("./configs/master.json"),
		"servers": helper.OpenFile("./configs/servers.json"),
		"route":   helper.OpenFile("./configs/route.json"),
	}

	h := helper.NewHandeln()

	var mc configs.MasterConfig
	helper.GetConfig(mm["master"], &mc)

	configs.Name = "Master"
	configs.Port = helper.GetPort(mc.Devel.Port)
	conf := configs.NewConfigs(&mm)
	conf.StartServers()
	conf.StartConnect()

	var b int
	m := protocol.InitRequest{
		Conf:  mm,
		State: 1,
	}
	for _, v := range conf.AllConnect() {
		err := v.Call("Status.Init", m, &b)
		if err != nil {
			fmt.Println(err)
		}

	}

	master := NewMaster()
	h.Register(master)
	helper.EchoPortInfo(configs.Name, configs.Port)
	h.Start(configs.Port)
}
