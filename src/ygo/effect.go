package ygo

func (pl *Player) DestroyMzoneFromIndex(index int) {
	pl.Grave.BeginPush(pl.Mzone.Picked(index))
}

func (pl *Player) DestroySzoneFromIndex(index int) {
	pl.Grave.BeginPush(pl.Szone.Picked(index))
}

func (pl *Player) DestroyField() {
	pl.Grave.BeginPush(pl.Field.Picked(0))
}

func (pl *Player) SummonMzoneFromIndex(index int, c *Card) {
	pl.Mzone.Placed(index, c)
}

func (pl *Player) SummonSzoneFromIndex(index int, c *Card) {
	pl.Szone.Placed(index, c)
}

func (pl *Player) SummonField(c *Card) {
	pl.Field.Placed(0, c)
}
