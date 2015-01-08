package session

func NewUniqUint() func() uint {
	num := make(chan uint, 8)
	go func() {
		for i := uint(0); ; i++ {
			num <- i
		}
	}()
	return func() uint {
		return <-num
	}
}

type Unique struct {
	Server uint
	Id     uint
}

func NewUnique(server, id uint) Unique {
	return Unique{
		Server: server,
		Id:     id,
	}
}

func (s *Unique) IsFrom(server uint) bool {
	return s.Server == server
}
