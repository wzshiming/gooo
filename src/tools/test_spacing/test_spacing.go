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
	gooo.MsgInfo(msg)
}

func (h *test) Join(c *gooo.Connect) {
	log.Printf("%v %v Join\n", c.ToUint(), c.Conn.RemoteAddr())
	b := route.ClientRequestForm(conf, "Random", "Random", "Range100Spacing", protocol.SpacingRequest{
		Size:  10,
		Space: 10000,
	})

	go func() {
		for i := 0; i != 5; i++ {
			c.Write(b)
		}
	}()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ior := route.NewIORange(1024)
	ttt := &test{}
	for i := 0; i != 2; i++ {
		gooo.NewClientTCP("127.0.0.1:3005", ior, ttt)
	}
	log.Printf("end\n")
	time.Sleep(time.Second * 120)
}
