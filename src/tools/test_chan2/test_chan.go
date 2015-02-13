package main

import (
	"fmt"
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
	index int
}

func (h *test) Mess(c *gooo.Connect, msg []byte) {
	gooo.MsgInfo(conf, msg)
	//log.Printf("\n%s\n%s\n\n", conf.Rc.Info(msg[1], msg[2], msg[3]), msg[4:])
}

func (h *test) Join(c *gooo.Connect) {
	log.Printf("%v %v Join\n", c.ToUint(), c.Conn.RemoteAddr())
	var b []byte

	cc := fmt.Sprintf("hallo%d", h.index)

	b = route.ClientRequestForm(conf, "Auth", "Auth", "Register", protocol.RegisterRequest{
		Username: cc,
		Password: "aaasssss",
	})
	c.Write(b)

	b = route.ClientRequestForm(conf, "Auth", "Auth", "LogIn", protocol.LogInRequest{
		Username: cc,
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

	b = route.ClientRequestForm(conf, "Chan", "Chan", "Join", protocol.JoinRequest{
		RoomId: 0,
	})
	c.Write(b)

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
	for i := 10; i != 20; i++ {
		ttt := &test{index: i}
		gooo.NewClientTCP("127.0.0.1:3005", ior, ttt)
	}
	time.Sleep(time.Second * 120)
}
