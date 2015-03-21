package ygo

type Action func()

type Event struct {
	Owner *Card
	Event Action
	Name  string
}

type Events struct {
	Events map[string]Event
	Type   string
}

func NewEvent(max int, name string) *Events {
	r := Events{
		Events: make(map[string]Event, max),
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
