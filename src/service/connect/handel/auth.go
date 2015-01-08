package handel

import (
	"gooo/configs"
	"gooo/connser"
	"gooo/helper"
	"log"
	"sync"
)

type Auth struct {
	helper.HandelInterface
	hand    connser.EventHandel
	reg     map[string]bool
	regLock sync.Mutex
}

func NewAuth(e connser.EventHandel) *Auth {
	return &Auth{
		hand: e,
		reg:  make(map[string]bool),
	}
}

func (h *Auth) Register(b string) {
	h.regLock.Lock()
	defer h.regLock.Unlock()
	h.reg[b] = true
}

func (h *Auth) Cancel(b string) bool {
	h.regLock.Lock()
	defer h.regLock.Unlock()
	if h.reg[b] == true {
		delete(h.reg, b)
		return true
	}
	return false
}

func (h *Auth) Mess(c *connser.Connect, msg []byte) {

	if h.Cancel(string(msg)) {
		c.Bc = h.hand
		c.Bc.Join(c)
		log.Println(configs.Name, "join", c.ToUint(), string(msg))
	} else {
		c.Close()
		log.Println(configs.Name, "not join", c.ToUint(), string(msg))
	}
}
