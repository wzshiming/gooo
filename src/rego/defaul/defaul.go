package defaul

import (
	"rego"
	"rego/agent"
	"rego/cfg"
	"rego/route"
)

var mapFile = "map.json"

func DefaulAgent() *agent.Agent {
	ro := route.NewRoute(cfg.Whole.Apps)
	ro.Code().WriteFile(mapFile)
	ag := agent.NewAgent(1024, func(conn agent.Conn, sess *rego.Session, msg []byte) {
		var reply rego.Response
		b := rego.EncodeStream(msg[4:])
		//rego.INFO(string(msg))
		err := ro.CallCode(msg[1], msg[2], msg[3], rego.Request{
			Session: sess,
			Request: &b,
		}, &reply)
		if err != nil {
			rego.ERR(err)
		}
		var ret []byte
		if reply.Error != "" {
			ret = []byte(`{"error":"` + reply.Error + `"}`)
		} else if reply.Response != nil {
			ret = reply.Response.Bytes()
		}
		//rego.INFO(string(ret))
		conn.WriteMsg(append(msg[:4], ret...))
	})
	return ag
}

func DefaulConn() agent.Conn {
	ca := cfg.Whole.Agents[0].ClientAddr()
	return agent.NewConn(ca)
}

func DefaulClient(hand func(byte, byte, byte, []byte)) func(byte, byte, byte, []byte) {
	conn := DefaulConn()
	size := 0
	isEnd := false
	go func() {
		for {
			b, err := conn.ReadMsg()
			if err != nil {
				isEnd = true
				return
			}
			hand(b[1], b[2], b[3], b[4:])
		}
	}()
	return func(m1, m2, m3 byte, b []byte) {
		if isEnd {
			return
		}
		err := conn.WriteMsg(append([]byte{byte(size), m1, m2, m3}, b...))
		if err != nil {
			isEnd = true
		}
	}
}

func DefaultClientCode(hand func(code string, v interface{})) func(code string, v interface{}) {
	cod := route.NewCodeMaps()
	cod.ReadFile(mapFile)
	recod := cod.MakeReCodeMap()
	rego.INFO(recod)
	c := DefaulClient(func(m1 byte, m2 byte, m3 byte, b []byte) {
		c1, c2, c3, err := cod.Map(m1, m2, m3)
		if err != nil {
			rego.ERR(err)
			return
		}
		es := rego.NewEncodeStream()
		es.Set(b)
		var r interface{}
		es.DeJson(&r)
		hand(c1+"."+c2+"."+c3, r)
	})
	return func(code string, v interface{}) {
		m1, m2, m3, err := recod.Map(code)
		if err != nil {
			rego.ERR(err)
			return
		}
		es := rego.NewEncodeStream()
		es.EnJson(v)
		c(m1, m2, m3, es.Bytes())
	}
}
