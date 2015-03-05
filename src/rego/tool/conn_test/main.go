package main

import (
	"rego"
	"rego/cfg"
	"rego/defaul"
)

func init() {
	cfg.Whole = cfg.NewWholeConfig(cfg.DirConf + "server.json")
}

func main() {
	send := defaul.DefaultClientCode(func(code string, v interface{}) {
		rego.INFO(code, v)
	})
	send("Auth.Auth.Register", map[string]interface{}{
		"hello": 10,
	})
	send("Auth.Auth.Register", map[string]interface{}{
		"hello": 11,
	})
	rego.Wait(nil)
}
