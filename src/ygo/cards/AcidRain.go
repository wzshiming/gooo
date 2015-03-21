package cards

import (
	"ygo"
)

func init() {
	c := &ygo.CardOriginal{
		Id:       1,
		Name:     "酸性风暴",
		Describe: "场上表侧表示存在的机械族怪兽全部破坏。",
		Lc:       ygo.LC_OrdinaryMagic,
		Initiative: func() {
		},
	}
	ygo.CardBag_VOL01.Register(c)
}
