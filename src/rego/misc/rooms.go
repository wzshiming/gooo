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

type data struct {
	ID   uint `json:",string"`
	Head []byte
}

type datafmt struct {
	Rooms map[string]data `json:"__Rooms__"`
}

func GetFromRoom(sess *agent.Session, name string) uint {
	var r datafmt
	sess.Data.DeJson(&r)
	return r.Rooms[name].ID
}

func GetFromRoomHead(sess *agent.Session, name string) []byte {
	var r datafmt
	sess.Data.DeJson(&r)
	return r.Rooms[name].Head
}

func NewRooms(name string) *Rooms {
	return &Rooms{
		name: name,
		list: make(map[uint]*agent.Session),
	}
}

func (ro *Rooms) JoinFromMutex(uniq uint, sess *agent.Session, head []byte) error {
	sess.LockSession()
	ret, err := ro.JoinFrom(uniq, sess, head)
	sess.UnlockSession(&agent.Response{
		Data: ret,
	})
	return err
}

func (ro *Rooms) JoinMutex(sess *agent.Session, head []byte) error {
	sess.LockSession()
	ret, err := ro.Join(sess, head)
	sess.UnlockSession(&agent.Response{
		Data: ret,
	})
	return err
}

func (ro *Rooms) JoinFrom(uniq uint, sess *agent.Session, head []byte) (*rego.EncodeBytes, error) {
	var r datafmt
	sess.Data.DeJson(&r)
	if r.Rooms == nil {
		r.Rooms = make(map[string]data)
	}
	if se := r.Rooms[ro.name].ID; se != 0 {
		return nil, errors.New("Rooms.Join: already join")
	}
	if un := ro.list[uniq]; un != nil {
		return nil, errors.New("Rooms.Join: repetition join")
	}
	r.Rooms[ro.name] = data{
		ID:   uniq,
		Head: head,
	}
	ro.list[uniq] = sess
	ret := rego.EnJson(r)
	sess.Data = rego.SumJson(sess.Data, ret)
	return ret, nil
}

func (ro *Rooms) Join(sess *agent.Session, head []byte) (*rego.EncodeBytes, error) {
	return ro.JoinFrom(sess.ToUint(), sess, head)
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
	re := rego.EnJson(r)
	sess.Data = rego.SumJson(sess.Data, re)
	return re
}

func (ro *Rooms) SyncMutex(sess *agent.Session) *agent.Session {
	se := ro.Sync(sess)
	if se != nil {
		se.SyncSession()
	}
	return se
}

func (ro *Rooms) Uniq(sess *agent.Session) uint {
	return GetFromRoom(sess, ro.name)
}

func (ro *Rooms) Head(sess *agent.Session) []byte {
	return GetFromRoomHead(sess, ro.name)
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

func (ro *Rooms) Random() *agent.Session {
	for _, v := range ro.list {
		return v
	}
	return nil
}

func (ro *Rooms) Broadcast(reply *agent.Response, fail func(*agent.Session)) {
	var err error
	ro.ForEach(func(sess *agent.Session) {
		if reply.Head == nil {
			reply.Head = ro.Head(sess)
		}
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
