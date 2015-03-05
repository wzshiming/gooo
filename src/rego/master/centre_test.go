package main

import (
	"rego/cfg"
	"testing"
)

func Test_get(t *testing.T) {
	go start()
	cfg.TakeConf()
}
