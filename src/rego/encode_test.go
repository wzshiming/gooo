package rego

import (
	"testing"
)

func Test_encode(t *testing.T) {
	es := NewEncodeStream()
	s := []byte(`{"asdasd":"aaa"}`)
	es.Set(s)
	m := make(map[string]string)
	es.DnJson(&m)
	es.ReSet()
	es.EnJson(m)
	if string(s) != string(es.Bytes()) {
		t.Fatal(s, es.Bytes())
	}

	m2 := make(map[string]string)
	es.EnGob(m)
	es.DnGob(&m2)
	t.Log(m2)
}
