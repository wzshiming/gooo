package route

import (
	"encoding/binary"
	"io"
	//"errors"
	"gooo"
)

const (
	LEN_BYTE_SIZE = 2
	BLOCK_SIZE    = 256 * 256
)

type IORange struct {
	rbuf *gooo.ReBuf
	wbuf *gooo.ReBuf
}

func NewIORange(size uint) *IORange {
	return &IORange{
		rbuf: gooo.NewReBuf(size),
		wbuf: gooo.NewReBuf(size),
	}
}

func (i *IORange) Reader(r io.Reader) (b []byte, err error) {
	ss := []byte{0, 0}
	_, err = r.Read(ss)
	if err != nil {
		panic(err)
	}
	size := binary.BigEndian.Uint16(ss)
	i.rbuf.Lock()
	defer i.rbuf.Unlock()
	_, err = r.Read(i.rbuf.Buf[:size])
	b = i.rbuf.Buf[:size]
	return
}

func (i *IORange) Writer(w io.Writer, b []byte) (err error) {
	i.wbuf.Lock()
	defer i.wbuf.Unlock()
	size := uint16(len(b))
	binary.BigEndian.PutUint16(i.wbuf.Buf[:LEN_BYTE_SIZE], size)
	copy(i.wbuf.Buf[LEN_BYTE_SIZE:LEN_BYTE_SIZE+size], b)
	_, err = w.Write(i.wbuf.Buf[:LEN_BYTE_SIZE+size])
	return
}
