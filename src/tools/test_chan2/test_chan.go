package main

import (
	"gooo/configs"
	"gooo/connser"
	"gooo/handeln"
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
	log.Printf("\n%s\n%s\n\n", conf.Rc.Info(msg[1], msg[2], msg[3]), msg[4:])
}

func (h *test) Join(c *connser.Connect) {
	log.Printf("%v %v Join\n", c.ToUint(), c.Conn.RemoteAddr())
	var b []byte

	b = route.ClientRequestForm(conf, "Info", "Info", "Info", infoprtc.InfoRequest{})
	c.Write(b)

	b = route.ClientRequestForm(conf, "Auth", "Auth", "Register", authprtc.RegisterRequest{
		Username: "hallo2",
		Password: "aaasssss",
	})
	c.Write(b)

	b = route.ClientRequestForm(conf, "Auth", "Auth", "LogIn", authprtc.LogInRequest{
		Username: "hallo2",
		Password: "aaasssss",
	})
	c.Write(b)

	b = route.ClientRequestForm(conf, "Info", "Use", "Chan", infoprtc.UseChanRequest{
		Use: 1,
	})
	c.Write(b)

	b = route.ClientRequestForm(conf, "Chan", "Info", "Name", chanprtc.NameRequest{
		Get: true,
	})
	c.Write(b)

	b = route.ClientRequestForm(conf, "Chan", "Chan", "Join", chanprtc.JoinRequest{
		RoomId: 0,
	})
	c.Write(b)

	b = route.ClientRequestForm(conf, "Chan", "Info", "Rooms", chanprtc.RoomsRequest{
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
