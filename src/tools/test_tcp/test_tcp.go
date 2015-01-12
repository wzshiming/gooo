package main

import (
	"encoding/binary"
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
	c1 := msg[0]
	c2 := msg[1]
	c3 := binary.BigEndian.Uint16(msg[2:4])
	log.Println("from:", c.RemoteAddr().String())
	log.Println("type:", c1)
	log.Println("subtype:", c2)
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
