package main

import (
	"fmt"
	"gooo"
	//"time"
)

func main() {

	conf := gooo.NewConfigsFrom("./conf")
	conn := gooo.NewConnTcp(fmt.Sprintf("%s:%d", conf.Mc.Host, conf.Mc.Port))
	var b int
	conn.Call("Master.Stop", 222, &b)
	//time.Sleep(time.Second * 5)
}
