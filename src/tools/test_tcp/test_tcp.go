package main

import (
	//"encoding/binary"
	"fmt"
	"gooo"
	"log"
	"service/connect/route"
)

type Handel struct {
	gooo.HandelInterface
}

func NewHandel() *Handel {
	return &Handel{}
}

func (h *Handel) Join(c *gooo.Connect) {
	h.HandelInterface.Join(c)
	sms := make([]byte, 1024)
	go func() {
		for {
			fmt.Scan(&sms)
			c.Write(sms)
		}
	}()
}

func (h *Handel) Mess(c *gooo.Connect, msg []byte) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Client %v Mess error: %v", c.RemoteAddr(), err)
			gooo.RecoverInfo()
		}
	}()
	c0 := msg[0]
	c1 := msg[1]
	c2 := msg[2]
	c3 := msg[3]
	log.Println("from:", c.RemoteAddr().String())
	log.Println("flag:", c0)
	log.Println("type:", c1)
	log.Println("sub:", c2)
	log.Println("code:", c3)
	log.Println("info:", string(msg[4:]))
	//h.HandelInterface.Mess(c, msg)
	//c.Write(msg)
}

func main() {
	port := ":3005"
	log.Println("port", port)
	gooo.NewServer(NewHandel(), route.NewIORange(1024)).StartTCP(port)
}
