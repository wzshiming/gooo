package rego

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

const (
	_unknown = byte(iota)
	_json
	_gob
)

type EncodeStream []byte

func NewEncodeStream() *EncodeStream {
	return &EncodeStream{}
}

func (en *EncodeStream) ReSet() {
	*en = []byte{}
}

func (en *EncodeStream) Set(b []byte) {
	*en = b
	return
}

func (en *EncodeStream) Bytes() []byte {
	return []byte(*en)
}

func (en *EncodeStream) EnJson(s interface{}) {
	*en, _ = json.Marshal(s)
	return
}

func (en *EncodeStream) DeJson(s interface{}) {
	json.Unmarshal(*en, s)
	return
}

func (en *EncodeStream) EnGob(s interface{}) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(s); err == nil {
		*en = buf.Bytes()
	}
	return
}

func (en *EncodeStream) DeGob(s interface{}) {
	buf := bytes.NewBuffer(*en)
	dec := gob.NewDecoder(buf)
	dec.Decode(s)
	return
}
