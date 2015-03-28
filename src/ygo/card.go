package ygo

type CardOriginal struct {
	Id       uint    // 卡牌id
	Lc       LC_TYPE // 卡牌类型
	Name     string  // 名字
	Describe string  // 描述

	// 主动效果
	Initiative Action // 发动效果

	// 怪兽卡 属性
	La      LA_TYPE // 怪兽属性
	Lr      LR_TYPE // 怪兽种族
	Star    int     // 星级
	Attack  int     // 攻击力
	Defense int     // 防御力

	// 事件
	Declaration   Action // 攻击宣言
	Damage        Action // 伤害计算
	Freedom       Action // 解放    送去墓地
	Destroy       Action // 战斗破坏 送去墓地
	Flip          Action // 反转
	Summon        Action // 召唤
	SummonCover   Action // 覆盖召唤
	SummonFlip    Action // 反转召唤
	SummonSpecial Action // 特殊召唤

}

func (co *CardOriginal) Make(ow *Player) *Card {
	return &Card{
		CardOriginal: *co,
		IsValid:      true,
		Owner:        ow,
		Le:           LE_None,
		AttackRound:  0,
	}
}

type CardsOriginal []CardOriginal

func (co *CardsOriginal) Make(ow *Player) (c *CardPile) {
	c = NewCardPile()
	for _, v := range *co {
		c.BeginPush(v.Make(ow))
	}
	return
}

type Card struct {
	CardOriginal
	Place   Cards   // 所在位置
	Owner   *Player // 所有者
	IsValid bool    // 是否有效
	Le      LE_TYPE // 表示形式
	//怪兽卡 属性
	AttackRound int // 最后攻击的回合 判断该回合是否攻击
}

func (ca *Card) Placed() {
	if ca.Place != nil {
		ca.Place.PickedFor(ca)
	}
}
