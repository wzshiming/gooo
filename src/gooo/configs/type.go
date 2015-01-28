package configs

import (
	"fmt"
	"net/rpc"
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

func (s *RouteConfig) Info(c1, c2, c3 uint8) string {
	b1 := (*s)[c1]
	b2 := b1.Map[c2]
	b3 := b2.Map[c3]
	return fmt.Sprintf("%s(%d).%s(%d).%s(%d) <0x%08X>",
		b1.Name, c1,
		b2.Name, c2,
		b3, c3,
		b1.Auth|b2.Auth)
}

type ServerConfig struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Name       string `json:"name"`
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
	NoRoom  uint32 `json:"noroom"`
	InRoom  uint32 `json:"inroom"`
	NoGame  uint32 `json:"nogame"`
	InGame  uint32 `json:"ingame"`
	Vip     uint32 `json:"vip"`
	Vip1    uint32 `json:"vip1"`
	Vip2    uint32 `json:"vip2"`
	Vip3    uint32 `json:"vip3"`
	Admin   uint32 `json:"admin"`
	Admin1  uint32 `json:"admin1"`
	Admin2  uint32 `json:"admin2"`
	Admin3  uint32 `json:"admin3"`
}
