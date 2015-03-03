package main

import (
	"github.com/gorilla/websocket"
	"gooo"
	"io"
	"time"
)

type Websocket struct {
	*websocket.Conn
}

func NewWebsocket(conn *websocket.Conn) *Websocket {
	return &Websocket{Conn: conn}
}

func (s *Websocket) Write(w []byte) (n int, err error) {
	return len(w), s.Conn.WriteMessage(websocket.TextMessage, w)
}

func (s *Websocket) Read(r []byte) (n int, err error) {
	var b []byte
	_, b, err = s.Conn.ReadMessage()
	if err != nil {
		return 0, err
	}
	copy(r, b)
	return len(b), nil
}

func (s *Websocket) SetDeadline(t time.Time) error {
	return s.Conn.UnderlyingConn().SetDeadline(t)
}

type IOFF struct {
	rbuf *gooo.ReBuf
}

func NewIOFF(size uint) *IOFF {
	return &IOFF{
		rbuf: gooo.NewReBuf(size),
	}
}

func (i *IOFF) Reader(r io.Reader) (b []byte, err error) {
	var j int
	i.rbuf.Lock()
	defer i.rbuf.Unlock()
	j, err = r.Read(i.rbuf.Buf)
	return i.rbuf.Buf[:j], err
}

func (i *IOFF) Writer(w io.Writer, b []byte) (err error) {
	_, err = w.Write(b)
	return
}
