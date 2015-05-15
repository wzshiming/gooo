package defaul

var DefaultDeck = []uint{}

func init() {
	for i := uint(0); i != 60; i++ {
		DefaultDeck = append(DefaultDeck, i+1)
	}
}
