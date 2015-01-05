package balance


type Balance struct {
    size    int
    i       int
}

func NewBalance(size int) *Balance{
    return &Balance{
        size:   size,
        i:      0,
    }
}

func (b *Balance)Allot()int {
    if b.i == b.size {
        b.i = 0
        return 0
    }
    i := b.i
    b.i++
    return i
}
