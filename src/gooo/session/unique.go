package session

func NewUniqUint() func() uint64 {
	num := make(chan uint64, 8)
	go func() {
		for i := uint64(0); ; i++ {
			num <- i
		}
	}()
	return func() uint64 {
		return <-num
	}
}

type Unique struct {
	Server uint64
	Id     uint64
}

func NewUnique(server, id uint64) Unique {
	return Unique{
		Server: server,
		Id:     id,
	}
}

func (s *Unique) IsFrom(server uint64) bool {
	return s.Server == server
}
