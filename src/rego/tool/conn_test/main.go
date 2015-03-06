package main

import (
	"rego"
	"rego/agent/defaul"
	"rego/cfg"
)

func init() {
	cfg.Whole = cfg.NewWholeConfig(cfg.DirConf + "server.json")
}

func main() {
	send := defaul.DefaultClientCode(func(code string, v interface{}) error {
		rego.INFO(code, v)
		return nil
	})
	for {
		err := send("Auth.Auth.Register", map[string]interface{}{
			"hello": 10,
		})
		if err != nil {
			break
		}

	}
	rego.NOTICE("close")
	rego.Wait(nil)
}
