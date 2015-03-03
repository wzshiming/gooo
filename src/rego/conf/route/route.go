package route

import (
	"rego/conf"
)

type Route struct {
	conf *conf.AppsConfig
}

func NewRoute(conf *conf.AppsConfig) *Route {
	return &Route{
		conf: conf,
	}
}
