package ygo

import (
	"rego"
)

type CardOriginal struct {
	IsValid  bool    // 是否有效
	Id       uint    // 卡牌id
	Name     string  // 名字
	Password string  // 卡牌密码
	Lc       LC_TYPE // 卡牌类型

	// 怪兽卡 属性
	La      LA_TYPE // 怪兽属性
	Lr      LR_TYPE // 怪兽种族
	Level   int     // 星级
	Attack  int     // 攻击力
	Defense int     // 防御力

	Initialize Action // 初始化
	// 战斗 事件
	// 攻击方
	DeclarationPre Action // 攻击宣言之前
	DeclarationSuf Action // 攻击宣言之后
	DamageCalPre   Action // 伤害计算之前
	DamageCalSuf   Action // 伤害计算之后
	ResultPre      Action // 结果执行之前
	ResultSuf      Action // 结果执行之后

	// 承受方
	BearDeclarationPre Action // 被攻击宣言之前
	BearDeclarationSuf Action // 被攻击宣言之后
	BearDamageCalPre   Action // 被伤害计算之前
	BearDamageCalSuf   Action // 被伤害计算之后
	BearResultPre      Action // 被结果执行之前
	BearResultSuf      Action // 被结果执行之后

	// 主动效果
	Initiative    Action // 我方主阶段发动
	Removed       Action // 移除
	Freedom       Action // 解放 送去墓地
	Destroy       Action // 战斗破坏 送去墓地
	Flip          Action // 反转
	Summon        Action // 召唤
	SummonSpecial Action // 特殊召唤
}

