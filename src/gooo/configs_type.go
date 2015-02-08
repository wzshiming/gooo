package gooo

import (
	"fmt"
	"net/rpc"
)

type RouteConfig []struct {
	Name    string `json:"name"`
	Allow   uint32 `json:"allow"`
	Unallow uint32 `json:"unallow"`
	Map     []struct {
		Name    string   `json:"name"`
		Allow   uint32   `json:"allow"`
		Unallow uint32   `json:"unallow"`
		Map     []string `json:"map"`
	} `json:"map"`
}

func (s *RouteConfig) FindIndex(c1, c2, c3 string) (i1, i2, i3 uint8) {
	for k1, v1 := range *s {
		if v1.Name == c1 {
			for k2, v2 := range v1.Map {
				if v2.Name == c2 {
					for k3, v3 := range v2.Map {
						if v3 == c3 {
							return uint8(k1), uint8(k2), uint8(k3)
						}
					}
				}
			}
		}
	}
	return 255, 255, 255
}

func (s *RouteConfig) Info(c1, c2, c3 uint8) string {
	b1 := (*s)[c1]
	b2 := b1.Map[c2]
	b3 := b2.Map[c3]
	return fmt.Sprintf("%s(%d).%s(%d).%s(%d) <0x%08X^0x%08X>",
		b1.Name, c1,
		b2.Name, c2,
		b3, c3,
		b1.Allow|b2.Allow,
		b1.Unallow|b2.Unallow)
}

type ServerConfig struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Name       string `json:"name"`
	ClientPort int    `json:"clientport"`
	Bin        string `json:"bin"`
	conn       *rpc.Client
}

func (s *ServerConfig) Conn() *rpc.Client {
	if s.conn == nil {
		addr := fmt.Sprintf("%v:%v", s.Host, s.Port)
		//log.Println(Name, "connect to", addr)
		s.conn = NewConnTcp(addr)
	}
	return s.conn
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

const (
	FlagLogin   uint32 = (1 << iota)
	FlagRoom    uint32 = (1 << iota)
	FlagGame    uint32 = (1 << iota)
	FlagUnallow uint32 = (1 << iota)
	FlagWarning uint32 = (1 << iota)
	FlagVip     uint32 = (1 << iota)
	FlagAdmin   uint32 = (1 << 31)
	FlagNone    uint32 = 0
)
