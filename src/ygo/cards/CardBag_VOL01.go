package cards

import (
	"ygo"
)

var CardBag_VOL01 = ygo.NewCardVersion()

func init() {
	var co *ygo.CardOriginal
	co = &ygo.CardOriginal{
		Id:       21323861,
		Name:     "酸性风暴",
		Describe: "效果：场上表侧表示存在的机械族怪兽全部破坏。",
		Lc:       ygo.LC_OrdinaryMagic,
		Initiative: func(pl *ygo.Player) {
			pl.ForEachPlayer(func(p *ygo.Player) {
				p.Mzone.Find(func(c *ygo.Card) bool {
					return false
				})
			})
		},
	}
	CardBag_VOL01.Register(co)

	co = &ygo.CardOriginal{
		Id:       84285623,
		Name:     "埴轮",
		Describe: "描述：古代王墓中守护宝物的土人偶。",
		Lc:       ygo.LC_OrdinaryMonster,
		Star:     2,
		La:       ygo.LA_Earth,
		Lr:       ygo.LR_Rock,
		Attack:   500,
		Defense:  500,
	}
	CardBag_VOL01.Register(co)
}
