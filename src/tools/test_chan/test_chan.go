package main

import (
	"gooo/configs"
	"gooo/connser"
	"gooo/handeln"
	"gooo/helper"
	"log"
	"runtime"
	authprtc "service/Auth/protocol"
	chanprtc "service/chan/protocol"
	"service/connect/iorange"
	"service/connect/route"
	infoprtc "service/info/protocol"
	"time"
)

var conf = configs.NewConfigsFrom("./conf")

type test struct {
	handeln.HandelInterface
}

func (h *test) Mess(c *connser.Connect, msg []byte) {
	helper.MsgInfo(msg)
}

func (h *test) Join(c *connser.Connect) {
	log.Printf("%v %v Join\n", c.ToUint(), c.Conn.RemoteAddr())
	var b []byte

	b = route.ClientRequestForm(conf, "Info", "Info", "Info", infoprtc.InfoRequest{})
	c.Write(b)

	b = route.ClientRequestForm(conf, "Auth", "Auth", "LogIn", authprtc.LogInRequest{
		Username: "hallo1",
		Password: "aaasssss",
	})
	c.Write(b)

	b = route.ClientRequestForm(conf, "Info", "Use", "Chan", infoprtc.UseChanRequest{
		Use: 1,
	})
	c.Write(b)

	b = route.ClientRequestForm(conf, "Chan", "Chan", "Name", chanprtc.NameRequest{
		Get: true,
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