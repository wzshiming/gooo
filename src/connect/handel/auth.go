package handel

import (
	"gooo/connser"
	"log"
	"sync"
)

type Auth struct {
	hand    connser.EventHandel
	reg 	map[string]bool
	regLock sync.Mutex
}

func NewAuth(e connser.EventHandel) *Auth {
	return &Auth{
		hand:  e,
		reg:   make(map[string]bool),
	}
}

func (h *Auth) Register(b []byte){
	h.regLock.Lock()
	defer h.regLock.Unlock()
	h.reg[string(b)] = true
}

func (h *Auth) Cancel(b []byte)(o bool){
	h.regLock.Lock()
	defer h.regLock.Unlock()
	s := string(b)
	if h.reg[s] == true {
		delete(h.reg,s)
		return true
	}
	return false
}


func (h *Auth) Join(c *connser.Connect) {
	log.Printf("%v %v join\n", c.ToUint(), c.RemoteAddr())
}
func (h *Auth) Mess(c *connser.Connect, msg []byte) {
	log.Println(1231231,string(msg))
	if h.Cancel(msg){
		c.Bc = h.hand
		c.Bc.Join(c)
	}else{
		c.Close()
	}
}
func (h *Auth) Exit(c *connser.Connect) {
	log.Printf("%v is quiting\n", c.RemoteAddr())
}
func (h *Auth) Timeout(c *connser.Connect) {
	log.Printf("%v Timeout\n", c.RemoteAddr())
}

func (h *Auth) Recover(c *connser.Connect, err error) {
	log.Printf("%v error %v\n", c.RemoteAddr(), err)
}
