package configs

import (
	"net/rpc"
	//"log"
)

type RouteConfig struct {
	Devel []struct {
		Name  string `json: name`
		Class []struct {
			Name    string   `json: name`
			Methods []string `json: methods`
		} `json: class`
		Conn *rpc.Client
	} `json: devel`
}

type ServersConfig struct {
	Devel map[string][]struct {
		Name       string `json: name`
		Host       string `json: host`
		Port       int    `json: port`
		ClientPort int    `json: clientport`
		Control    bool   `json: control`
		Conn       *rpc.Client
	} `json: devel`
}

type MasterConfig struct {
	Devel struct {
		Name string `json: name`
		Host string `json: host`
		Port int    `json: port`
	} `json: devel`
}
