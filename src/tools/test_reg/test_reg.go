package main

import (
	//"bufio"
	//"encoding/binary"
	"gooo/connser"
	//"fmt"
	"gooo/configs"
	"gooo/encoder"
	"gooo/helper"
	"log"
	"runtime"
	authprtc "service/Auth/protocol"
	"service/connect/iorange"
	"time"
)

var conf = configs.NewConfigs(&map[string][]byte{
	"master":   helper.OpenFile("./configs/master.json"),
	"servers":  helper.OpenFile("./configs/servers.json"),
	"route":    helper.OpenFile("./configs/route.json"),
	"database": helper.OpenFile("./configs/database.json"),
})

type test struct {
	helper.HandelInterface
}

func (h *test) Mess(c *connser.Connect, msg []byte) {
	helper.MsgInfo(msg)
}
func (h *test) Join(c *connser.Connect) {
	log.Printf("%v %v Join\n", c.ToUint(), c.Conn.RemoteAddr())
	c1, c2, c3 := conf.Rc.FindIndex("Auth", "Auth", "ChangePwd")
	sms := make([]byte, 1024)
	sms[1] = c1
	sms[2] = c2
	sms[3] = c3
	log.Println(sms[:4])
	//binary.BigEndian.PutUint16(sms[2:4], 1)

	b, _ := encoder.Encode(authprtc.ChangePwdRequest{
		Username:    "hallo1",
		Password:    "yougeegge",
		NewPassword: "aaaaaaa",
	})
	//fmt.Printf("%s",b)
	copy(sms[4:], b)
	s := len(b) + 4
	c.Write(sms[:s])

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ior := iorange.NewIORange(1024)
	ttt := &test{}
	connser.NewClientTCP("127.0.0.1:3005", ior, ttt)
	time.Sleep(time.Second * 120)
}
