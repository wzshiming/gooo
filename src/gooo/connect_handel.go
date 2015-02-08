package gooo

import (
	"log"
)

type HandelInterface struct {
}

func (h *HandelInterface) Join(c *Connect) {
	log.Printf("%v join\n", c.RemoteAddr())
}

func (h *HandelInterface) Mess(c *Connect, msg []byte) {
	log.Printf("%v Mess %v", c.RemoteAddr(), msg)
}

func (h *HandelInterface) Exit(c *Connect) {
	log.Printf("%v quit\n", c.RemoteAddr())
}

func (h *HandelInterface) Timeout(c *Connect) {
	log.Printf("%v Timeout\n", c.RemoteAddr())
}

func (h *HandelInterface) Recover(c *Connect, err error) {
	log.Printf("%v error %v\n", c.RemoteAddr(), err)
}
