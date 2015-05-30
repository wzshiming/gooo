package ygo

import (
	"rego"
)

type CardOriginal struct {
	Id       uint    // 卡牌id
	Lc       LC_TYPE // 卡牌类型
	Name     string  // 名字
	Describe string  // 描述
	Password string  // 卡牌密码

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
	IsValid       bool   // 是否有效

}

func (co *CardOriginal) Make(ow *Player) *Card {
	c := &Card{
		CardOriginal: *co,
		Owner:        ow,
		Le:           LE_None,
		AttackRound:  0,
	}
	c.InitUint()
	return c
}

var handMethods = map[LC_TYPE][]LI_TYPE{
	LC_None:            []LI_TYPE{},
	LC_Monster:         []LI_TYPE{LI_Call, LI_Cover}, //怪兽
	LC_Magic:           []LI_TYPE{LI_Use, LI_Cover},  //魔法
	LC_Trap:            []LI_TYPE{LI_Cover},          //陷阱
	LC_OrdinaryMonster: []LI_TYPE{LI_Call, LI_Cover}, //普通怪兽 通常 黄色
	LC_EffectMonster:   []LI_TYPE{LI_Call, LI_Cover}, //效果怪兽 橙色
	LC_FusionMonster:   []LI_TYPE{LI_Call, LI_Cover}, //融合怪兽 紫色
	LC_ExcessMonster:   []LI_TYPE{LI_Call, LI_Cover}, //超量怪兽 xyz 黑色
	LC_HomologyMonster: []LI_TYPE{LI_Call, LI_Cover}, //同调怪兽 白色
	LC_RiteMonster:     []LI_TYPE{LI_Call, LI_Cover}, //仪式怪兽 蓝色
	LC_OrdinaryMagic:   []LI_TYPE{LI_Use, LI_Cover},  //普通魔法 通常
	LC_RiteMagic:       []LI_TYPE{LI_Use, LI_Cover},  //仪式魔法
	LC_SustainsMagic:   []LI_TYPE{LI_Use, LI_Cover},  //永续魔法 速度2
	LC_EquipMagic:      []LI_TYPE{LI_Use, LI_Cover},  //装备魔法
	LC_PlaceMagic:      []LI_TYPE{LI_Use, LI_Cover},  //场地魔法
	LC_RushMagic:       []LI_TYPE{LI_Use, LI_Cover},  //速攻魔法
	LC_OrdinaryTrap:    []LI_TYPE{LI_Cover},          //普通陷阱 速度2
	LC_SustainsTrap:    []LI_TYPE{LI_Cover},          //永续陷阱 速度2
	LC_ReactionTrap:    []LI_TYPE{LI_Cover},          //反击陷阱 速度3
}

type Card struct {
	rego.Unique
	CardOriginal
	Place *CardPile // 所在位置
	Owner *Player   // 所有者
	Le    LE_TYPE   // 表示形式
	//怪兽卡 属性
	AttackRound int // 最后攻击的回合 判断该回合是否攻击
}

func (ca *Card) HandMethods() []LI_TYPE {
	return handMethods[ca.CardOriginal.Lc]
}

func (ca *Card) Call() {
	ca.Owner.Mzone.EndPush(ca)
	ca.Owner.CallAll(SetFront(ca))
	ca.SetLE(LE_FaceUpAttack)
}

func (ca *Card) Cover() {
	if (ca.Lc & LC_MagicAndTrap) != 0 {
		ca.Owner.Szone.EndPush(ca)
		ca.SetLE(LE_FaceDownAttack)
	} else if (ca.Lc & LC_Monster) != 0 {
		ca.Owner.Mzone.EndPush(ca)
		ca.SetLE(LE_FaceDownDefense)
	}
}

func (ca *Card) Use() {
	ca.Owner.Szone.EndPush(ca)
	ca.Owner.CallAll(SetFront(ca))
	ca.SetLE(LE_FaceDownDefense)
}

func (ca *Card) SetLE(l LE_TYPE) {
	ca.Le = l
	ca.Owner.CallAll(ExprCard(ca, l))
}

func (ca *Card) Placed() {
	if ca.Place != nil {
		ca.Place.PickedFor(ca)
	}
}
