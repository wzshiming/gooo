package ygo

import (
	"rego"
)

type CardOriginal struct {
	Id       uint    // 卡牌id
	Name     string  // 名字
	Describe string  // 描述
	Password string  // 卡牌密码
	Lc       LC_TYPE // 卡牌类型

	// 主动效果
	Initiative Action // 发动效果

	// 怪兽卡 属性
	La      LA_TYPE // 怪兽属性
	Lr      LR_TYPE // 怪兽种族
	Level   int     // 星级
	Attack  int     // 攻击力
	Defense int     // 防御力

	// 事件
	Declaration   Action // 攻击宣言
	Damage        Action // 伤害计算
	Freedom       Action // 解放    送去墓地
	Destroy       Action // 战斗破坏 送去墓地
	Flip          Action // 反转
	Summon        Action // 召唤
	SummonSpecial Action // 特殊召唤
	IsValid       bool   // 是否有效

}

func (co *CardOriginal) Make(ow *Player) *Card {
	c := &Card{
		original: *co,
		Owner:    ow,
		Le:       LE_FaceUpAttack,
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
	original CardOriginal
	Place    *Cards  // 所在位置
	Owner    *Player // 所有者
	Le       LE_TYPE // 表示形式
	//怪兽卡 属性
	Counter  int // 计数器
	UseRound int // 最后攻击的回合 判断该回合是否攻击
}

func (ca *Card) GetId() uint {
	return ca.original.Id
}

func (ca *Card) Summon() bool {
	return ca.original.Summon.Call(ca)
}

func (ca *Card) GetType() LC_TYPE {
	return ca.original.Lc
}

func (ca *Card) GetAttribute() LA_TYPE {
	return ca.original.La
}

func (ca *Card) GetRace() LR_TYPE {
	return ca.original.Lr
}

func (ca *Card) GetBaseAttack() int {
	return ca.original.Attack
}

func (ca *Card) GetAttack() int {
	return ca.original.Attack
}

func (ca *Card) GetBaseDefense() int {
	return ca.original.Defense
}

func (ca *Card) GetDefense() int {
	return ca.original.Defense
}

func (ca *Card) GetLevel() int {
	return ca.original.Level
}

func (ca *Card) HandMethods() []LI_TYPE {
	return handMethods[ca.original.Lc]
}

func (ca *Card) Battle(tar *Card) {
	ca.UseRound = ca.Owner.RoundSize
	if (tar.Le & LE_Attack) != 0 {
		t := ca.GetAttack() - tar.GetAttack()
		if t > 0 {
			tar.Owner.ChangeHp(-t)
			tar.Owner.Grave.EndPush(tar)
		} else if t < 0 {
			ca.Owner.ChangeHp(t)
			ca.Owner.Grave.EndPush(ca)
		} else {
			tar.Owner.Grave.EndPush(tar)
			ca.Owner.Grave.EndPush(ca)
		}
	} else {
		t := ca.GetAttack() - tar.GetDefense()
		if t > 0 {
			tar.Owner.Grave.EndPush(tar)
		} else if t < 0 {
			tar.SetLE(LE_FaceUpDefense)
			ca.Owner.ChangeHp(t)
		}
	}
}

func (ca *Card) Call() {
	if (ca.GetType() & LC_Monster) != 0 {
		if ca.Summon() {
			if ca.Owner.Mzone.Len() < 5 {
				ca.SetLE(LE_FaceUpAttack)
				ca.Owner.Mzone.EndPush(ca)
			}
		}
	}
}

func (ca *Card) Cover() {
	if (ca.GetType() & LC_MagicAndTrap) != 0 {
		if ca.Owner.Szone.Len() < 5 {
			ca.UseRound = ca.Owner.RoundSize
			ca.Owner.Szone.EndPush(ca)
			ca.SetLE(LE_FaceDownAttack)
		}
	} else if (ca.GetType() & LC_Monster) != 0 {
		if ca.Summon() {
			if ca.Owner.Mzone.Len() < 5 {
				ca.SetLE(LE_FaceDownDefense)
				ca.Owner.Mzone.EndPush(ca)
			}
		}
	}
}

func (ca *Card) Use() {
	if (ca.GetType() & LC_Magic) != 0 {
		ca.original.Initiative.Call(ca)
		ca.Owner.Grave.EndPush(ca)
	}
}

func (ca *Card) SetLE(l LE_TYPE) {
	ca.Le = l
	if (ca.Le & LE_FaceUp) != 0 {
		ca.Owner.CallAll(SetFront(ca))
	}
	ca.Owner.CallAll(ExprCard(ca, l))
}

func (ca *Card) Placed() {
	if ca.Place != nil {
		ca.Place.PickedFor(ca)
	}
}

func SuperiorCall(ca *Card) (ok bool) {
	own := ca.Owner
	if ca.GetLevel() > 6 {
		if own.Mzone.Len() < 2 {
			own.Call(Message("无法满足召唤条件"))
			return false
		}
		own.SelectTo(2, own.Mzone, own.Grave, "选择献祭的怪兽")
	} else if ca.GetLevel() > 4 && ca.Owner.Mzone.Len() < 1 {
		if own.Mzone.Len() < 1 {
			own.Call(Message("无法满足召唤条件"))
			return false
		}
		own.SelectTo(1, own.Mzone, own.Grave, "选择献祭的怪兽")
	}
	return true
}
