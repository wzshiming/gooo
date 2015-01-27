package configs

import (
	"net/rpc"
	//"log"
)

//type RouteConfig map[string]map[string][]string

type RouteConfig []struct {
	Name string `json:"name"`
	Auth uint32 `json:"auth"`
	Map  []struct {
		Name string   `json:"name"`
		Auth uint32   `json:"auth"`
		Map  []string `json:"map"`
	} `json:"map"`
}

type ServerConfig struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	ClientPort int    `json:"clientport"`
	conn       *rpc.Client
}

type ServersConfig map[string][]ServerConfig

type MasterConfig struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

type DataBaseConfig map[string]struct {
	Dialect string `json:"dialect"`
	Source  string `json:"source"`
}

type StatusConfig struct {
	None    uint32 `json:"none"`
	NoLogin uint32 `json:"nologin"`
	Login   uint32 `json:"login"`
	Vip     uint32 `json:"vip"`
	Vip1    uint32 `json:"vip1"`
	Vip2    uint32 `json:"vip2"`
	Vip3    uint32 `json:"vip3"`
	Admin   uint32 `json:"admin"`
	Admin1  uint32 `json:"admin1"`
	Admin2  uint32 `json:"admin2"`
	Admin3  uint32 `json:"admin3"`
}
