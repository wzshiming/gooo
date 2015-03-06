package agent

import (
	"errors"
)

type PushRequest struct {
	Uniq uint
	Msg  []byte
}

type Push struct {
	agent *Agent
}

func NewPush(ag *Agent) *Push {
	r := Push{
		agent: ag,
	}
	return &r
}
func (r *Push) Send(args PushRequest, reply *int) (err error) {
	if conn := r.agent.Get(args.Uniq); conn != nil {
		conn.WriteMsg(args.Msg)
		return nil
	}
	return errors.New("use of closed network connection")
}
