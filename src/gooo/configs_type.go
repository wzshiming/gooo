package gooo

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"os/exec"
	"time"
)

type names struct {
	Name    string `json:"name"`
	Allow   uint32 `json:"allow"`
	Unallow uint32 `json:"unallow"`
}

type RoutesConfig []struct {
	names
	Map []struct {
		names
		Map []string `json:"map"`
	} `json:"map"`
}

func (s *RoutesConfig) FindIndex(c1, c2, c3 string) (i1, i2, i3 uint8) {
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

func (s *ServerConfig) Start(k1 string, k2 int) {
	b := fmt.Sprintf("./%s", ToLower(k1))
	args := []string{}
	if s.Bin != "" {
		b = fmt.Sprintf("./%s", s.Bin)
	}
	if s.Port != 0 {
		args = append(args, "-p", fmt.Sprintf("%d", s.Port))
	}
	args = append(args, "-t", fmt.Sprintf("%s", k1))
	args = append(args, "-i", fmt.Sprintf("%d", k2))
	args = append(args, "&")
	cmd := exec.Command(b, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	go cmd.Run()
}

func (s *ServerConfig) Call(class, method string, m, b interface{}) {
	t := fmt.Sprintf("%s.%s", class, method)
	err := s.Conn().Call(fmt.Sprintf("%s.%s", class, method), m, &b)
	if err != nil {
		log.Printf("Can't Call %s from %s:%d", t, s.Host, s.Port)
	}
}

type ServersConfig map[string][]ServerConfig

func (c *ServersConfig) Size(name string) int {
	defer Recover()
	return len((*c)[name])
}

func (c *ServersConfig) Self() *ServerConfig {
	defer Recover()
	return &((*c)[Type][Id])
}

func (c *ServersConfig) Start() {
	defer Recover()
	for k1, v1 := range *c {
		for k2, v2 := range v1 {
			v2.Start(k1, k2)
		}
	}
	time.Sleep(time.Second)
}

func (c *ServersConfig) Foreach(class, method string, m, b interface{}) {
	defer Recover()
	for _, v1 := range *c {
		for _, v2 := range v1 {
			v2.Call(class, method, m, b)
		}
	}
}

type MasterConfig struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

type DataBaseConfig struct {
	Dialect string `json:"dialect"`
	Source  string `json:"source"`
}

type DataBasesConfig map[string]DataBaseConfig

const (
	FlagLogin   uint32 = (1 << iota)
	FlagRoom    uint32 = (1 << iota)
	FlagGame    uint32 = (1 << iota)
	FlagChan    uint32 = (1 << iota)
	FlagUnallow uint32 = (1 << iota)
	FlagWarning uint32 = (1 << iota)
	FlagVip     uint32 = (1 << iota)
	FlagAdmin   uint32 = (1 << 31)
	FlagNone    uint32 = 0
)
