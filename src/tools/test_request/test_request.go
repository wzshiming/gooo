package main

import (
	"gooo/configs"
	"gooo/connser"
	"gooo/handeln"
	"gooo/helper"
	"log"
	"runtime"
	"service/connect/iorange"
	"service/connect/route"
	randprtc "service/random/protocol"
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
	b := route.ClientRequestForm(conf, "Random", "Random", "Range100", randprtc.RandRequest{
		Size: 10,
	})

	go func() {
		for i := 0; i != 100; i++ {
			c.Write(b)
		}
	}()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ior := iorange.NewIORange(1024)
	ttt := &test{}
	for i := 0; i != 100; i++ {
		connser.NewClientTCP("127.0.0.1:3005", ior, ttt)
	}
	log.Printf("end\n")
	time.Sleep(time.Second * 120)
}
