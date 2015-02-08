package gooo

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"path"
	"regexp"
	"strings"
)

func NewConnTcp(addr string) (c *rpc.Client) {
	c, err := jsonrpc.Dial("tcp", addr)
	if err != nil {
		log.Fatal("Server tcp Connect error:", err)
	}
	return
}

func NewConnUdp(addr string) (c *rpc.Client) {
	c, err := jsonrpc.Dial("udp", addr)
	if err != nil {
		log.Fatal("Server udp Connect error:", err)
	}
	return
}

func OpenFile(path string) []byte {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

var (
	reg1 = regexp.MustCompile(`//.*\r`)
	reg2 = regexp.MustCompile(`#.*\r`)
	reg3 = regexp.MustCompile(`\s`)
	reg4 = regexp.MustCompile(`/\*.*\*/`)
	reg5 = regexp.MustCompile(`\b0[xX][0-9a-fA-F]+\b`)
)

func ReplaceJson(str []byte) []byte {
	str = reg1.ReplaceAll(str, []byte{})
	str = reg2.ReplaceAll(str, []byte{})
	str = reg3.ReplaceAll(str, []byte{})
	str = reg4.ReplaceAll(str, []byte{})
	str = reg5.ReplaceAllFunc(str, func(b []byte) []byte {
		var p uint32
		fmt.Sscanf(string(b[2:]), "%x", &p)
		return []byte(fmt.Sprintln(p))
	})
	return str
}

func GetConfig(conf []byte, s interface{}) {
	err := json.Unmarshal(conf, s)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func GetBaseName() string {
	return path.Base(os.Args[0])
}

func EchoPortInfo(name, port string) {
	log.Printf("%s Server starts from port %s\n", name, port)
}

func EchoPublicPortInfo(name, port string) {
	log.Printf("%s Server starts from public port %s\n", name, port)
}

func FlagVar() (name, port, typ string, id uint64) {
	inputType := flag.String("t", "", "Type")
	inputPort := flag.Int("p", 0, "Port")
	inputId := flag.Uint64("i", 0, "Id")
	flag.Parse()
	typ = fmt.Sprintf("%s", *inputType)
	port = fmt.Sprintf(":%d", *inputPort)
	name = fmt.Sprintf("%s_%d", *inputType, *inputId)
	id = *inputId
	if typ == "" {
		name = ""
	}
	return
}

func GetPort(port int) string {
	return fmt.Sprintf(":%d", port)
}
