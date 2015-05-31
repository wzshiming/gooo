package ygo

type Action func(ca *Card) bool

func (ac Action) Call(ca *Card) (ok bool) {
	if ac != nil {
		return ac(ca)
	}
	return false
}

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
