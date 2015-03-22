package ygo

type Action func(pl *Player)

type Event struct {
	Owner *Card
	Event Action
	Name  string
}

type Events struct {
	Events map[string]Event
	Type   string
}

func NewEvent(name string) *Events {
	r := Events{
		Events: make(map[string]Event),
		Type:   name,
	}
	return &r
}

func (ev *Events) Register(name string, fun Event) {
	ev.Events[name] = fun
}
func (ev *Events) Unregister(name string) {
	delete(ev.Events, name)
}
