package main

import (
	//"encoding/base64"
	"errors"
	"net/http"
	"rego"
	"time"
)

type HttpConn struct {
	write    http.ResponseWriter
	read     *http.Request
	writebuf chan []byte
	readbuf  chan []byte
}

func NewHttpConn(max int) *HttpConn {
	return &HttpConn{
		readbuf:  make(chan []byte, max),
		writebuf: make(chan []byte, max),
	}
}

func (ht *HttpConn) Refresh(w http.ResponseWriter, r *http.Request, b []byte) {
	ht.write = w
	ht.read = r
	ht.readbuf <- b
	ht.push()
}

func (ht *HttpConn) ReadMsg() ([]byte, error) {

	if b, ok := <-ht.readbuf; ok {
		return b, nil
	}
	return nil, errors.New("HttpConn.ReadMsg: close")
}

func (ht *HttpConn) WriteMsg(b []byte) error {
	select {
	case ht.writebuf <- b:
		return nil
	default:
		return errors.New("HttpConn.WriteMsg: close")
	}
}
func hstr(b []byte) string {
	//return base64.StdEncoding.EncodeToString(b)
	return string(b)
}
func (ht *HttpConn) Close() error {
	close(ht.writebuf)
	close(ht.readbuf)
	return nil
}
func (ht *HttpConn) push() error {
	ret := []string{}
	ret = append(ret, hstr(<-ht.writebuf))
loop:
	for {
		select {
		case t, ok := <-ht.writebuf:
			ret = append(ret, hstr(t))
			if !ok {
				return errors.New("HttpConn.Close: close")
			}
		default:
			break loop
		}
	}
	r := rego.EnJson(ret)
	ht.write.Write(r.Bytes())
	return nil
}

func (ht *HttpConn) LocalAddr() string {
	return ht.read.Host
}

func (ht *HttpConn) RemoteAddr() string {
	return ht.read.RemoteAddr
}

func (ht *HttpConn) SetDeadline(t time.Time) error {
	return nil
}

func (ht *HttpConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (ht *HttpConn) SetWriteDeadline(t time.Time) error {
	return nil
}
