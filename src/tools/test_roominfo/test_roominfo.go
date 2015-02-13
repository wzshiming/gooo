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
	gooo.MsgInfo(conf, msg)
}

func (h *test) Join(c *gooo.Connect) {
	log.Printf("%v %v Join\n", c.ToUint(), c.Conn.RemoteAddr())
	var b []byte

	//b = route.ClientRequestForm(conf, "Auth", "Auth", "Register", protocol.RegisterRequest{
	//	Username: "hallo111",
	//	Password: "aaasssss",
	//})
	//c.Write(b)

	b = route.ClientRequestForm(conf, "Auth", "Auth", "LogIn", protocol.LogInRequest{
		Username: "hallo111",
		Password: "aaasssss",
	})
	c.Write(b)

	b = route.ClientRequestForm(conf, "Chan", "Use", "Chan", protocol.UseChanRequest{
		Use: 1,
	})
	c.Write(b)

	b = route.ClientRequestForm(conf, "Chan", "Info", "Name", protocol.NameRequest{
		Get: true,
	})
	c.Write(b)

	//b = route.ClientRequestForm(conf, "Chan", "Chan", "Join", protocol.JoinRequest{
	//	RoomId: 0,
	//})
	//c.Write(b)

	b = route.ClientRequestForm(conf, "Chan", "Info", "Rooms", protocol.RoomsRequest{
		Get: true,
	})
	c.Write(b)

	//b = route.ClientRequestForm(conf, "Chan", "InChan", "Leave", protocol.RoomsRequest{
	//	Get: true,
	//})
	//c.Write(b)

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ior := route.NewIORange(1024)
	ttt := &test{}
	gooo.NewClientTCP("127.0.0.1:3005", ior, ttt)
	time.Sleep(time.Second * 120)
}
