package uniq

func NewUniq() func() uint {
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
