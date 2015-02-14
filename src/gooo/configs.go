package gooo

import (
	"fmt"
)

type Configs struct {
	Sc ServersConfig
	Mc MasterConfig
	Dc DataBasesConfig
}

func NewConfigsFrom(dir string) *Configs {
	mm := map[string][]byte{}
	ss := []string{"master", "servers", "database"}
	for _, v := range ss {
		file := ReadFile(fmt.Sprintf("%s/%s.json", dir, v))
		mm[v] = ReplaceJson(file)
	}

	return NewConfigs(&mm)
}

func NewConfigs(m *map[string][]byte) *Configs {
	var sc ServersConfig
	GetConfig((*m)["servers"], &sc)

	var mc MasterConfig
	GetConfig((*m)["master"], &mc)

	var dc DataBasesConfig
	GetConfig((*m)["database"], &dc)

	return &Configs{
		Sc: sc,

		Mc: mc,
		Dc: dc,
	}
}

func (c *Configs) StartServers() {
	c.Sc.Start()
}

func (c *Configs) Foreach(class, method string, m, b interface{}) {
	c.Sc.Foreach(class, method, m, b)
}

func (c *Configs) Self() *ServerConfig {
	return c.Sc.Self()
}

func (c *Configs) Servers() *ServersConfig {
	return &c.Sc
}

func (c *Configs) Master() *MasterConfig {
	return &c.Mc
}

func (c *Configs) DataBase(name string) DataBaseConfig {
	return c.Dc[name]
}

func (c *Configs) TypeSize(name string) int {
	return c.Sc.Size(name)
}
