package conf

import (
	"fmt"
	"rego"
	"rego/server"
)

type ServerConfig struct {
	Name       string `json:"name"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	ClientPort int    `json:"clientport"`
	Bin        string `json:"bin"`
}

func (se *ServerConfig) Addr() string {
	return fmt.Sprintf("%s:%d", se.Host, se.Port)
}

func (se *ServerConfig) ExposedAddr() string {
	if se.ClientPort != 0 {
		return fmt.Sprintf("%s:%d", rego.LocalIP, se.Port)
	}
	return ""
}

func (se *ServerConfig) Client() *server.Client {
	return server.NewClient(se.Addr())
}

func (se *ServerConfig) Server() *server.Server {
	return server.NewServer(se.Port)
}
