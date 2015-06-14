package ygo

type Action func(ca *Card) bool

func (ac Action) IsExits() bool {
	if ac != nil {
		return true
	}
	return false
}

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
	events map[*Card]Effect
	name   string
}

func NewEffects(name string) *Effects {
	r := Effects{
		events: make(map[*Card]Effect),
		name:   name,
	}
	return &r
}

func (ev *Effects) GetNmae() string {
	return ev.name
}

func (ev *Effects) Call() {
	for k, v := range ev.events {
		v.Call(k)
	}
}

func (ev *Effects) Register(card *Card, fun Action) {
	ev.events[card] = Effect{
		Owner: card,
		Place: ev,
		Event: fun,
	}
}
func (ev *Effects) Unregister(card *Card) {
	delete(ev.events, card)
}
