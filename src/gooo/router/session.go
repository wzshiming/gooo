package router

import (
	"gooo/helper"
	"gooo/session"
)

func (s *CallServer) CallBySession(sess *session.Session, method string, args, reply interface{}) (err error) {
	defer helper.Recover()
	return s.Client[sess.Uniq.Server].Call(method, args, reply)
}

func (s *CallServer) CallBySessions(sesss []*session.Session, method string, args, reply interface{}) (err error) {
	for _, sess := range sesss {
		s.CallBySession(sess, method, args, reply)
	}
	return
}
