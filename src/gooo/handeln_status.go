package gooo

import ()

type Status struct {
	Hand *Handeln
}

func NewStatus(hand *Handeln) *Status {
	s := Status{
		Hand: hand,
	}
	return &s
}

func (s *Status) Stop(args int, reply *int) error {
	if args == 222 {
		s.Hand.Stop()
	}
	return nil
}
