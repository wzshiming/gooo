package room

import (
	"gooo/used"
)

type GameRoom struct {
	use     *used.Used
	Users   []uint64 `json:"u"`
	RoomId  int      `json:"r"`
	Master  int      `json:"m"`
	Name    string   `json:"n"`
	Started bool     `json:"s"`
}

func NewGameRoom(roomid int, size int, userid uint64) *GameRoom {
	s := GameRoom{
		Users:  make([]uint64, size),
		use:    used.NewUsed(size),
		RoomId: roomid,
		Master: 0,
	}
	return &s
}

func (s *GameRoom) IsStarted() bool {
	return s.Started
}

func (s *GameRoom) Play() {
	s.Started = true
}

func (s *GameRoom) End() {
	s.Started = false
}

func (s *GameRoom) Size() int {
	return s.use.Size()
}

func (s *GameRoom) List() []uint64 {
	return s.Users
}

func (s *GameRoom) Join(userid uint64) int {
	if i := s.use.Join(); i >= 0 {
		s.Users[i] = userid
		return i
	}
	return -1
}

func (s *GameRoom) Leave(seat int) int {
	if i := s.use.Leave(seat); i == 0 {
		s.Users[seat] = 0
		if seat == s.Master && s.Size() != 0 {
			s.Master = s.use.Unusable()
		}
		return 0
	}
	return -1
}
