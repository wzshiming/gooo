package cards

import (
	//"fmt"
	"ygo"
)

var CardBag_test = ygo.NewCardVersion()

func init() {
	original(CardBag_test)
	elses1(CardBag_test)
	elses2(CardBag_test)
	elses3(CardBag_test)
	elses4(CardBag_test)
	elses5(CardBag_test)
	elses6(CardBag_test)
	elses7(CardBag_test)
	//	for i := uint(0); i != 10000; i++ {
	//		var co *ygo.CardOriginal
	//		co = &ygo.CardOriginal{
	//			Id:            i,
	//			Name:          "test_" + fmt.Sprint(i),
	//			Describe:      "描述：测试卡 " + fmt.Sprint(i),
	//			Lc:            ygo.LC_None,
	//			Initiative:    func(pl *ygo.Player) {}, // 主动发动
	//			Declaration:   func(pl *ygo.Player) {}, // 攻击宣言
	//			Damage:        func(pl *ygo.Player) {}, // 伤害计算
	//			Freedom:       func(pl *ygo.Player) {}, // 解放    送去墓地
	//			Destroy:       func(pl *ygo.Player) {}, // 战斗破坏 送去墓地
	//			Flip:          func(pl *ygo.Player) {}, // 反转
	//			Summon:        func(pl *ygo.Player) {}, // 召唤
	//			SummonCover:   func(pl *ygo.Player) {}, // 覆盖召唤
	//			SummonFlip:    func(pl *ygo.Player) {}, // 反转召唤
	//			SummonSpecial: func(pl *ygo.Player) {}, // 特殊召唤
	//		}
	//		CardBag_test.Register(co)
	//	}
}
