package route

import (
	"errors"
	"fmt"
	"gooo"
	"log"
)

type MethodServer struct {
	caller  gooo.CallServer
	Methods [][]gooo.MethodsResponse `json:"methods"`
	Name    string                   `json:"name"`
	nameuse string
}

func NewMethodServer(name string, conf *gooo.Configs) *MethodServer {

	c := gooo.NewCallServer(name, conf)

	s := MethodServer{
		caller:  *c,
		Methods: make([][]gooo.MethodsResponse, c.Size()),
		Name:    name,
		nameuse: fmt.Sprintf("use%s", name),
	}

	for i := 0; i != c.Size(); i++ {
		var mr1 gooo.MethodsResponse
		err := c.CallBy(i, "Status.Method", 0, &mr1)
		if err != nil {
			log.Fatalln(err, name)
		}
		s.Methods[i] = make([]gooo.MethodsResponse, len(mr1.Method))
		for k1, v1 := range mr1.Method {

			err := c.CallBy(i, fmt.Sprintf("%s.Method", v1), 0, &s.Methods[i][k1])
			if err != nil {
				log.Fatalln(err, name)
			}
			s.Methods[i][k1].Allow |= mr1.Allow
			s.Methods[i][k1].Unallow |= mr1.Unallow
			for k2, v2 := range s.Methods[i][k1].Method {
				s.Methods[i][k1].Method[k2] = fmt.Sprintf("%s.%s", v1, v2)
			}
		}
	}

	return &s
}

func (s *MethodServer) Call(c2 uint8, c3 uint8, args gooo.RpcRequest, reply *gooo.RpcResponse) error {

	var t map[string]uint32
	gooo.Decode(args.Session.Data, &t)
	ind := t[s.nameuse]
	m0 := s.Methods[ind]

	if c2 >= uint8(len(m0)) {
		return errors.New("MethodServer Call index c2 error")
	}
	m := m0[c2]

	if c3 >= uint8(len(m.Method)) {
		return errors.New("MethodServer Call index c3 error")
	}

	if a := m.Allow; a != (a & t["flag"]) {
		return errors.New("Permission denied")
	}

	if a := m.Unallow; 0 != (a & t["flag"]) {
		return errors.New("Permission denied")
	}

	return s.caller.CallBy(int(ind), m.Method[c3], args, reply)
}

func (s *MethodServer) Type() string {
	return s.Name
}

type MethodServers []MethodServer

func NewMethodServers(conf *gooo.Configs, names ...string) *MethodServers {
	s := make(MethodServers, len(names))
	for k, v := range names {
		s[k] = *NewMethodServer(v, conf)
	}
	return &s
}

func (s *MethodServers) Call(msg []byte, args gooo.RpcRequest, reply *gooo.RpcResponse) error {
	c1 := msg[1]
	c2 := msg[2]
	c3 := msg[3]
	//c3 := binary.BigEndian.Uint16(msg[2:4])

	//log.Println(c1, c2, c3)
	if c1 >= byte(len(*s)) {
		return errors.New("MethodServers Call index c1 error")
	}
	ss := (*s)[c1]
	return ss.Call(c2, c3, args, reply)
}

func (s *MethodServers) Map() *Remap {
	b := NewRemap()
	for k1, v1 := range *s {
		if len(v1.Methods) != 0 {
			for k2, v2 := range v1.Methods[0] {
				for k3, v3 := range v2.Method {
					b.Set(fmt.Sprintf("%s.%s", v1.Name, v3), [3]byte{byte(k1), byte(k2), byte(k3)})
				}
			}
		}
	}
	return b
}

func (s *MethodServers) MsgInfo(msg []byte) {
	//c0 := msg[0]
	c1 := msg[1]
	c2 := msg[2]
	c3 := msg[3]
	b1 := (*s)[c1]
	b2 := b1.Methods[0][c2]
	b3 := b2.Method[c3]

	log.Printf("%s.%s(%d.%d.%d)  <0x%08X^0x%08X>",
		b1.Name, b3,
		c1, c2, c3,
		b2.Allow,
		b2.Unallow)
	log.Printf(string(msg[4:]))
}
