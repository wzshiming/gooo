package room

import (
	"gooo/used"
)

type GameRoom struct {
	use     *used.Used
	users   []uint64
	roomId  int
	name    string
	started bool
}

func NewGameRoom(roomid int, size int, userid uint64) *GameRoom {
	s := GameRoom{
		users:  make([]uint64, size),
		use:    used.NewUsed(size),
		roomId: roomid,
	}
	return &s
}

func (s *GameRoom) Size() int {
	return s.use.Size()
}

func (s *GameRoom) List() []uint64 {
	return s.users
}

func (s *GameRoom) Join(userid uint64) int {
	if i := s.use.Join(); i >= 0 {
		s.users[i] = userid
		return i
	}
	return -1
}

func (s *GameRoom) Leave(seat int) int {
	if i := s.use.Leave(seat); i == 0 {
		s.users[seat] = 0
		return 0
	}
	return -1
}
