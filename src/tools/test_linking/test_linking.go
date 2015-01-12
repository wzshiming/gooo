package main

import (
	//"bufio"
	//. "./gogogo/connect"
	//."./gogogo/coding"
	"fmt"
	"log"
	"net"
	"time"
	//"encoding/json"
)

var ls = AllLocalAddr()

func TCPDial(raddr string) (net.Conn, error) {
	laddr := <-ls
	d := net.Dialer{
		LocalAddr: laddr,
	}
	return d.Dial("tcp", raddr)
}

func AllLocalAddr() (addrs chan net.Addr) {
	addrs = make(chan net.Addr, 10)
	go func() {
		for i1 := 0; i1 != 256; i1++ {
			for i2 := 0; i2 != 256; i2++ {
				for i3 := 1; i3 != 256; i3++ {
					for i4 := 0; i4 != 65536; i4++ {
						str := fmt.Sprintf("127.%v.%v.%v:%v", i1, i2, i3, i4)
						add, err := net.ResolveTCPAddr("tcp", str)
						if err == nil {
							addrs <- add
						} else {
							break
						}
					}
				}
			}
		}
	}()
	return
}

func test() {

	conn, err := TCPDial("127.0.0.1:3000")
	if err != nil {
		return
	}
	defer conn.Close()
	b := make([]byte, 10)
	conn.Write([]byte{0, 1})
	i, _ := conn.Read(b)
	log.Println(b[:i])
}

func main() {
	for i := 0; i < 100000; i++ {
		go test()
		time.Sleep(1)
		log.Println(i)
	}
	time.Sleep(100000 * time.Second)
	log.Println("end")
}
