package configs

import (
	"fmt"
	"gooo/helper"
	"log"
	"os"
	"os/exec"
	"time"
)

type Configs struct {
	Sc ServersConfig
	Rc RouteConfig
	Mc MasterConfig
	Dc DataBaseConfig
}

func NewConfigsFrom(dir string) *Configs {
	mm := map[string][]byte{}
	ss := []string{"master", "servers", "route", "database"}
	for _, v := range ss {
		file := helper.OpenFile(fmt.Sprintf("%s/%s.json", dir, v))
		mm[v] = helper.ReplaceJson(file)
	}

	return NewConfigs(&mm)
}

func NewConfigs(m *map[string][]byte) *Configs {
	var sc ServersConfig
	helper.GetConfig((*m)["servers"], &sc)

	var rc RouteConfig
	helper.GetConfig((*m)["route"], &rc)

	var mc MasterConfig
	helper.GetConfig((*m)["master"], &mc)

	var dc DataBaseConfig
	helper.GetConfig((*m)["database"], &dc)

	return &Configs{
		Sc: sc,
		Rc: rc,
		Mc: mc,
		Dc: dc,
	}
}

func (c *Configs) StartServers() {
	csd := c.Sc
	for k1, v1 := range csd {
		for k2, v2 := range v1 {
			b := fmt.Sprintf("./%s", helper.ToLower(k1))
			args := []string{}
			if v2.Port != 0 {
				args = append(args, "-p", fmt.Sprintf("%d", v2.Port))
			}
			if v2.Bin != "" {
				b = fmt.Sprintf("./%s", v2.Bin)
			}
			args = append(args, "-i", fmt.Sprintf("%d", k2))
			args = append(args, "-t", fmt.Sprintf("%s", k1))
			args = append(args, "&")

			cmd := exec.Command(b, args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			go cmd.Run()
		}
	}
	time.Sleep(time.Second)
}

func (c *Configs) Foreach(class, method string, m, b interface{}) {
	for k1, v1 := range c.Sc {
		for k2, v2 := range v1 {
			if v2.Conn == nil {
				log.Printf("%s_%d Not connect\n", k1, k2)
				continue
			}
			var err error
			var t string
			if class == "" {
				t = fmt.Sprintf("%v.%s", k1, method)
				err = v2.Conn().Call(t, m, &b)

			} else {
				t = fmt.Sprintf("%s.%s", class, method)
				err = v2.Conn().Call(fmt.Sprintf("%s.%s", class, method), m, &b)
			}
			if err != nil {
				log.Printf("Can't Call %s from %s_%d %s:%d", t, k1, k2, v2.Host, v2.Port)
			}
		}
	}
}
