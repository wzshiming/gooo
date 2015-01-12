package main

import (
	"fmt"
	"gooo/configs"
	"gooo/helper"
	//"time"
)

func main() {
	mm := map[string][]byte{
		"master":  helper.OpenFile("./configs/master.json"),
		"servers": helper.OpenFile("./configs/servers.json"),
		"route":   helper.OpenFile("./configs/route.json"),
	}

	conf := configs.NewConfigs(&mm)
	conn := helper.NewConn(fmt.Sprintf("%s:%d", conf.Mc.Host, conf.Mc.Port))
	var b int
	conn.Call("Master.Stop", 222, &b)
	//time.Sleep(time.Second * 5)
}
