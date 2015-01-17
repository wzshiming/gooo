package handeln

import (
	"gooo/connser"
	"log"
)

type HandelInterface struct {
}

func (h *HandelInterface) Join(c *connser.Connect) {
	log.Printf("%v join\n", c.RemoteAddr())
}

func (h *HandelInterface) Mess(c *connser.Connect, msg []byte) {
	log.Printf("%v Mess %v", c.RemoteAddr(), msg)
}

func (h *HandelInterface) Exit(c *connser.Connect) {
	log.Printf("%v quit\n", c.RemoteAddr())
}

func (h *HandelInterface) Timeout(c *connser.Connect) {
	log.Printf("%v Timeout\n", c.RemoteAddr())
}

func (h *HandelInterface) Recover(c *connser.Connect, err error) {
	log.Printf("%v error %v\n", c.RemoteAddr(), err)
}
