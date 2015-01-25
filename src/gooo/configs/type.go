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
