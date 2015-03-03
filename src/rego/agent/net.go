package agent

import (
	"bufio"
	"encoding/binary"
	"net"
)

type ConnNet struct {
	net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
}

func NewConnNet(conn net.Conn) *ConnNet {
	if conn == nil {
		return nil
	}
	ret := ConnNet{
		Conn:   conn,
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
	}
	return &ret
}

func (conn *ConnNet) ReadMsg() (b []byte, err error) {
	head := []byte{0, 0}
	_, err = conn.reader.Read(head)
	if err != nil {
		return nil, err
	}
	size := binary.BigEndian.Uint16(head)
	body := make([]byte, size)
	_, err = conn.reader.Read(body)
	return body, err
}

func (conn *ConnNet) WriteMsg(b []byte) (err error) {
	size := uint16(len(b))
	err = binary.Write(conn.writer, binary.BigEndian, size)
	if err != nil {
		return err
	}
	_, err = conn.writer.Write(b)
	if err != nil {
		return err
	}
	return conn.writer.Flush()
}
