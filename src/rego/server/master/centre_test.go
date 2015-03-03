package main

import (
	"rego"
	"rego/conf"
	"testing"
)

func Test_get(t *testing.T) {
	go start()
	var w conf.WholeConfig
	selfconf.Client().Call("Master.WholeConfig", 0, &w)
	rego.INFO(w)
}
