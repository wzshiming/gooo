package main

import (
	//"bufio"
	//"encoding/binary"
	"gooo/connser"
	//"fmt"
	"gooo/configs"
	//"gooo/encoder"
	"gooo/helper"

	"log"
	"runtime"
	authprtc "service/Auth/protocol"
	"service/connect/iorange"
	"service/connect/route"
	"time"
)

var conf = configs.NewConfigsFrom("./configs")

type test struct {
	helper.HandelInterface
}

func (h *test) Mess(c *connser.Connect, msg []byte) {
	helper.MsgInfo(msg)
}

func (h *test) Join(c *connser.Connect) {
	log.Printf("%v %v Join\n", c.ToUint(), c.Conn.RemoteAddr())
	b := route.ClientRequestForm(conf, "Auth", "Auth", "ChangePwd", authprtc.ChangePwdRequest{
		Username:    "hallo1",
		Password:    "aaaaaaa",
		NewPassword: "aaasssss",
	})

	c.Write(b)

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ior := iorange.NewIORange(1024)
	ttt := &test{}
	connser.NewClientTCP("127.0.0.1:3005", ior, ttt)
	time.Sleep(time.Second * 120)
}
