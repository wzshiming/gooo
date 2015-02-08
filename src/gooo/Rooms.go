package gooo

import ()

type GameRooms struct {
	use  *Used
	list []*GameRoom `json:"l"`
}

func NewGameRooms(size int) *GameRooms {
	s := GameRooms{
		list: make([]*GameRoom, size),
		use:  NewUsed(size),
	}
	return &s
}

func (s *GameRooms) List() []GameRoom {
	l := s.use.List()
	r := make([]GameRoom, len(l))
	for k, v := range l {
		if d := s.list[v]; d != nil {
			r[k] = *d
		}
	}
	return r
}

func (s *GameRooms) Room(roomid int) *GameRoom {
	if s.use.IsUse(roomid) {
		return s.list[roomid]
	}
	return nil
}

func (s *GameRooms) Join(roomid int, userid uint64) int {
	if s.use.IsUse(roomid) && !s.list[roomid].IsStarted() {
		return s.list[roomid].Join(userid)
	}
	return -1
}

func (s *GameRooms) Leave(roomid int, seat int) int {
	if s.use.IsUse(roomid) {
		s.list[roomid].Leave(seat)
		if s.list[roomid].Size() == 0 {
			s.use.Leave(roomid)
			s.list[roomid] = nil
		}
	}
	return -1
}

func (s *GameRooms) Create(size int, userid uint64) (int, int) {
	if i := s.use.Join(); i >= 0 {
		s.list[i] = NewGameRoom(i, size, userid)
		seat := s.list[i].Join(userid)
		return i, seat
	}
	return -1, -1
}

func (s *GameRooms) Play(roomid int) {
	if s.use.IsUse(roomid) {
		s.list[roomid].Play()
	}
}

func (s *GameRooms) End(roomid int) {
	if s.use.IsUse(roomid) {
		s.list[roomid].End()
	}
}
