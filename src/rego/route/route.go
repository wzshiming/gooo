package route

import (
	"rego"
	"rego/cfg"
	"rego/server"
)

type Route struct {
	maps     CodeMaps
	connmaps map[string][]*server.Client
	conns    []*server.Client
}

func NewRoute(conf []cfg.ServerConfig) *Route {
	ro := &Route{
		connmaps: make(map[string][]*server.Client),
	}
	for _, v1 := range conf {
		ro.Register(v1)
		rego.INFO("hello")
	}
	return ro
}

func (ro *Route) Register(c cfg.ServerConfig) {
	conn := c.Client()
	if conn == nil {
		return
	}

	ro.conns = append(ro.conns, conn)
	if len(ro.connmaps[c.Type]) == 0 && c.Type != "" {
		classs, err := conn.TakeClasss()
		if err != nil {
			return
		}
		ro.maps = append(ro.maps, CodeMap{
			Name:   c.Type,
			Classs: classs,
		})
	}
	ro.connmaps[c.Type] = append(ro.connmaps[c.Type], conn)
}

func (ro *Route) CallCode(c1, c2, c3 byte, args rego.Request, reply *rego.Response) error {
	m1, m2, m3, err := ro.maps.Map(c1, c2, c3)
	if err != nil {
		return err
	}
	return ro.Call(m1, m2, m3, args, reply)
}

func (ro *Route) Call(m1, m2, m3 string, args rego.Request, reply *rego.Response) error {
	defer func() {
		if err := recover(); err != nil {
			rego.PANIC(err)
			return
		}
	}()
	return ro.connmaps[m1][0].Call(m2+"."+m3, args, reply)
}
