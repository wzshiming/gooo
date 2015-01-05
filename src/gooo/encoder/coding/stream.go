package coding

import (
	"encoding/binary"
	"io"
	"bytes"
	"sync"

)

type HeadSize uint16
// A Decoder reads and decodes coding objects from an input stream.
var gdata = make([]byte,1024*1024)
var gbuf = bytes.NewBuffer([]byte{})
var dmutex,emutex sync.Mutex
type Decoder struct {
	r    io.Reader
}

// NewDecoder returns a new decoder that reads from r.
//
// The decoder introduces its own buffering and may
// read data from r beyond the coding values requested.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r:r}
}

// Decode reads the next coding-encoded value from its
// input and stores it in the value pointed to by v.
//
// See the documentation for Unmarshal for details about
// the conversion of coding into a Go value.
func (dec *Decoder) Decode(v interface{}) error {

	// the connection is still usable since we read a complete coding
	// object from it before the error happened.
	size := HeadSize(0)
	binary.Read(dec.r, binary.BigEndian, &size)
	dmutex.Lock()
	defer dmutex.Unlock()
	data := gdata[:size]
    dec.r.Read(data)
    data,err := Decompr(data)
    if err!=nil {
		return err
	}
	err = Unmarshal(data,v)
	return err
}


// An Encoder writes coding objects to an output stream.
type Encoder struct {
	w   io.Writer
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

// Encode writes the coding encoding of v to the stream,
// followed by a newline character.
//
// See the documentation for Marshal for details about the
// conversion of Go values to coding.

func (enc *Encoder) Encode(v interface{}) error {
	d,err := Marshal(v)
	if err != nil {
		return err
	}
	d, err = Compr(d)
	if err!=nil {
		return err
	}
	emutex.Lock()
	defer emutex.Unlock()
	gbuf.Reset()
	binary.Write(gbuf, binary.BigEndian, HeadSize(len(d)))
	gbuf.Write(d)

	_, err = gbuf.WriteTo(enc.w)
	return err
}

