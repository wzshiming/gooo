package configs

import (
	"fmt"
	"gooo/helper"
	"log"
	"net/rpc"
	"os"
	"os/exec"
)

type Configs struct {
	Sc ServersConfig
	Rc RouteConfig
	Mc MasterConfig
}

func NewConfigs(m *map[string][]byte) *Configs {
	var sc ServersConfig
	helper.GetConfig((*m)["servers"], &sc)

	var rc RouteConfig
	helper.GetConfig((*m)["route"], &rc)

	var mc MasterConfig
	helper.GetConfig((*m)["master"], &mc)
	return &Configs{
		Sc: sc,
		Rc: rc,
		Mc: mc,
	}
}

func (c *Configs) StartServers() {
	csd := c.Sc
	for k1, v1 := range csd {
		b := fmt.Sprintf("./%s", helper.ToLower(k1))
		for k2, v2 := range v1 {
			args := []string{}
			if v2.Port != 0 {
				args = append(args, "-p", fmt.Sprintf("%d", v2.Port))
			}

			if v2.Control {
				args = append(args, "-i", fmt.Sprintf("%d", k2))
				args = append(args, "-t", fmt.Sprintf("%s", k1))
			}

			//Control
			//fmt.Println(args)
			cmd := exec.Command(b, args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			go cmd.Run()
		}
	}
}

func (c *Configs) AllConnect() (t []*rpc.Client) {
	csd := c.Sc
	for k1, v1 := range csd {
		for k2, v2 := range v1 {
			name := fmt.Sprintf("%v_%v", k1, k2)
			if name != Name && v2.Control {
				t = append(t, v2.Conn)
			}
		}
	}
	return
}

func (c *Configs) StartConnect() {
	csd := c.Sc
	for k1, v1 := range csd {
		for k2, v2 := range v1 {
			name := fmt.Sprintf("%v_%v", k1, k2)
			addr := fmt.Sprintf("%v:%v", v2.Host, v2.Port)
			if name != Name && v2.Control {
				c.Sc[k1][k2].Conn = helper.NewConn(addr)
				log.Println(Name, "connect to", addr, name)
			}
		}
	}
	return
}
