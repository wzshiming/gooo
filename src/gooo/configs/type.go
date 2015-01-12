package configs

import (
	"net/rpc"
	//"log"
)

//type RouteConfig map[string]map[string][]string

type RouteConfig []struct {
	Name string `json: name`
	Map  []struct {
		Name string   `json: name`
		Map  []string `json: map`
	} `json: map`
}

type ServerConfig struct {
	Name       string `json: name`
	Host       string `json: host`
	Port       int    `json: port`
	ClientPort int    `json: clientport`
	Control    bool   `json: control`
	Conn       *rpc.Client
}

type ServersConfig map[string][]ServerConfig

type MasterConfig struct {
	Name string `json: name`
	Host string `json: host`
	Port int    `json: port`
}
