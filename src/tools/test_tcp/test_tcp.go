package main

import (
	//"encoding/binary"
	"fmt"
	"gooo/connser"
	"gooo/helper"
	"log"
	"service/connect/iorange"
)

type Handel struct {
	helper.HandelInterface
}

func NewHandel() *Handel {
	return &Handel{}
}

func (h *Handel) Join(c *connser.Connect) {
	h.HandelInterface.Join(c)
	sms := make([]byte, 1024)
	go func() {
		for {
			fmt.Scan(&sms)
			c.Write(sms)
		}
	}()
}

func (h *Handel) Mess(c *connser.Connect, msg []byte) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Client %v Mess error: %v", c.RemoteAddr(), err)
			helper.RecoverInfo()
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
	connser.NewServer(NewHandel(), iorange.NewIORange(1024)).StartTCP(port)
}
