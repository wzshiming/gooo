package cards

import (
	"fmt"
	"ygo"
)

var CardBag_test = ygo.NewCardVersion()

func init() {
	for i := uint(0); i != 10000; i++ {
		var co *ygo.CardOriginal
		co = &ygo.CardOriginal{
			Id:       i,
			Name:     "test_" + fmt.Sprint(i),
			Describe: "描述：测试卡 " + fmt.Sprint(i),
			Lc:       ygo.LC_OrdinaryMonster,
		}
		CardBag_test.Register(co)
	}
}
