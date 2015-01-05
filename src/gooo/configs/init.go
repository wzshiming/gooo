package configs

import (
	"gooo/helper"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"os/exec"
)

type Configs struct {
	Sc   ServersConfig
	Rc   RouteConfig
	Mc   MasterConfig
	Type string
	Name string
	Port string
	Id   int
}

func NewConfigs(m *map[string][]byte) *Configs {
	var sc ServersConfig
	helper.GetConfig((*m)["servers"], &sc)

	var rc RouteConfig
	helper.GetConfig((*m)["route"], &rc)

	var mc MasterConfig
	helper.GetConfig((*m)["master"], &mc)
	return &Configs{
		Sc:   sc,
		Rc:   rc,
		Mc:   mc,
		Name: Name,
		Type: Type,
		Port: Port,
		Id:   Id,
	}
}

func (c *Configs) StartServers() {
	csd := c.Sc.Devel
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
	csd := c.Sc.Devel
	for k1, v1 := range csd {
		for k2, v2 := range v1 {
			Name := fmt.Sprintf("%v_%v", k1, k2)
			if Name != c.Name && v2.Control {
				t = append(t, v2.Conn)
			}

		}
	}
	return
}

func (c *Configs) StartConnect() {
	csd := c.Sc.Devel
	for k1, v1 := range csd {
		for k2, v2 := range v1 {
			name := fmt.Sprintf("%v_%v", k1, k2)
			addr := fmt.Sprintf("%v:%v", v2.Host, v2.Port)
			if name != c.Name && v2.Control {
				c.Sc.Devel[k1][k2].Conn = helper.NewConn(addr)
				log.Println(c.Name, "connect to", addr, name)
			}
		}
	}
	return
}
