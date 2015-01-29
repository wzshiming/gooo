package main

import (
	"gooo/configs"
	"gooo/connser"
	"gooo/handeln"
	//"gooo/helper"
	"log"
	"runtime"
	authprtc "service/Auth/protocol"
	"service/connect/iorange"
	"service/connect/route"
	"time"
)

var conf = configs.NewConfigsFrom("./conf")

type test struct {
	handeln.HandelInterface
}

func (h *test) Mess(c *connser.Connect, msg []byte) {
	log.Printf("\n%s\n%s\n\n", conf.Rc.Info(msg[1], msg[2], msg[3]), msg[4:])
}

func (h *test) Join(c *connser.Connect) {
	log.Printf("%v %v Join\n", c.ToUint(), c.Conn.RemoteAddr())
	var b []byte
	b = route.ClientRequestForm(conf, "Auth", "Info", "Info", authprtc.InfoRequest{})
	c.Write(b)

	b = route.ClientRequestForm(conf, "Auth", "PassAuth", "LogOut", authprtc.LogOutRequest{
		LogOut: true,
	})
	c.Write(b)

	//b = route.ClientRequestForm(conf, "Auth", "Auth", "Register", authprtc.RegisterRequest{
	//	Username: "hallo1",
	//	Password: "aaasssss",
	//})

	//c.Write(b)

	b = route.ClientRequestForm(conf, "Auth", "Auth", "LogIn", authprtc.LogInRequest{
		Username: "hallo1",
		Password: "aaasssss",
	})

	c.Write(b)
	//c.Write(b)

	//b2 := route.ClientRequestForm(conf, "Auth", "PassAuth", "LogOut", authprtc.LogOutRequest{
	//	LogOut: true,
	//})
	//c.Write(b2)
	//c.Write(b)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ior := iorange.NewIORange(1024)
	ttt := &test{}
	connser.NewClientTCP("127.0.0.1:3005", ior, ttt)
	time.Sleep(time.Second * 120)
}
