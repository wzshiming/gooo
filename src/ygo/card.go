package ygo

import (
	"github.com/wzshiming/dispatcher"
	"github.com/wzshiming/rego"
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

}

func (co *CardOriginal) Make(pl *Player) *Card {
	c := &Card{
		Events:       dispatcher.NewForkEvent(pl.GetFork()),
		baseOriginal: co,
		owner:        pl,
		summoner:     pl,
		le:           LE_FaceDownAttack,
	}
	c.InitUint()
	c.Init()
	return c
}

var handMethods = map[LC_TYPE][]LI_TYPE{
	LC_None:            []LI_TYPE{},
	LC_Monster:         []LI_TYPE{LI_Use1, LI_Use2}, //怪兽
	LC_Magic:           []LI_TYPE{LI_Use1, LI_Use2}, //魔法
	LC_Trap:            []LI_TYPE{LI_Use2},          //陷阱
	LC_OrdinaryMonster: []LI_TYPE{LI_Use1, LI_Use2}, //普通怪兽 通常 黄色
	LC_EffectMonster:   []LI_TYPE{LI_Use1, LI_Use2}, //效果怪兽 橙色
	LC_FusionMonster:   []LI_TYPE{LI_Use1, LI_Use2}, //融合怪兽 紫色
	LC_ExcessMonster:   []LI_TYPE{LI_Use1, LI_Use2}, //超量怪兽 xyz 黑色
	LC_HomologyMonster: []LI_TYPE{LI_Use1, LI_Use2}, //同调怪兽 白色
	LC_RiteMonster:     []LI_TYPE{LI_Use1, LI_Use2}, //仪式怪兽 蓝色
	LC_OrdinaryMagic:   []LI_TYPE{LI_Use1, LI_Use2}, //普通魔法 通常
	LC_RiteMagic:       []LI_TYPE{LI_Use1, LI_Use2}, //仪式魔法
	LC_SustainsMagic:   []LI_TYPE{LI_Use1, LI_Use2}, //永续魔法 速度2
	LC_EquipMagic:      []LI_TYPE{LI_Use1, LI_Use2}, //装备魔法
	LC_PlaceMagic:      []LI_TYPE{LI_Use1, LI_Use2}, //场地魔法
	LC_RushMagic:       []LI_TYPE{LI_Use1, LI_Use2}, //速攻魔法
	LC_OrdinaryTrap:    []LI_TYPE{LI_Use2},          //普通陷阱 速度2
	LC_SustainsTrap:    []LI_TYPE{LI_Use2},          //永续陷阱 速度2
	LC_ReactionTrap:    []LI_TYPE{LI_Use2},          //反击陷阱 速度3
}

type Card struct {
	dispatcher.Events
	rego.Unique
	baseOriginal *CardOriginal
	original     CardOriginal
	place        *Group  // 所在位置
	target       *Card   // 目标卡牌
	effects      []*Card // 对此卡牌影响的卡牌
	summoner     *Player // 召唤者
	owner        *Player // 所有者
	le           LE_TYPE // 表示形式
	//怪兽卡 属性
	counter         int // 计数器
	lastAttackRound int // 最后攻击回合
	lastChangeRound int // 最后改变表示形式回合
	isValid         bool
}

func (ca *Card) ShowInfo() {
	pl := ca.GetSummoner()
	pl.CallAll(SetCardFace(ca, Arg{"攻击力": ca.GetAttack(), "防御力": ca.GetDefense()}))
}

func (ca *Card) HideInfo() {
	pl := ca.GetSummoner()
	pl.CallAll(SetCardFace(ca, Arg{}))
}

func (ca *Card) IsValid() bool {
	return ca.isValid
}

func (ca *Card) Priority() int {
	switch ca.baseOriginal.Lc {
	case LC_OrdinaryMonster: //普通怪兽 黄色
		return 1
	case LC_EffectMonster: //效果怪兽 橙色
		return 1
	case LC_OrdinaryMagic: //普通魔法 通常
		return 1
	case LC_SustainsMagic: //永续魔法
		return 1
	case LC_EquipMagic: //装备魔法
		return 1
	case LC_PlaceMagic: //场地魔法
		return 1
	case LC_RushMagic: //速攻魔法 速度2
		return 2
	case LC_OrdinaryTrap: //普通陷阱 速度2
		return 2
	case LC_SustainsTrap: //永续陷阱 速度2
		return 2
	case LC_ReactionTrap: //反击陷阱 速度3
		return 3
	}
	return 0
}

