package route

import (
	"rego"
	"rego/server"
	"testing"
)

func Test_code(t *testing.T) {
	cm := NewCodeMaps()
	cm.Append("he", server.Classs{
		server.Methods{
			Name: "ll",
			Methods: []string{
				"oo",
				"o",
			},
		},
	})
	rego.INFO(cm)
	rego.INFO(cm.MakeReCodeMap().Map("he.ll.o"))
}
