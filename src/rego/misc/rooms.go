package misc

import (
	"errors"
	"rego"
	"rego/agent"
)

type Rooms struct {
	name string
	list map[uint]*agent.Session
}
type datafmt struct {
	Rooms map[string]uint `json:"__Rooms__"`
}

func GetFromRoom(sess *agent.Session, name string) uint {
	var r datafmt
	sess.Data.DeJson(&r)
	return r.Rooms[name]
}

func NewRooms(name string) *Rooms {
	return &Rooms{
		name: name,
		list: make(map[uint]*agent.Session),
	}
}

func (ro *Rooms) JoinFromMutex(uniq uint, sess *agent.Session) error {
	sess.LockSession()
	ret, err := ro.JoinFrom(uniq, sess)
	sess.UnlockSession(&agent.Response{
		Data: ret,
	})
	return err
}

func (ro *Rooms) JoinMutex(sess *agent.Session) error {
	sess.LockSession()
	ret, err := ro.Join(sess)
	sess.UnlockSession(&agent.Response{
		Data: ret,
	})
	return err
}

func (ro *Rooms) JoinFrom(uniq uint, sess *agent.Session) (*rego.EncodeBytes, error) {
	var r datafmt
	sess.Data.DeJson(&r)
	if r.Rooms == nil {
		r.Rooms = make(map[string]uint, 1)
	}
	if se := r.Rooms[ro.name]; se != 0 {
		return nil, errors.New("Rooms.Join: already join")
	}
	if un := ro.list[uniq]; un != nil {
		return nil, errors.New("Rooms.Join: repetition join")
	}
	r.Rooms[ro.name] = uniq
	ro.list[uniq] = sess
	return rego.EnJson(r), nil
}

func (ro *Rooms) Join(sess *agent.Session) (*rego.EncodeBytes, error) {
	return ro.JoinFrom(sess.ToUint(), sess)
}

func (ro *Rooms) LeaveMutex(sess *agent.Session) {
	sess.LockSession()
	sess.UnlockSession(&agent.Response{
		Data: ro.Leave(sess),
	})
}

func (ro *Rooms) Leave(sess *agent.Session) *rego.EncodeBytes {
	uniq := ro.Uniq(sess)
	var r datafmt
	sess.Data.DeJson(&r)
	delete(r.Rooms, ro.name)
	delete(ro.list, uniq)
	sess.Data = rego.EnJson(r)
	return rego.EnJson(map[string]interface{}{ro.name: nil})
}

func (ro *Rooms) SyncMutex(sess *agent.Session) *agent.Session {
	se := ro.Sync(sess)
	if se != nil {
		se.SyncSession()
	}
	return se
}

func (ro *Rooms) Uniq(sess *agent.Session) uint {
	var r datafmt
	sess.Data.DeJson(&r)
	return r.Rooms[ro.name]
}

func (ro *Rooms) Sync(sess *agent.Session) *agent.Session {
	if se := ro.Uniq(sess); se != 0 {
		return ro.Get(se)
	}
	return nil
}

func (ro *Rooms) Get(uniq uint) *agent.Session {
	return ro.list[uniq]
}

func (ro *Rooms) Len() int {
	return len(ro.list)
}

func (ro *Rooms) ForEach(fun func(*agent.Session)) {
	for _, v := range ro.list {
		fun(v)
	}
}

func (ro *Rooms) Broadcast(reply *agent.Response, fail func(*agent.Session)) {
	var err error
	ro.ForEach(func(sess *agent.Session) {
		err = sess.Send(reply)
		if err != nil {
			rego.ERR(err)
			ro.Leave(sess)
			if fail != nil {
				fail(sess)
			}
		}
	})
}