func (ca *Card) Dispatch(eventName string, args ...interface{}) {
	yg := ca.GetSummoner().Game()
	if Pay != eventName {
		ca.Events.Dispatch(Pay, append(args, eventName)...)
	}
	yg.Chain(eventName, ca, ca.GetSummoner(), append(args))
	ca.Events.Dispatch(eventName, args...)
}

func (ca *Card) GetPlace() *Group {
	return ca.place
}

func (ca *Card) HandMethods() []LI_TYPE {
	return handMethods[ca.original.Lc]
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

// 获得基础类型
func (ca *Card) GetBaseType() LC_TYPE {
	return ca.baseOriginal.Lc
}

// 获得类型
func (ca *Card) GetType() LC_TYPE {
	return ca.original.Lc
}

// 是魔法卡
func (ca *Card) IsMagicAndTrap() bool {
	return (ca.GetType() & LC_MagicAndTrap) != 0
}

// 是魔法卡
func (ca *Card) IsMagic() bool {
	return (ca.GetType() & LC_Magic) != 0
}

// 是陷阱卡
func (ca *Card) IsTrap() bool {
	return (ca.GetType() & LC_Trap) != 0
}

// 是怪兽卡
func (ca *Card) IsMonster() bool {
	return (ca.GetType() & LC_Monster) != 0
}

// 是普通怪兽
func (ca *Card) IsOrdinaryMonster() bool {
	return (ca.GetType() & LC_OrdinaryMonster) != 0
}

// 是效果怪兽
func (ca *Card) IsEffectMonster() bool {
	return (ca.GetType() & LC_EffectMonster) != 0
}

// 是融合怪兽
func (ca *Card) IsFusionMonster() bool {
	return (ca.GetType() & LC_FusionMonster) != 0
}

// 是超量怪兽
func (ca *Card) IsExcessMonster() bool {
	return (ca.GetType() & LC_ExcessMonster) != 0
}

// 是同调怪兽
func (ca *Card) IsHomologyMonster() bool {
	return (ca.GetType() & LC_HomologyMonster) != 0
}

// 是仪式怪兽
func (ca *Card) IsRiteMonster() bool {
	return (ca.GetType() & LC_RiteMonster) != 0
}

// 是普通魔法
func (ca *Card) IsOrdinaryMagic() bool {
	return (ca.GetType() & LC_OrdinaryMagic) != 0
}

// 是仪式魔法
func (ca *Card) IsRiteMagic() bool {
	return (ca.GetType() & LC_RiteMagic) != 0
}

// 是永续魔法
func (ca *Card) IsSustainsMagic() bool {
	return (ca.GetType() & LC_SustainsMagic) != 0
}

// 是装备魔法
func (ca *Card) IsEquipMagic() bool {
	return (ca.GetType() & LC_EquipMagic) != 0
}

// 是场地魔法
func (ca *Card) IsPlaceMagic() bool {
	return (ca.GetType() & LC_PlaceMagic) != 0
}

// 是速攻魔法
func (ca *Card) IsRushMagic() bool {
	return (ca.GetType() & LC_RushMagic) != 0
}

// 是普通陷阱
func (ca *Card) IsOrdinaryTrap() bool {
	return (ca.GetType() & LC_OrdinaryTrap) != 0
}

// 是永续陷阱
func (ca *Card) IsSustainsTrap() bool {
	return (ca.GetType() & LC_SustainsTrap) != 0
}

// 是反击陷阱
func (ca *Card) IsReactionTrap() bool {
	return (ca.GetType() & LC_ReactionTrap) != 0
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

// 设置表示形式
func (ca *Card) setLE(l LE_TYPE) {
	ca.le = l
	//	if (l & LE_Attack) != 0 {
	//		ca.le &= ^LE_Defense
	//	} else if (l & LE_Defense) != 0 {
	//		ca.le &= ^LE_Attack
	//	}

	//	if (l & LE_FaceUp) != 0 {
	//		ca.le &= ^LE_FaceDown
	//	} else if (l & LE_FaceDown) != 0 {
	//		ca.le &= ^LE_Attack
	//	}
	pl := ca.GetSummoner()
	pl.Dispatch(Expres, ca)
	pl.CallAll(ExprCard(ca, l))
	if ca.IsFaceUp() {
		pl.CallAll(SetFront(ca))
	}
}

// 判断是攻击表示
func (ca *Card) IsAttack() bool {
	return (ca.le & LE_Attack) == LE_Attack
}

// 设置攻击表示
func (ca *Card) Attack() {
	ca.setLE(LE_Attack | (ca.le & LE_fd))
}

func Attack(ca *Card) {
	ca.Attack()
}

// 判断是防御表示
func (ca *Card) IsDefense() bool {
	return (ca.le & LE_Defense) == LE_Defense
}

// 设置是表示
func (ca *Card) Defense() {
	ca.setLE(LE_Defense | (ca.le & LE_fd))
}

func Defense(ca *Card) {
	ca.Defense()
}

// 判断是面朝
func (ca *Card) IsFaceUp() bool {
	return (ca.le & LE_FaceUp) == LE_FaceUp
}

// 设置是面朝上
func (ca *Card) FaceUp() {
	ca.setLE(LE_FaceUp | (ca.le & LE_ad))
}

func FaceUp(ca *Card) {
	ca.FaceUp()
}

// 判断是面朝下
func (ca *Card) IsFaceDown() bool {
	return (ca.le & LE_FaceDown) == LE_FaceDown
}

// 设置是面朝下
func (ca *Card) FaceDown() {
	ca.setLE(LE_FaceDown | (ca.le & LE_ad))
}

func FaceDown(ca *Card) {
	ca.FaceDown()
}

// 判断是面朝上攻击表示
func (ca *Card) IsFaceUpAttack() bool {
	return (ca.le & LE_FaceUpAttack) == LE_FaceUpAttack
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
	return (ca.le & LE_FaceDownAttack) == LE_FaceDownAttack
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
	return (ca.le & LE_FaceUpDefense) == LE_FaceUpDefense
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
	return (ca.le & LE_FaceDownDefense) == LE_FaceDownDefense
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
		f.ForEach(ToGrave)
	}
	ca.GetOwner().Field.EndPush(ca)
	return true
}

func ToField(ca *Card) bool {
	return ca.ToField()
}

// 是在场地
func (ca *Card) IsInField() bool {
	return ca.GetPlace().GetName() == LL_Field
}

func IsInField(ca *Card) bool {
	return ca.IsInField()
}

// 是在卡组
func (ca *Card) IsInDeck() bool {
	return ca.GetPlace().GetName() == LL_Deck
}

func IsInDeck(ca *Card) bool {
	return ca.IsInDeck()
}

// 是在额外
func (ca *Card) IsInExtra() bool {
	return ca.GetPlace().GetName() == LL_Extra
}

func IsInExtra(ca *Card) bool {
	return ca.IsInExtra()
}

// 是在墓地
func (ca *Card) IsInGrave() bool {
	return ca.GetPlace().GetName() == LL_Grave
}

func IsInGrave(ca *Card) bool {
	return ca.IsInGrave()
}

// 是在手牌
func (ca *Card) IsInHand() bool {
	return ca.GetPlace().GetName() == LL_Hand
}

func IsInHand(ca *Card) bool {
	return ca.IsInHand()
}

// 是在怪兽区
func (ca *Card) IsInMzone() bool {
	return ca.GetPlace().GetName() == LL_Mzone
}

func IsInMzone(ca *Card) bool {
	return ca.IsInMzone()
}

// 是在魔陷区
func (ca *Card) IsInSzone() bool {
	return ca.GetPlace().GetName() == LL_Szone
}

func IsInSzone(ca *Card) bool {
	return ca.IsInSzone()
}

// 是在手牌
func (ca *Card) IsInRemoved() bool {
	return ca.GetPlace().GetName() == LL_Removed
}

func IsInRemoved(ca *Card) bool {
	return ca.IsInRemoved()
}
