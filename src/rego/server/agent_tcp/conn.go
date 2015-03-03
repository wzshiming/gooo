package main

import (
	"rego"
	"rego/agent"
)

func main() {
	ag := agent.NewAgent(16, func(conn agent.Conn, sess *rego.Session, msg []byte) {
		rego.NOTICE(string(msg))
		conn.WriteMsg(msg)
		conn.Close()
	})
	agent.NewListener(7710, func(conn agent.Conn) {
		ag.Join(conn)
	}).Start()
}
