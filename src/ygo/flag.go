package ygo

const (
	FindCards = "FindCards"

	DPPre = "DPPre" // 抽排阶段之前
	DP    = "DP"
	DPSuf = "DPSuf" // 抽排阶段之后

	SPPre = "SPPre" // 准备阶段之前
	SP    = "SP"
	SPSuf = "SPSuf" // 准备阶段之后

	MPPre = "MPPre" // 主阶段之前
	MP    = "MP"
	MPSuf = "MPSuf" // 主阶段之后

	BPPre = "BPPre" // 战斗阶段之前
	BP    = "BP"
	BPSuf = "BPSuf" // 战斗阶段之后

	EPPre = "EPPre" // 结束阶段之前
	EP    = "EP"
	EPSuf = "EPSuf" // 结束阶段之后

	Draw    = "Draw"    // 每次抽排
	DrawNum = "DrawNum" // 每抽一张牌

	HPChange = "HPChange" // 生命值改变

	Expres = "Expres" // 表示形式改变

	DeclarationPre = "DeclarationPre" // 攻击宣言之前
	DeclarationSuf = "DeclarationSuf" // 攻击宣言之后
	DamageCalPre   = "DamageCalPre"   // 伤害计算之前
	DamageCalSuf   = "DamageCalSuf"   // 伤害计算之后
	ResultPre      = "ResultPre"      // 结果执行之前
	ResultSuf      = "ResultSuf"      // 结果执行之后

	BearDeclarationPre = "BearDeclarationPre" // 被攻击宣言之前
	BearDeclarationSuf = "BearDeclarationSuf" // 被攻击宣言之后
	BearDamageCalPre   = "BearDamageCalPre"   // 被伤害计算之前
	BearDamageCalSuf   = "BearDamageCalSuf"   // 被伤害计算之后
	BearResultPre      = "BearResultPre"      // 被结果执行之前
	BearResultSuf      = "BearResultSuf"      // 被结果执行之后

	Onset   = "onset"   // 我方主阶段发动
	Removed = "removed" // 移除
	Freedom = "freedom" // 解放 送去墓地
	Destroy = "Destroy" // 破坏 送去墓地
	Flip    = "flip"    // 反转
	Summon  = "summon"  // 召唤
	Cover   = "cover"   // 覆盖
	Battle  = "battle"  // 战斗
	Used    = "used"    // 发动了的
	Use     = "use"     // 手牌发动

	Discard = "discard" // 丢弃
	Deduct  = "deduct"  // 对玩家造成伤害

)
