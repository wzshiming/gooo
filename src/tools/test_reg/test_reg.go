package main

import (
	"gooo"
	"gooo/protocol"
	"log"
	"runtime"
	"service/connect/route"
	"time"
)

var conf = gooo.NewConfigsFrom("./conf")

type test struct {
	gooo.HandelInterface
}

func (h *test) Mess(c *gooo.Connect, msg []byte) {
	log.Printf("\n%s\n%s\n\n", conf.Rc.Info(msg[1], msg[2], msg[3]), msg[4:])
}

func (h *test) Join(c *gooo.Connect) {
	log.Printf("%v %v Join\n", c.ToUint(), c.Conn.RemoteAddr())
	var b []byte

	b = route.ClientRequestForm(conf, "Auth", "PassAuth", "LogOut", protocol.LogOutRequest{
		LogOut: true,
	})
	c.Write(b)

	//b = route.ClientRequestForm(conf, "Auth", "Auth", "Register", protocol.RegisterRequest{
	//	Username: "hallo1",
	//	Password: "aaasssss",
	//})

	//c.Write(b)

	b = route.ClientRequestForm(conf, "Auth", "Auth", "LogIn", protocol.LogInRequest{
		Username: "hallo1",
		Password: "aaasssss",
	})

	c.Write(b)
	//c.Write(b)

	//b2 := route.ClientRequestForm(conf, "Auth", "PassAuth", "LogOut", protocol.LogOutRequest{
	//	LogOut: true,
	//})
	//c.Write(b2)
	//c.Write(b)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ior := route.NewIORange(1024)
	ttt := &test{}
	gooo.NewClientTCP("127.0.0.1:3005", ior, ttt)
	time.Sleep(time.Second * 120)
}
