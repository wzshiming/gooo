package helper

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
	"strings"
)

func NewConn(addr string) (c *rpc.Client) {
	c, err := jsonrpc.Dial("tcp", addr)
	if err != nil {
		log.Fatal("Server Connect error:", err)
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

func FlagVar() (name, port, typ string, id int) {
	inputType := flag.String("t", "", "Type")
	inputPort := flag.Int("p", 0, "Port")
	inputId := flag.Int("i", 0, "Id")
	flag.Parse()
	name = fmt.Sprintf("%s", *inputType)
	port = fmt.Sprintf(":%d", *inputPort)
	typ = fmt.Sprintf("%s_%d", *inputType, *inputId)
	id = *inputId
	return
}

func GetPort(port int) string {
	return fmt.Sprintf(":%d", port)
}
