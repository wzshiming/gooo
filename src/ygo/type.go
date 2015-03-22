package ygo

// 卡牌类型 Card Types
type LC_TYPE uint32

const (
	LC_None            = LC_TYPE(1 << (32 - 1 - iota))
	LC_OrdinaryMonster //普通怪兽 黄色
	LC_EffectMonster   //效果怪兽 橙色
	LC_FusionMonster   //融合怪兽 紫色
	LC_ExcessMonster   //超量怪兽 黑色
	LC_HomologyMonster //同调怪兽 白色
	LC_RiteMonster     //仪式怪兽 蓝色
	LC_OrdinaryMagic   //普通魔法
	LC_RiteMagic       //仪式魔法
	LC_SustainsMagic   //永续魔法 速度2
	LC_EquipMagic      //装备魔法
	LC_PlaceMagic      //场地魔法
	LC_RushMagic       //速攻魔法
	LC_OrdinaryTrap    //普通陷阱 速度2
	LC_SustainsTrap    //永续陷阱 速度2
	LC_ReactionTrap    //反击陷阱 速度3

	// 怪物卡
	LC_Monster = LC_OrdinaryMonster | LC_EffectMonster | LC_ExcessMonster | LC_HomologyMonster | LC_FusionMonster | LC_RiteMonster
	// 魔法卡
	LC_Magic = LC_OrdinaryMagic | LC_RiteMagic | LC_SustainsMagic | LC_EquipMagic | LC_PlaceMagic | LC_RushMagic
	// 陷阱卡
	LC_Trap = LC_OrdinaryTrap | LC_SustainsTrap | LC_ReactionTrap
	// 魔法卡与陷阱卡
	LC_MagicAndTrap = LC_Magic | LC_Trap
	// 所有类型的卡
	LC_Card = LC_Monster | LC_MagicAndTrap
)

// 表示形式 Expression
type LE_TYPE uint32

const (
	LE_None     = LE_TYPE(1 << (32 - 1 - iota))
	LE_FaceUp   // 正面朝上
	LE_FaceDown // 正面朝下
	LE_Attack   // 攻击状态
	LE_Defense  // 守备状态

	LE_FaceUpAttack    = LE_FaceUp | LE_Attack    // 表侧
	LE_FaceDownAttack  = LE_FaceUp | LE_Defense   //
	LE_FaceUpDefense   = LE_FaceDown | LE_Attack  // 侧面
	LE_FaceDownDefense = LE_FaceDown | LE_Defense // 里侧
)

// 怪兽属性 Attributes
type LA_TYPE uint32

const (
	LA_None   = LA_TYPE(1 << (32 - 1 - iota))
	LA_Earth  //地
	LA_Water  //水
	LA_Fire   //火
	LA_Wind   //风
	LA_Light  //光
	LA_Dark   //暗
	LA_Devine //神
)

// 怪兽种族 Races
type LR_TYPE uint32

const (
	LR_None        = LR_TYPE(1 << (32 - 1 - iota))
	LR_Warrior     //战士族
	LR_Spellcaster //魔法使用族
	LR_Fairy       //精灵族
	LR_Fiend       //恶魔族
	LR_Zombie      //不死族
	LR_Machine     //机械族
	LR_Aqua        //水族
	LR_Pyro        //炎族
	LR_Rock        //岩石族
	LR_WindBeast   //鸟兽族
	LR_Plant       //植物族
	LR_Insect      //昆虫族
	LR_Thunder     //雷族
	LR_Dragon      //龙族
	LR_Beast       //兽族
	LR_BeastWarror //兽战士族
	LR_Dinosaur    //恐龙族
	LR_Fish        //鱼族
	LR_Seaserpent  //海龙族
	LR_Reptile     //爬虫族
	LR_Psycho      //念动力族
	LR_Devine      //神族
	LR_CreatorGod  //创造神族
)
