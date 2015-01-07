package addr

import (
	"net"
	"strings"
)

func GetIP() string {
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		return "127.0.0.1"
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}
