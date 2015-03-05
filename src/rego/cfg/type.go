package cfg

import (
	"io/ioutil"
	"rego"
)

type WholeConfig struct {
	Agents []ServerConfig `json:"agents"`
	Apps   []ServerConfig `json:"apps"`
	Dbs    DbsConfig      `json:"dbs"`
}

func (wh *WholeConfig) Shutdown() {
	for _, v := range wh.Agents {
		v.Client().ShutdownNow()
	}
	for _, v := range wh.Apps {
		v.Client().ShutdownNow()
	}
}
func (wh *WholeConfig) Start() {
	for _, v := range wh.Apps {
		v.Start()
	}
	for _, v := range wh.Agents {
		v.Start()
	}
}

type DbsConfig map[string]DbConfig
type DbConfig struct {
	Dialect string `json:"dialect"`
	Source  string `json:"source"`
}

func NewServerConfig(path string) *ServerConfig {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		rego.ERR(err)
		return nil
	}
	en := rego.NewEncodeStream()
	var conf ServerConfig
	en.Set(b)
	en.DeJson(&conf)
	conf.makeId()
	return &conf
}

func NewWholeConfig(path string) *WholeConfig {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		rego.ERR(err)
		return nil
	}
	en := rego.NewEncodeStream()
	var conf WholeConfig
	en.Set(b)
	en.DeJson(&conf)
	for k, _ := range conf.Apps {
		conf.Apps[k].makeId()
	}
	for k, _ := range conf.Agents {
		conf.Agents[k].makeId()
	}
	return &conf
}
