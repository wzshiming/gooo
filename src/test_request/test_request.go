package main

import (
	//"bufio"
	"gooo/connser"
	"encoding/binary"
	//"fmt"
	"gooo/protocol"
	"gooo/encoder"
	"connect/iorange"
	"log"
	"runtime"
	"time"
)

type test struct {
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
func (h *test) Mess(c *connser.Connect, msg []byte) {
	log.Printf("%v %v Mess  %v\n", c.ToUint(), c.Conn.RemoteAddr(), msg)
}
func (h *test) Exit(c *connser.Connect) {
	log.Printf("%v %v Exit\n", c.ToUint(), c.Conn.RemoteAddr())
}
func (h *test) Timeout(c *connser.Connect) {
	log.Printf("%v %v Timeout\n", c.ToUint(), c.Conn.RemoteAddr())
}
func (h *test) Recover(c *connser.Connect, err error) {
	log.Printf("%v error %v\n", c.Conn.RemoteAddr(), err)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ior := iorange.NewIORange(1024)
	ttt := &test{}
	for i := 0; i < 100; i++ {
		connser.NewClientTCP("121.42.149.4:3000", ior, ttt)
	}
	log.Printf("end\n")
	time.Sleep(time.Second * 120)
}
