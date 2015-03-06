package defaul

import (
	"errors"
	"rego"
	"rego/agent"
	"rego/cfg"
	"rego/route"
)

var mapFile = "map.json"

func DefaulAgent() *agent.Agent {
	ro := route.NewRoute(cfg.Whole.Apps)
	ro.Code().WriteFile(mapFile)

	ag := agent.NewAgent(1024, func(user *agent.User, msg []byte) (err error) {
		defer rego.PanicErr(&err)
		var reply agent.Response
		//rego.INFO(string(msg))
		user.Refresh()
		err = ro.CallCode(msg[1], msg[2], msg[3], agent.Request{
			Session: &user.Session,
			Request: rego.NewEncodeBytes(msg[4:]),
		}, &reply)

		var ret []byte
		if err != nil {
			ret = []byte(`{"error":"` + err.Error() + `"}`)
		} else {
			if reply.Coverage != nil {
				user.Data = reply.Coverage
			} else if reply.Data != nil {
				var t interface{}
				user.Data.DeJson(&t)
				reply.Data.DeJson(&t)
				user.Data.EnJson(&t)
			}
			if reply.Error != "" {
				ret = []byte(`{"error":"` + reply.Error + `"}`)
			} else if reply.Response != nil {
				ret = reply.Response.Bytes()
			} else {
				return nil
			}
		}

		user.WriteMsg(append(msg[:4], ret...))
		return nil
	})
	return ag
}

func DefaulConn() agent.Conn {
	ca := cfg.Whole.Agents[0].ClientAddr()
	return agent.NewConn(ca)
}

func DefaulClient(hand func(byte, byte, byte, []byte) error) func(byte, byte, byte, []byte) error {
	conn := DefaulConn()
	size := 0
	isEnd := false
	errmsg := errors.New("use of closed network connection")
	go func() {
		for {
			b, err := conn.ReadMsg()
			if err != nil {
				break
			}
			if len(b) == 0 {
				continue
			}
			err = hand(b[1], b[2], b[3], b[4:])
			if err != nil {
				break
			}
		}
		isEnd = true
		conn.Close()
	}()
	return func(m1, m2, m3 byte, b []byte) error {
		if isEnd {
			return errmsg
		}
		err := conn.WriteMsg(append([]byte{byte(size), m1, m2, m3}, b...))
		if err != nil {
			isEnd = true
			return errmsg
		}
		return nil
	}
}

func DefaultClientCode(hand func(code string, v interface{}) error) func(code string, v interface{}) error {
	cod := route.NewCodeMaps()
	cod.ReadFile(mapFile)
	recod := cod.MakeReCodeMap()
	rego.INFO(recod)
	c := DefaulClient(func(m1 byte, m2 byte, m3 byte, b []byte) error {
		c1, c2, c3, err := cod.Map(m1, m2, m3)
		if err != nil {
			return err
		}
		es := rego.NewEncodeBytes(b)
		var r interface{}
		es.DeJson(&r)
		return hand(c1+"."+c2+"."+c3, r)
	})
	return func(code string, v interface{}) error {
		m1, m2, m3, err := recod.Map(code)
		if err != nil {
			return err
		}
		es := rego.EnJson(v)
		return c(m1, m2, m3, es.Bytes())
	}
}
