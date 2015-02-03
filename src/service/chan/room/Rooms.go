package room

import (
	"gooo/used"

	chanprtc "service/chan/protocol"
)

type GameRooms struct {
	use  *used.Used
	list []*GameRoom
}

func NewGameRooms(size int) *GameRooms {
	s := GameRooms{
		list: make([]*GameRoom, size),
		use:  used.NewUsed(size),
	}
	return &s
}

func (s *GameRooms) List() chanprtc.RoomsResponse {
	l := s.use.List()
	r := make(chanprtc.RoomsResponse, len(l))
	for k, v := range l {
		if d := s.list[v]; d != nil {
			r[k].Users = d.List()
			r[k].RoomId = d.roomId
			r[k].Name = d.name
			r[k].Started = d.started
		}
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
	return -1, 0
}