func (co *CardOriginal) Make(ow *Player) *Card {
	c := &Card{
		baseOriginal: co,
		original:     *co,
		owner:        ow,
		summoner:     ow,
		le:           LE_None,
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
	baseOriginal *CardOriginal
	original     CardOriginal
	place        *Cards  // 所在位置
	summoner     *Player // 召唤者
	owner        *Player // 所有者
	le           LE_TYPE // 表示形式
	//怪兽卡 属性
	counter         int // 计数器
	lastAttackRound int // 最后攻击回合
	lastChangeRound int // 最后改变表示形式回合
}

func (ca *Card) HandMethods() []LI_TYPE {
	return handMethods[ca.original.Lc]
}

// 战斗
func (ca *Card) Battle(tar *Card) bool {
	pl := ca.GetSummoner()
	plt := tar.GetSummoner()

	f := tar.IsFaceDown()
	b := tar.IsAttack()

	//战斗宣言
	ca.original.DeclarationPre.Call(ca)
	tar.original.BearDeclarationPre.Call(tar)
	if f {
		tar.FaceUp()
	}
	ca.original.DeclarationSuf.Call(ca)
	tar.original.BearDeclarationSuf.Call(tar)

	//伤害判定
	ca.original.DamageCalPre.Call(ca)
	tar.original.BearDamageCalPre.Call(tar)
	var t int
	if b {
		t = ca.GetAttack() - tar.GetAttack()
	} else {
		t = ca.GetAttack() - tar.GetDefense()
	}
	ca.SetNotCanAttack()
	ca.SetNotCanChange()
	ca.original.DamageCalPre.Call(ca)
	tar.original.BearDamageCalPre.Call(tar)

	//结果执行
	ca.original.ResultPre.Call(ca)
	tar.original.BearResultPre.Call(tar)
	if b {
		if t > 0 {
			plt.ChangeHp(-t)
			tar.ToGrave()
		} else if t < 0 {
			pl.ChangeHp(t)
			ca.ToGrave()
		} else {
			tar.ToGrave()
			ca.ToGrave()
		}
	} else {
		if t > 0 {
			tar.ToGrave()
		} else if t < 0 {
			pl.ChangeHp(t)
		}
	}
	if f {
		tar.EffectFlip()
	}
	ca.original.ResultSuf.Call(ca)
	tar.original.BearResultSuf.Call(tar)

	return true
}

// 执行翻转效果
func (ca *Card) EffectFlip() bool {
	if ca.original.Flip.Call(ca) {
		pl := ca.GetSummoner()
		pl.Call(Message("发动翻转效果"))
		return true
	}
	return false
}

// 发动主动效果
func (ca *Card) EffectInitiative() bool {
	if ca.original.Initiative.IsExits() {
		ca.FaceUp()
		if ca.original.Initiative.Call(ca) {
			pl := ca.GetSummoner()
			pl.Call(Message("发动主动效果"))
			return true
		}
	}
	return false
}

// 获得召唤者
func (ca *Card) GetSummoner() *Player {
	return ca.summoner
}

// 设置召唤者
func (ca *Card) SetSummoner(c *Player) {
	ca.summoner = c
}

// 获得所有者
func (ca *Card) GetOwner() *Player {
	return ca.owner
}

// 获得 id
func (ca *Card) GetId() uint {
	return ca.original.Id
}

// 召唤
func (ca *Card) Summon() bool {
	return ca.original.Summon.Call(ca)
}

// 获得基础类型
func (ca *Card) GetBaseType() LC_TYPE {
	return ca.baseOriginal.Lc
}

// 获得类型
func (ca *Card) GetType() LC_TYPE {
	return ca.original.Lc
}

// 设置类型
func (ca *Card) SetType(l LC_TYPE) {
	ca.original.Lc = l
}

// 获得基础属性
func (ca *Card) GetBaseAttribute() LA_TYPE {
	return ca.baseOriginal.La
}

// 获得属性
func (ca *Card) GetAttribute() LA_TYPE {
	return ca.original.La
}

//  设置属性
func (ca *Card) SetAttribute(l LA_TYPE) {
	ca.original.La = l
}

// 获得基础种族
func (ca *Card) GetBaseRace() LR_TYPE {
	return ca.baseOriginal.Lr
}

// 获得种族
func (ca *Card) GetRace() LR_TYPE {
	return ca.original.Lr
}

// 设置种族
func (ca *Card) SetRace(l LR_TYPE) {
	ca.original.Lr = l
}

// 获得基础攻击
func (ca *Card) GetBaseAttack() int {
	return ca.baseOriginal.Attack
}

// 获得攻击
func (ca *Card) GetAttack() int {
	return ca.original.Attack
}

// 设置攻击
func (ca *Card) SetAttack(i int) {
	ca.original.Attack = i
}

// 获得基础防御
func (ca *Card) GetBaseDefense() int {
	return ca.baseOriginal.Defense
}

// 获得防御
func (ca *Card) GetDefense() int {
	return ca.original.Defense
}

// 设置防御
func (ca *Card) SetDefense(i int) {
	ca.original.Defense = i
}

// 获得基础等级
func (ca *Card) GetBaseLevel() int {
	return ca.baseOriginal.Level
}

// 获得等级
func (ca *Card) GetLevel() int {
	return ca.original.Level
}

// 设置等级
func (ca *Card) SetLevel(i int) {
	ca.original.Level = i
}

// 判断能够改变表示形式
func (ca *Card) IsCanChange() bool {
	if ca.lastChangeRound == ca.GetSummoner().GetRound() {
		return false
	}
	return true
}

func IsCanChange(ca *Card) bool {
	return ca.IsCanChange()
}

// 设置能够改变表示形式
func (ca *Card) SetCanChange() bool {
	ca.lastChangeRound = 0
	return true
}

func SetCanChange(ca *Card) bool {
	return ca.SetCanChange()
}

// 设置不能够改变表示形式
func (ca *Card) SetNotCanChange() bool {
	ca.lastChangeRound = ca.GetSummoner().GetRound()
	return true
}

func SetNotCanChange(ca *Card) bool {
	return ca.SetNotCanChange()
}

// 判断能够攻击
func (ca *Card) IsCanAttack() bool {
	if ca.lastAttackRound == ca.GetSummoner().GetRound() {
		return false
	}
	return true
}

func IsCanAttack(ca *Card) bool {
	return ca.IsCanAttack()
}

// 设置能够攻击
func (ca *Card) SetCanAttack() bool {
	ca.lastAttackRound = 0
	return true
}

func SetCanAttack(ca *Card) bool {
	return ca.SetCanAttack()
}

// 设置不能够攻击
func (ca *Card) SetNotCanAttack() bool {
	ca.lastAttackRound = ca.GetSummoner().GetRound()
	return true
}

func SetNotCanAttack(ca *Card) bool {
	return ca.SetNotCanAttack()
}

// 召唤
func (ca *Card) Call() {
	if (ca.GetType() & LC_Monster) != 0 {
		if ca.Summon() {
			if ca.GetSummoner().Mzone.Len() < 5 {
				ca.FaceUpAttack()
				ca.ToMzone()
			}
		}
	}
}

// 放置
func (ca *Card) Cover() {
	if (ca.GetType() & LC_PlaceMagic) != 0 { // 场地卡
		if ca.GetSummoner().Field.Len() != 0 {
			ca.GetSummoner().Field.ForEach(ToGrave)
			ca.ToField()
			ca.FaceDownAttack()
		}
	} else if (ca.GetType() & LC_MagicAndTrap) != 0 { // 魔陷
		if ca.GetSummoner().Szone.Len() < 5 {
			ca.ToSzone()
			ca.FaceDownAttack()
		}
	} else if (ca.GetType() & LC_Monster) != 0 { // 怪兽
		if ca.Summon() {
			if ca.GetSummoner().Mzone.Len() < 5 {
				ca.FaceDownDefense()
				ca.ToMzone()
				ca.SetNotCanChange()
			}
		}
	}
}

// 手牌发动
func (ca *Card) Use() {
	if (ca.GetType() & LC_PlaceMagic) != 0 { // 场地卡
		if ca.GetSummoner().Field.Len() != 0 {
			ca.GetSummoner().Field.ForEach(ToGrave)
			ca.ToField()
			ca.FaceUpAttack()
			ca.EffectInitiative()
		}
	} else if (ca.GetType() & LC_Magic) != 0 { // 魔陷
		if ca.GetSummoner().Szone.Len() < 5 {
			ca.ToSzone()
			ca.FaceUpAttack()
			ca.EffectInitiative()
		}
	}
}

// 设置表示形式
func (ca *Card) setLE(l LE_TYPE) {
	ca.GetSummoner().CallAll(ExprCard(ca, l))
	ca.le = l
	if ca.IsFaceUp() {
		ca.GetSummoner().CallAll(SetFront(ca))
	}
}

// 判断是攻击表示
func (ca *Card) IsAttack() bool {
	return (ca.le & LE_Attack) != 0
}

// 设置攻击表示
func (ca *Card) Attack() {
	ca.setLE(LE_Attack)
}

func Attack(ca *Card) {
	ca.Attack()
}

// 判断是防御表示
func (ca *Card) IsDefense() bool {
	return (ca.le & LE_Defense) != 0
}

// 设置是表示
func (ca *Card) Defense() {
	ca.setLE(LE_Defense)
}

func Defense(ca *Card) {
	ca.Defense()
}

// 判断是面朝
func (ca *Card) IsFaceUp() bool {
	return (ca.le & LE_FaceUp) != 0
}

// 设置是面朝上
func (ca *Card) FaceUp() {
	ca.setLE(LE_FaceUp)
}

func FaceUp(ca *Card) {
	ca.FaceUp()
}

// 判断是面朝下
func (ca *Card) IsFaceDown() bool {
	return (ca.le & LE_FaceDown) != 0
}

// 设置是面朝下
func (ca *Card) FaceDown() {
	ca.setLE(LE_FaceDown)
}

func FaceDown(ca *Card) {
	ca.FaceDown()
}

// 判断是面朝上攻击表示
func (ca *Card) IsFaceUpAttack() bool {
	return ca.le == LE_FaceUpAttack
}

// 设置是面朝上攻击表示
func (ca *Card) FaceUpAttack() {
	ca.setLE(LE_FaceUpAttack)
}

func FaceUpAttack(ca *Card) {
	ca.FaceUpAttack()
}

// 判断是面朝下攻击表示
func (ca *Card) IsFaceDownAttack() bool {
	return ca.le == LE_FaceDownAttack
}

// 设置是面朝下攻击表示
func (ca *Card) FaceDownAttack() {
	ca.setLE(LE_FaceDownAttack)
}

func FaceDownAttack(ca *Card) {
	ca.FaceDownAttack()
}

// 判断是面朝上防御表示
func (ca *Card) IsFaceUpDefense() bool {
	return ca.le == LE_FaceUpDefense
}

// 设置是面朝上防御表示
func (ca *Card) FaceUpDefense() {
	ca.setLE(LE_FaceUpDefense)
}

func FaceUpDefense(ca *Card) {
	ca.FaceUpDefense()
}

// 判断是面朝下防御表示
func (ca *Card) IsFaceDownDefense() bool {
	return ca.le == LE_FaceDownDefense
}

// 设置是面朝下防御表示
func (ca *Card) FaceDownDefense() {
	ca.setLE(LE_FaceDownDefense)
}

func FaceDownDefense(ca *Card) {
	ca.FaceDownDefense()
}

// 拿起  不属于任何牌堆里
func (ca *Card) Placed() {
	if ca.place != nil {
		ca.place.PickedFor(ca)
	}
}

// 改变表示形式
func (ca *Card) ChangeExpression() {
	pl := ca.GetSummoner()
	if ca.IsCanChange() {
		if ca.IsFaceDown() {
			ca.FaceUpAttack()
			ca.EffectFlip()
		} else if ca.IsFaceUpDefense() {
			ca.FaceUpAttack()
		} else if ca.IsFaceUpAttack() {
			ca.FaceUpDefense()
		} else {
			pl.Call(Message("当前状态无法改变表示形式"))
			return
		}
		pl.Call(Message("改变表示形式"))
		ca.SetNotCanChange()
	} else {
		pl.Call(Message("该怪兽当前不能改变表示形式"))
	}
}

func ChangeExpression(ca *Card) {
	ca.ChangeExpression()
}

// 普通召唤 和 上级召唤
func (ca *Card) SuperiorCall() bool {
	pl := ca.GetSummoner()
	if ca.GetLevel() > 6 {
		if pl.Mzone.Len() < 2 {
			pl.Call(Message("需要2只献祭的怪兽，无法满足召唤条件"))
			return false
		} else if pl.Mzone.Len() == 2 {
			pl.Mzone.ForEach(ToGrave)
		} else {
			pl.Call(Message("选择2只献祭的怪兽"))
			pl.SelectFunc(2, pl.Mzone, ToGrave)
		}
	} else if ca.GetLevel() > 4 {
		if pl.Mzone.Len() < 1 {
			pl.Call(Message("需要1只献祭的怪兽，无法满足召唤条件"))
			return false
		} else if pl.Mzone.Len() == 1 {
			pl.Mzone.ForEach(ToGrave)
		} else {
			pl.Call(Message("选择1只献祭的怪兽"))
			pl.SelectFunc(1, pl.Mzone, ToGrave)
		}
	}
	pl.Call(Message("召唤成功"))
	ca.SetNotCanChange()
	return true
}

func SuperiorCall(ca *Card) bool {
	return ca.SuperiorCall()
}

// 链接
func Chain(ca *Card) bool {
	return true
}

// 移动到墓地
func (ca *Card) ToGrave() bool {
	ca.GetOwner().Grave.EndPush(ca)
	return true
}

func ToGrave(ca *Card) bool {
	return ca.ToGrave()
}

// 移动到除外
func (ca *Card) ToRemoved() bool {
	ca.GetOwner().Removed.EndPush(ca)
	return true
}

func ToRemoved(ca *Card) bool {
	return ca.ToRemoved()
}

// 移动到手牌
func (ca *Card) ToHand() bool {
	ca.GetOwner().Hand.EndPush(ca)
	return true
}

func ToHand(ca *Card) bool {
	return ca.ToHand()
}

// 移动到额外
func (ca *Card) ToExtra() bool {
	ca.GetOwner().Extra.EndPush(ca)
	return true
}

func ToExtra(ca *Card) bool {
	return ca.ToExtra()
}

// 移动到怪兽
func (ca *Card) ToMzone() bool {
	ca.GetSummoner().Mzone.EndPush(ca)
	return true
}

func ToMzone(ca *Card) bool {
	return ca.ToMzone()
}

// 移动到魔法
func (ca *Card) ToSzone() bool {
	ca.GetOwner().Szone.EndPush(ca)
	return true
}

func ToSzone(ca *Card) bool {
	return ca.ToSzone()
}

// 移动到场地
func (ca *Card) ToField() bool {
	f := ca.GetOwner().Field
	for f.Len() != 0 {
		ToGrave(f.EndPop())
	}
	ca.GetOwner().Field.EndPush(ca)
	return true
}

func ToField(ca *Card) bool {
	return ca.ToField()
}
