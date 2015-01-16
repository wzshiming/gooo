package main

import (
	"fmt"
	"gooo/configs"
	"gooo/helper"
	//"time"
)

func main() {

	conf := configs.NewConfigsFrom("./conf")
	conn := helper.NewConn(fmt.Sprintf("%s:%d", conf.Mc.Host, conf.Mc.Port))
	var b int
	conn.Call("Master.Stop", 222, &b)
	//time.Sleep(time.Second * 5)
}
