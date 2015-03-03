package conf

import (
	"io/ioutil"
	"rego"
)

type WholeConfig struct {
	Agents AgentsConfig `json:"agents"`
	Apps   AppsConfig   `json:"apps"`
	Dbs    DbsConfig    `json:"dbs"`
}
type DbsConfig map[string]DbConfig
type DbConfig struct {
	Dialect string `json:"dialect"`
	Source  string `json:"source"`
}

type AppsConfig map[string][]ServerConfig

type AgentsConfig []ServerConfig

func NewServerConfig(path string) *ServerConfig {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		rego.ERR(err)
		return nil
	}
	en := rego.NewEncodeStream()
	var conf ServerConfig
	en.Set(b)
	en.DnJson(&conf)
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
	en.DnJson(&conf)
	return &conf
}
