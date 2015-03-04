package cfg

import (
	"rego"
	"testing"
)

func Test_read(t *testing.T) {
	b := NewWholeConfig("../conf/server.json")
	if b == nil {
		t.Fail()
	}
	c := NewServerConfig("../conf/master.json")
	if c == nil {
		t.Fail()
	}
	rego.INFO(b)
	rego.INFO(c)
}
