package main

import (
	"gooo"
	"gooo/protocol"
	"gooo/route"
	"log"
	"runtime"
	"time"
)

var conf = gooo.NewConfigsFrom("./conf")

type test struct {
	gooo.HandelInterface
}

var msf = route.NewRemapFrom("./conf/Connect_0_map.json")

func (h *test) Mess(c *gooo.Connect, msg []byte) {
	msf.MsgInfo(msg)
}

func (h *test) Join(c *gooo.Connect) {
	log.Printf("%v %v Join\n", c.ToUint(), c.Conn.RemoteAddr())
	var b []byte

	b = msf.ClientRequest("Auth.Auth.Register", protocol.RegisterRequest{
		Username: "hallo111",
		Password: "aaassss1s",
	})
	c.Write(b)

	b = msf.ClientRequest("Auth.Auth.LogIn", protocol.LogInRequest{
		Username: "hallo111",
		Password: "aaassss1s",
	})
	c.Write(b)

	b = msf.ClientRequest("Chan.Use.Chan", protocol.UseChanRequest{
		Use: 1,
	})
	c.Write(b)

	b = msf.ClientRequest("Chan.Info.Name", protocol.NameRequest{
		Get: true,
	})
	c.Write(b)

	//b = msf.ClientRequest("Chan.Chan.Join", protocol.JoinRequest{
	//	RoomId: 0,
	//})
	//c.Write(b)

	b = msf.ClientRequest("Chan.Info.Rooms", protocol.RoomsRequest{
		Get: true,
	})
	c.Write(b)

	//b = msf.ClientRequest("Chan.InChan.Leave", protocol.RoomsRequest{
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
