package gooo

import ()

type Status struct {
	hand *Handeln
	conf *Configs
}

func NewStatus(hand *Handeln) *Status {
	s := Status{
		hand: hand,
	}
	return &s
}

func (s *Status) Stop(args int, reply *int) error {
	if args == 222 {
		s.hand.Stop()
	}
	return nil
}

func (s *Status) Init(args InitRequest, reply *int) error {
	if args.State == 222 {
		s.conf = &args.Conf
	}
	return nil
}
