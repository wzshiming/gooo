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

func (h *Auth) Register(b []byte) {
	h.regLock.Lock()
	defer h.regLock.Unlock()
	h.reg[string(b)] = true
}

func (h *Auth) Cancel(b []byte) bool {
	h.regLock.Lock()
	defer h.regLock.Unlock()
	s := string(b)
	if h.reg[s] == true {
		delete(h.reg, s)
		return true
	}
	return false
}

func (h *Auth) Mess(c *connser.Connect, msg []byte) {

	if h.Cancel(msg) {
		c.Bc = h.hand
		c.Bc.Join(c)
		log.Println(configs.Name, "join", string(msg))
	} else {
		c.Close()
		log.Println(configs.Name, "not join", string(msg))
	}
}
