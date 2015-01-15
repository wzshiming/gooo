package main

import (
	//"bufio"
	//"encoding/binary"
	//"fmt"
	"gooo/configs"
	"gooo/connser"
	"gooo/encoder"
	"gooo/helper"
	"log"
	"runtime"
	"service/connect/iorange"
	randprtc "service/random/protocol"
	"strings"
	"time"
)

var (
	mm = map[string][]byte{
		"master":   helper.OpenFile("./configs/master.json"),
		"servers":  helper.OpenFile("./configs/servers.json"),
		"route":    helper.OpenFile("./configs/route.json"),
		"database": helper.OpenFile("./configs/database.json"),
	}

	conf = configs.NewConfigs(&mm)

	Ior = iorange.NewIORange(1024)
	T1  = &test{}

	gi = 0
	gj = 0
)

type test struct {
	helper.HandelInterface
}

func (h *test) Join(c *connser.Connect) {
	log.Printf("%v %v Join gate\n", c.ToUint(), c.Conn.RemoteAddr())
	c.Write([]byte{0, 0})
}

func (h *test) Mess(c *connser.Connect, msg []byte) {
	//log.Printf("%v %v Mess  %v\n", c.ToUint(), c.Conn.RemoteAddr(), msg)
	ss := strings.Split(string(msg), " ")
	//log.Println(ss[0], ss[1])
	connser.NewClientTCP(ss[0], Ior, &test2{hand: []byte(ss[1])})
	//log.Printf("mess end ")
	gi++
	log.Printf("%d %d\r\n", gi, gj)
}

type test2 struct {
	helper.HandelInterface
	hand []byte
}

func (h *test2) Join(c *connser.Connect) {
	c.Write(h.hand)
	log.Printf("%v %v Join Connect\n", c.ToUint(), c.Conn.RemoteAddr())

	c1, c2, c3 := conf.Rc.FindIndex("Random", "Random", "Range100Spacing")
	sms := make([]byte, 1024)
	sms[1] = c1
	sms[2] = c2
	sms[3] = c3
	log.Println(sms[:4])
	//binary.BigEndian.PutUint16(sms[2:4], 1)

	b, _ := encoder.Encode(randprtc.RandRequest{10})
	//fmt.Printf("%s",b)
	copy(sms[4:], b)
	s := len(b) + 4
	//c.Write(sms[:s])
	//conf.Rc["Random"]["Random"][]
	time.Sleep(100)
	go func() {
		for i := 0; i != 1; i++ {
			//time.Sleep(1)
			c.Write(sms[:s])
		}
	}()
}
func (h *test2) Mess(c *connser.Connect, msg []byte) {
	//log.Printf("%v Mess Test %v\n", c.Conn.RemoteAddr(), msg)
	gj++
	log.Printf("%d %d\r\n", gi, gj)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	for i := 0; i != 1; i++ {
		connser.NewClientTCP("127.0.0.1:6002", Ior, T1)
	}
	log.Printf("end\n")
	time.Sleep(time.Second * 120)
}
