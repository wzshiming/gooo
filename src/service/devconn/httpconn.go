package main

import (
	//"encoding/base64"
	"errors"
	"net/http"
	"time"

	"github.com/wzshiming/base"
)

type HttpConn struct {
	write    http.ResponseWriter
	read     *http.Request
	writebuf chan []byte
	readbuf  chan []byte
	pend     time.Duration
}

func NewHttpConn(max int) *HttpConn {
	return &HttpConn{
		readbuf:  make(chan []byte, max),
		writebuf: make(chan []byte, max),
		pend:     time.Second * 60,
	}
}

func (ht *HttpConn) Refresh(w http.ResponseWriter, r *http.Request, b []byte) {
	ht.write = w
	ht.read = r
	ht.readbuf <- b
	ht.push()
}

func (ht *HttpConn) ReadMsg() ([]byte, error) {
	select {
	case b, ok := <-ht.readbuf:
		if ok {
			return b, nil
		}
		return nil, errors.New("HttpConn.ReadMsg: close")
	case <-time.After(ht.pend):
		return nil, errors.New("HttpConn.ReadMsg: outtime")
	}
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
loop:
	for {
		select {
		case t, ok := <-ht.writebuf:
			if !ok {
				return errors.New("HttpConn.Close: close")
			}
			ret = append(ret, hstr(t))
		default:
			break loop
		}
	}
	t, ok := <-ht.writebuf
	if !ok {
		return errors.New("HttpConn.Close: close")
	}
	ret = append(ret, hstr(t))
	r := base.EnJson(ret)
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
	ht.pend = t.Sub(time.Now())
	return nil
}

func (ht *HttpConn) SetReadDeadline(t time.Time) error {
	return ht.SetDeadline(t)
}

func (ht *HttpConn) SetWriteDeadline(t time.Time) error {
	return ht.SetDeadline(t)
}
