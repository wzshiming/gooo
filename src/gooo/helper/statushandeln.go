package helper

import ()

type Status struct {
	hand *Handeln
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
