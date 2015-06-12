package ygo

type Action func(ca *Card) bool

func (ac Action) Call(ca *Card) bool {
	if ac != nil {
		return ac(ca)
	}
	return false
}

type Effect struct {
	Owner *Card
	Place *Effects
	Event Action
}

func (ac *Effect) Remove() {
	ac.Place.Unregister(ac.Owner)
}

func (ac *Effect) Call(ca *Card) bool {
	return ac.Event.Call(ca)
}

type Effects struct {
	Events map[*Card]Effect
	Name   string
}

func NewEffects(name string) *Effects {
	r := Effects{
		Events: make(map[*Card]Effect),
		Name:   name,
	}
	return &r
}

func (ev *Effects) Register(card *Card, fun Action) {
	ev.Events[card] = Effect{
		Owner: card,
		Place: ev,
		Event: fun,
	}
}
func (ev *Effects) Unregister(card *Card) {
	delete(ev.Events, card)
}
