package route

import (
	"gooo/protocol"
)

type ConnRpc struct {
	hand *Handel
}

func NewConnRpc(m *Handel) *ConnRpc {
	s := ConnRpc{
		hand: m,
	}
	return &s
}

func (s *ConnRpc) Send(args protocol.SendRequest, reply *protocol.SendResponse) error {
	for _, v := range args.Clients {
		c := s.hand.Session.Get(v)
		if c != nil {
			c.Conn().Write(append([]byte{255, 255, 255, 255}, args.Data...))
		}
	}
	return nil
}

func (s *ConnRpc) SetSession(args protocol.SetSessionRequest, reply *protocol.SetSessionResponse) error {
	obj := s.hand.Session.Get(args.Id)
	obj.Data = args.Data
	obj.Unlock()
	return nil
}

func (s *ConnRpc) GetSession(args protocol.GetSessionRequest, reply *protocol.GetSessionResponse) error {
	obj := s.hand.Session.Get(args.Id)
	obj.Lock()
	*reply = protocol.GetSessionResponse{
		Data: obj.Data,
	}
	return nil
}
