package protocol

import (
	"gooo/used"
)

type GameRooms struct {
	use  *used.Used
	list []GameRoom
}

func NewGameRooms(size int) *GameRooms {
	s := GameRooms{
		list: make([]GameRoom, size),
		use:  used.NewUsed(size),
	}
	return &s
}

func (s *GameRooms) List() RoomsResponse {
	l := s.use.List()
	r := make(RoomsResponse, len(l))
	for k, v := range l {
		r[k].Users = s.list[v].List()
		r[k].RoomId = s.list[v].roomId
		r[k].Name = s.list[v].name
		r[k].Started = s.list[v].started
	}
	return r
}

func (s *GameRooms) Join(roomid int, userid uint64) int {
	if s.use.IsUse(roomid) {
		return s.list[roomid].Join(userid)
	}
	return -1
}

func (s *GameRooms) Leave(roomid int, seat int) int {
	if s.use.IsUse(roomid) {
		return s.list[roomid].Leave(seat)
	}
	return -1
}

func (s *GameRooms) Create(size int, userid uint64) (int, int) {
	if i := s.use.Join(); i >= 0 {
		s.list[i] = *NewGameRoom(i, size, userid)
		seat := s.list[i].Join(userid)
		return i, seat
	}
	return -1, 0
}
