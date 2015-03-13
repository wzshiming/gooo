package agent

import (
	"errors"
)

type Connect struct {
	agent *Agent
}

func NewConnect(ag *Agent) *Connect {
	r := Connect{
		agent: ag,
	}
	return &r
}

type PushRequest struct {
	Uniq  uint
	Reply *Response
}

func (r *Connect) Push(args PushRequest, reply *int) (err error) {
	if conn := r.agent.Get(args.Uniq); conn != nil {
		if args.Reply != nil {
			args.Reply.Hand(conn, []byte{255, 255, 255, 255})
		}
		return nil
	}
	return errors.New("Connect.Push: use of closed network connection")
}

func (r *Connect) Sync(args LockRequest, reply *LockResponse) (err error) {
	if conn := r.agent.Get(args.Uniq); conn != nil {
		reply.Session = &conn.Session
		return nil
	}
	return errors.New("Connect.Lock: use of closed network connection")
}

type LockRequest struct {
	Uniq uint
}

type LockResponse struct {
	Session *Session
}

func (r *Connect) Lock(args LockRequest, reply *LockResponse) (err error) {
	if conn := r.agent.Get(args.Uniq); conn != nil {
		conn.Lock()
		reply.Session = &conn.Session
		return nil
	}
	return errors.New("Connect.Lock: use of closed network connection")
}

type UnlockRequest struct {
	Reply *Response
	Uniq  uint
}

func (r *Connect) Unlock(args UnlockRequest, reply *int) (err error) {
	if conn := r.agent.Get(args.Uniq); conn != nil {
		if args.Reply != nil {
			args.Reply.Hand(conn, []byte{255, 255, 255, 255})
		}
		conn.Unlock()
		return nil
	}
	return errors.New("Connect.Unlock: use of closed network connection")
}
