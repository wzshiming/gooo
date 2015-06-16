package main

import (
	"github.com/wzshiming/rego"
	"github.com/wzshiming/rego/agent/defaul"
	"github.com/wzshiming/rego/cfg"
	"service/proto"
)

func init() {
	cfg.Whole = cfg.NewWholeConfig(cfg.DirConf + "server.json")
}

func main() {
	send := defaul.DefaultClientCode(func(code string, v interface{}) error {
		rego.INFO(code, v)
		return nil
	})
	//	send("Auth.Auth.Loop", 1)
	send("Auth.Users.Register", proto.LogInRequest{
		Username: "hello2",
		Password: "123456",
	})

	send("Auth.Users.LogIn", proto.LogInRequest{
		Username: "hello2",
		Password: "123456",
	})

	//send("Auth.Users.LogOut", 1)

	//	for i := 0; i != 10; i++ {
	//		//for {
	//		err := send("Auth.Auth.Register", map[string]interface{}{
	//			"hello": 10,
	//		})
	//		if err != nil {
	//			break
	//		}

	//		//}
	//	}

	rego.NOTICE("close")
	rego.Wait(nil)
}
