package conf

import (
	"rego"
	"testing"
)

func Test_read(t *testing.T) {
	b := NewWholeConfig("./test/server.json")
	if b == nil {
		t.Fail()
	}
	c := NewServerConfig("./test/master.json")
	if c == nil {
		t.Fail()
	}
	rego.INFO(b)
	rego.INFO(c)
}
