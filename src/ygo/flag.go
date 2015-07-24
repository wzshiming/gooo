package ygo

const (
	//标识符
	Pre  = "Rre"  // 之前
	Suf  = "Suf"  // 之后
	Bear = "Bear" // 被
	In   = "In"   // 进入
	Out  = "Out"  // 离开

	// 玩家事件
	RoundBegin     = "roundbegin"     // 回合开始之前
	DP             = "DP"             // 抽排阶段
	SP             = "SP"             // 预备阶段
	MP             = "MP"             // 主要阶段
	BP             = "BP"             // 战斗阶段
	EP             = "EP"             // 结束阶段
	RoundEnd       = "roundend"       // 回合结束之后
	Chain          = "chain"          // 连锁
	DetectionChain = "Detectionchain" // 检查连锁
	Draw           = "Draw"           // 每次抽排
	DrawNum        = "DrawNum"        // 每抽一张牌

	HPChange = "HPChange" // 生命值改变

	Expres = "Expres" // 表示形式改变

	// 主阶段事件
	Summon = "summon" // 召唤
	Cover  = "cover"  // 覆盖
	Pay    = "pay"    // 支付代价

	//Open   = "open"   // 主动发动
	Use1 = "use1" // 发动
	Use2 = "use2" // 覆盖

	// 战斗阶段事件
	Declaration     = "Declaration"      // 攻击宣言
	DeclarationBear = Declaration + Bear // 被攻击宣言
	DamageStep      = "DamageStep"       // 伤害步骤
	DamageStepBear  = DamageStep + Bear  // 被伤害步骤
	//DamageStepPre   = DamageStep + Pre   // 伤害步骤开始
	DamagePenalty = "DamagePenalty" // 伤害判定
	//DamageStepSuf   = DamageStep + Suf   // 伤害步骤结束

	// 连锁事件

	// 怪兽事件
	Flip            = "Flip"            // 怪兽卡 反转
	Expression      = "Expression"      // 怪兽卡 改变表示形式
	DestroyBeBattle = "DestroyBeBattle" // 怪兽卡 战斗破坏
	Freedom         = "Freedom"         // 怪兽卡 解放
	Deduct          = "Deduct"          // 怪兽卡 对玩家造成伤害

	// 卡牌事件
	InDeck     = In + string(LL_Deck)     // 进入卡组
	OutDeck    = Out + string(LL_Deck)    // 离开卡组
	InHand     = In + string(LL_Hand)     // 进入手牌
	OutHand    = Out + string(LL_Hand)    // 离开手牌
	InMzone    = In + string(LL_Mzone)    // 进入怪兽区
	OutMzone   = Out + string(LL_Mzone)   // 离开怪兽区
	InSzone    = In + string(LL_Szone)    // 进入魔陷区
	OutSzone   = Out + string(LL_Szone)   // 离开魔陷区
	InGrave    = In + string(LL_Grave)    // 进入墓地
	OutGrave   = Out + string(LL_Grave)   // 离开墓地
	InRemoved  = In + string(LL_Removed)  // 进入移除
	OutRemoved = Out + string(LL_Removed) // 离开移除
	InExtra    = In + string(LL_Extra)    // 进入额外
	OutExtra   = Out + string(LL_Extra)   // 离开额外
	InField    = In + string(LL_Field)    // 进入场地
	OutField   = Out + string(LL_Field)   // 离开场地

	Onset         = "onset"            // 主动发动
	Condition     = "condition"        // 触发条件
	Trigger       = "Trigger"          // 诱发效果
	TriggerMust   = Trigger + "Must"   // 必发
	TriggerChoose = Trigger + "Choose" // 选发

	Removed       = "removed"       // 移除
	Disabled      = "disabled"      // 失效
	Destroy       = "Destroy"       // 破坏 送去墓地
	DestroyBeRule = "destroyberule" // 规则破坏

	Realize = "Realize" // 被知道
	Discard = "discard" // 丢弃

)
