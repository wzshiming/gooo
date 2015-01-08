package main

import (
	//"bufio"
	"encoding/binary"
	"gooo/connser"
	//"fmt"
	"gooo/encoder"
	"gooo/helper"
	"gooo/protocol"
	"log"
	"runtime"
	"service/connect/iorange"
	"time"
)

type test struct {
	helper.HandelInterface
}

func (h *test) Join(c *connser.Connect) {
	log.Printf("%v %v Join\n", c.ToUint(), c.Conn.RemoteAddr())
	sms := make([]byte, 1024)
	sms[0] = 1
	sms[1] = 1
	binary.BigEndian.PutUint16(sms[2:4], 1)

	b, _ := encoder.Encode(protocol.RandRequest{10})
	//fmt.Printf("%s",b)
	copy(sms[4:], b)
	s := len(b) + 4
	//time.Sleep(100)
	go func() {
		for i := 0; i != 100000; i++ {
			//time.Sleep(1)
			c.Write(sms[:s])
		}
	}()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ior := iorange.NewIORange(1024)
	ttt := &test{}
	for i := 0; i < 100; i++ {
		connser.NewClientTCP("127.0.0.1:3005", ior, ttt)
	}
	log.Printf("end\n")
	time.Sleep(time.Second * 120)
}
