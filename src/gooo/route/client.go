package route

import (
	"github.com/wzshiming/ffmt"
	"gooo"
	"log"
)

type Remap map[string]uint32

func NewRemap() *Remap {
	r := make(Remap)
	return &r
}

func NewRemapFrom(path string) *Remap {
	var s Remap
	gooo.Decode(gooo.ReadFile(path), &s)
	return &s
}

func (s *Remap) Get(name string) [3]byte {
	t := (*s)[name]
	return [3]byte{byte(t), byte(t >> 8), byte(t >> 16)}
}

func (s *Remap) Set(name string, v [3]byte) {
	(*s)[name] = (uint32(v[2]) << 16) | (uint32(v[1]) << 8) | uint32(v[0])
}

var lo byte = 0

func (s *Remap) ClientRequest(name string, arge interface{}) []byte {
	code := s.Get(name)
	if code[0] == 0 && code[1] == 0 && code[2] == 0 {
		return []byte{}
	}
	return s.ClientRequestCode(code, arge)
}

func (s *Remap) ClientRequestCode(code [3]byte, arge interface{}) []byte {
	lo++
	sms := append([]byte{lo}, code[:]...)
	b, err := gooo.Encode(arge)
	if err != nil {
		return []byte{}
	}
	return append(sms, b...)
}

func (s *Remap) MsgInfo(msg []byte) {
	c0 := msg[0]
	c1 := msg[1]
	c2 := msg[2]
	c3 := msg[3]

	log.Printf("%d %d.%d.%d", c0, c1, c2, c3)
	log.Printf(string(msg[4:]))
}

func (s *Remap) WriteFile(path string) {
	b, _ := gooo.Encode(s)
	gooo.WriteFile(path, []byte(ffmt.Sputs(string(b))))
}
