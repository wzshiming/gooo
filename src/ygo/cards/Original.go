
package cards

import (
    "ygo"
)


func original(cardBag *ygo.CardVersion) {
    var co *ygo.CardOriginal
    
    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE02，307，YSD05，SDM
        */
        Id:       1006,
        Password: "43793530",
        Name:     "歧蜥·魔蜥义豪", // "Giga Gagagigo"  "ギガ·ガガギゴ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Reptile, // 爬虫类
        Attack:   2450,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE02，307
        */
        Id:       1007,
        Password: "79182538",
        Name:     "暗黑之狂犬", // "Mad Dog of Darkness"  "暗黒の狂犬"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1900,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE02，307
        */
        Id:       1008,
        Password: "16587243",
        Name:     "新虫", // "Neo Bug"  "ネオバグ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   1800,
        Defense:  1700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE02，307，SD04
        */
        Id:       1009,
        Password: "42071342",
        Name:     "暗黑之海龙兵", // "Sea Serpent Warrior of Darkness"  "暗黒の海竜兵"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Seaserpent, // 海龙
        Attack:   1800,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE02，307
        */
        Id:       1010,
        Password: "78060096",
        Name:     "灭绝国王鲑", // "Terrorking Salmon"  "ジェノサイドキングサーモン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Fish, // 鱼
        Attack:   2400,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE02，307，SD03，YSD05
        */
        Id:       1011,
        Password: "05464695",
        Name:     "火炎木人18", // "Blazing Inpachi"  "火炎木人１８"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Pyro, // 炎
        Attack:   1850,
        Defense:  0,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE02，308，YSD04
        */
        Id:       1062,
        Password: "39674352",
        Name:     "巨歧蜥·魔蜥义豪", // "Gogiga Gagagigo"  "ゴギガ·ガガギゴ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     8,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Reptile, // 爬虫类
        Attack:   2950,
        Defense:  2800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        YSD01，EE02，308
        */
        Id:       1063,
        Password: "66073051",
        Name:     "杰拉的战士", // "Warrior of Zera"  "ゼラの戦士"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1600,
        Defense:  1600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，平罕NR
        EE02，308
        */
        Id:       1064,
        Password: "02468169",
        Name:     "封印师 明晴", // "Sealmaster Meisei"  "封印師 メイセイ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1100,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE02，308，SD20
        */
        Id:       1065,
        Password: "39552864",
        Name:     "神圣球体", // "Mystical Shine Ball"  "神聖なる球体"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   500,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE02，308
        */
        Id:       1066,
        Password: "65957473",
        Name:     "铁钢装甲虫", // "Metal Armored Bug"  "鉄鋼装甲虫"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     8,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   2800,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        MR，DL01
        */
        Id:       11,
        Password: "47879985",
        Name:     "王室前的守护者", // "Guardian of the Throne Room"  "王室前のガーディアン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1650,
        Defense:  1600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        银碎SER，平卡N
        WCPS2006，TU07，其他
        */
        Id:       1148,
        Password: "06740720",
        Name:     "圣夜龙", // "Seiyaryu"  "ホーリー·ナイト·ドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     7,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   2500,
        Defense:  2300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        银碎SER，平卡N，银字R
        SD11，TU04
        */
        Id:       1157,
        Password: "12493482",
        Name:     "神力女武神", // "Dunames Dark Witch"  "デュナミス·ヴァルキリア"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   1800,
        Defense:  1050,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，金字UR
        TP05，YU，其他，15AY
        */
        Id:       1163,
        Password: "99785935",
        Name:     "磁石战士α", // "Alpha The Magnet Warrior"  "磁石の戦士α"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   1400,
        Defense:  1700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，金字UR
        TP06，YU，其他，15AY
        */
        Id:       1164,
        Password: "39256679",
        Name:     "磁石战士β", // "Beta The Magnet Warrior"  "磁石の戦士β"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   1700,
        Defense:  1600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，金字UR
        YSD02，YSD04，其他
        */
        Id:       1165,
        Password: "48766543",
        Name:     "电子科技翼龙", // "Cyber-Tech Alligator"  "サイバティック·ワイバーン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Machine, // 机械
        Attack:   2500,
        Defense:  1600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR，平卡N
        TP07，YU，其他，15AY
        */
        Id:       1170,
        Password: "11549357",
        Name:     "磁石战士γ", // "Gamma The Magnet Warrior"  "磁石の戦士γ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   1500,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        MR
        */
        Id:       1191,
        Password: "55998462",
        Name:     "金属鱼", // "Metal Fish"  "メタル·フィッシュ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1600,
        Defense:  1900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        （这张卡在规则上不当作「恶魔/デーモン」卡使用）
        无限制
        金字UR
        其他
        */
        Id:       1203,
        Password: "24311372",
        Name:     "恶魔 佐亚", // "Zoa"  "デビルゾア"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     7,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   2600,
        Defense:  1900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        [14/04/27]
此卡是通常怪物。
        无限制
        金字UR，爆闪PR，金碎USR，平卡N
        PP05，LE01，GLD04，TP21，ST14
        */
        Id:       1211,
        Password: "32012841",
        Name:     "千年盾", // "Millennium Shield"  "千年の盾"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   0,
        Defense:  3000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR，金碎USR
        LE01，AT05
        */
        Id:       1212,
        Password: "05628232",
        Name:     "飞企鹅", // "Flying Penguin"  "トビペンギン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1200,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR，平卡N，金碎USR
        LE01，TP09
        */
        Id:       1213,
        Password: "71068263",
        Name:     "食人玩偶", // "Stuffed Animal"  "くいぐるみ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1200,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR，金碎USR
        LE01
        */
        Id:       1214,
        Password: "07562372",
        Name:     "超音速之眼", // "Megasonic Eye"  "メガソニック·アイ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1500,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR，金碎USR，平卡N
        LE01，TP23
        */
        Id:       1215,
        Password: "70345785",
        Name:     "三头巨龙", // "Yamadron"  "ヤマドラン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   1600,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR，平卡N，金碎USR
        LE01，TP04
        */
        Id:       1216,
        Password: "68401546",
        Name:     "妖精的赠礼", // "Fairy's Gift"  "妖精の贈りもの"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1400,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR，金碎USR，平卡N
        LE01，TP22
        */
        Id:       1217,
        Password: "44073668",
        Name:     "塔库利米诺斯", // "Takriminos"  "タクリミノス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Seaserpent, // 海龙
        Attack:   1500,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR，金碎USR
        LE01
        */
        Id:       1218,
        Password: "33734439",
        Name:     "2人3脚僵尸", // "Three-Legged Zombies"  "２人３脚ゾンビ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   1100,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR，金碎USR
        LE01
        */
        Id:       1219,
        Password: "71280811",
        Name:     "宝箱鬼", // "Yaranzo"  "ヤランゾ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   1300,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR，平卡N
        LE02，TP08，JY，SJ2
        */
        Id:       1220,
        Password: "64428736",
        Name:     "翼龙战士", // "Alligator's Sword"  "ワイバーンの戦士"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1500,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR，平卡N
        LE02，VOL07，GLD04
        */
        Id:       1221,
        Password: "24433920",
        Name:     "单摆刃拷问机械", // "Pendulum Machine"  "振り子刃の拷問機械"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1750,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，金字UR
        LE03，EX-R(EX)，Booster04，Booster R2，TP10，DPKB，KA，SK2
        */
        Id:       1227,
        Password: "97590747",
        Name:     "灯之魔精", // "La Jinn the Mystical Genie of the Lamp"  "ランプの魔精·ラ·ジーン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1800,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，金字UR，银字R
        LE04，EEN(406)，EE04，DPYG，15AY
        */
        Id:       1230,
        Password: "25652259",
        Name:     "王后骑士", // "Queen's Knight"  "クィーンズ·ナイト"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1500,
        Defense:  1600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，金字UR，银字R
        LE04，EEN(406)，EE04，DPYG，15AY
        */
        Id:       1231,
        Password: "90876561",
        Name:     "卫兵骑士", // "Jack's Knight"  "ジャックス·ナイト"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1900,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR，斜碎SCR，平卡N
        LE04，GLAS(506)
        */
        Id:       1232,
        Password: "38445524",
        Name:     "基尔·加斯", // "Gil Garth"  "ギル·ガース"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1800,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，金字UR，银字R
        LE05，DT05，DPKB，KA，SK2，其他
        */
        Id:       1240,
        Password: "14898066",
        Name:     "血腥魔兽人", // "Vorse Raider"  "ブラッド·ヴォルス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_BeastWarror, // 兽战士
        Attack:   1900,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR，平卡N
        PP01，ST13
        */
        Id:       1248,
        Password: "38999506",
        Name:     "宇宙女王", // "Cosmo Queen"  "コスモクイーン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     8,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   2900,
        Defense:  2450,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR
        PP01
        */
        Id:       1249,
        Password: "38277918",
        Name:     "新月龙", // "Mikazukinoyaiba"  "クレセント·ドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     7,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   2200,
        Defense:  2350,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR，面闪SR，平卡N
        PP01，PP03，TP04，PRC1
        */
        Id:       1250,
        Password: "64271667",
        Name:     "流星之龙", // "Meteor Dragon"  "メテオ·ドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   1800,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR
        PP01
        */
        Id:       1251,
        Password: "96967123",
        Name:     "大炮达磨", // "Dharma Cannon"  "大砲だるま"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   900,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR，平卡N
        PP01，TP05
        */
        Id:       1252,
        Password: "68638985",
        Name:     "蛙形史莱姆", // "Frog the Jam"  "カエルスライム"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   700,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR
        PP01
        */
        Id:       1253,
        Password: "59053232",
        Name:     "小水怪", // "Turu-Purun"  "ツルプルン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   450,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR，平卡N
        PP01，TP08
        */
        Id:       1254,
        Password: "59983499",
        Name:     "舞蹈精灵", // "Dancing Elf"  "ダンシング·エルフ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Devine, // 天使
        Attack:   300,
        Defense:  200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        面闪SR
        PP03，WCS2005，PRC1
        */
        Id:       1264,
        Password: "27054370",
        Name:     "火翼飞马", // "Firewing Pegasus"  "ファイヤー·ウイング·ペガサス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Beast, // 兽
        Attack:   2250,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        面闪SR，平卡N，平爆NPR
        PP03，TP19，AT03
        */
        Id:       1265,
        Password: "39111158",
        Name:     "三角魔龙", // "Tri-Horned Dragon"  "トライホーン·ドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     8,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   2850,
        Defense:  2350,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        面闪SR，斜碎SCR
        PP03，WCS2004
        */
        Id:       1266,
        Password: "76232340",
        Name:     "千年原人", // "Sengenjin"  "千年原人"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     8,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_BeastWarror, // 兽战士
        Attack:   2750,
        Defense:  2500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        面闪SR，平卡N
        PP03，TU07
        */
        Id:       1267,
        Password: "66516792",
        Name:     "邪恶骑士龙", // "Serpent Night Dragon"  "エビルナイト·ドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     7,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   2350,
        Defense:  2400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，金字UR
        TP11，其他
        */
        Id:       1287,
        Password: "11761845",
        Name:     "塔瓦弯刀恶魔", // "Beast of Talwar"  "タルワール·デーモン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   2400,
        Defense:  2150,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE02，309
        */
        Id:       1298,
        Password: "53776525",
        Name:     "义豪灵蜥", // "Gigobyte"  "ギゴバイト"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Reptile, // 爬虫类
        Attack:   350,
        Defense:  300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE02，309，TU04
        */
        Id:       1299,
        Password: "27288416",
        Name:     "悠悠", // "Mokey Mokey"  "もけもけ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   300,
        Defense:  100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平罕NR，平卡N
        EE02，309
        */
        Id:       1300,
        Password: "99171160",
        Name:     "平庸鬼", // "Kozaky"  "コザッキー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   400,
        Defense:  400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        （这张卡在规则上不当作「恶魔/デーモン」卡使用）
        无限制
        平卡N
        EE02，309
        */
        Id:       1301,
        Password: "26566878",
        Name:     "恶魔蝎", // "Fiend Scorpion"  "デビルスコーピオン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   900,
        Defense:  200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE02，309
        */
        Id:       1302,
        Password: "52550973",
        Name:     "法老的仆人", // "Pharaoh's Servant"  "ファラオのしもべ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   900,
        Defense:  0,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE02，309
        */
        Id:       1303,
        Password: "89959682",
        Name:     "王家的守护者", // "Pharaonic Protector"  "王家の守護者"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   900,
        Defense:  0,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        CA，DL01
        */
        Id:       131,
        Password: "78861134",
        Name:     "炎之剑豪", // "Darkfire Soldier #2"  "炎の剣豪"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Pyro, // 炎
        Attack:   1700,
        Defense:  1100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        CA，DL01
        */
        Id:       133,
        Password: "30655537",
        Name:     "机械猎鹰", // "Cyber Falcon"  "メカファルコン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1400,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        CA，DL01
        */
        Id:       134,
        Password: "03134241",
        Name:     "飞螳螂", // "Flying Kamakiri #2"  "フライングマンティス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   1500,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE01，CA，DL01，AT01
        */
        Id:       135,
        Password: "30532390",
        Name:     "鸟人", // "Sky Scout"  "バードマン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   1800,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        银字R，平卡N
        EE03，SOD(401)，YSD03
        */
        Id:       1360,
        Password: "13179332",
        Name:     "大木炭18", // "Charcoal Inpachi"  "大木炭１８"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Pyro, // 炎
        Attack:   100,
        Defense:  2100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE03，SOD(401)
        */
        Id:       1361,
        Password: "49563947",
        Name:     "新水魔道士", // "Neo Aqua Madoor"  "ネオアクア·マドール"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1200,
        Defense:  3000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE03，SOD(401)
        */
        Id:       1362,
        Password: "86652646",
        Name:     "骨犬 小栗子", // "Skull Dog Marron"  "骨犬マロン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1350,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE03，SOD(401)
        */
        Id:       1363,
        Password: "12057781",
        Name:     "落空哥布林", // "Goblin Calligrapher"  "スカゴブリン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   400,
        Defense:  400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        立体UTR，金字UR，爆闪PR，平卡N，面闪SR
        LB，BE01，SM，EX-R(EX)，DL02，Starter Box，DT01，YAP01，DTP01，DPKB，KA，SK2，MFC03，LC01，SD22，SD25
        */
        Id:       139,
        Password: "89631139",
        Name:     "青眼白龙", // "Blue-Eyes White Dragon"  "青眼の白龍"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     8,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   3000,
        Defense:  2500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，BE01，EX-R(EX)，VOL01，DL02，DPKB
        */
        Id:       140,
        Password: "76184692",
        Name:     "独眼巨人", // "Hitotsu-Me Giant"  "サイクロプス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_BeastWarror, // 兽战士
        Attack:   1200,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，BE01，DL02，Starter Box
        */
        Id:       142,
        Password: "32274490",
        Name:     "白骨", // "Skull Servant"  "ワイト"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   300,
        Defense:  200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE03，RDS(402)
        */
        Id:       1421,
        Password: "35322812",
        Name:     "人造木人18", // "Woodborg Inpachi"  "人造木人１８"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Machine, // 机械
        Attack:   500,
        Defense:  2500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE03，RDS(402)
        */
        Id:       1422,
        Password: "62327910",
        Name:     "强力守卫", // "Mighty Guard"  "マイティガード"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   500,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE03，RDS(402)
        */
        Id:       1423,
        Password: "08715625",
        Name:     "魔货物车辆 博科伊奇", // "Bokoichi the Freightening Car"  "魔貨物車両 ボコイチ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   500,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE03，RDS(402)，SD08，YSD03
        */
        Id:       1424,
        Password: "34100324",
        Name:     "鹰身少女", // "Harpie Girl"  "ハーピィ·ガール"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   500,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR，爆闪PR，银字R，平卡N，立体UTR
        LB，BE01，PP04，LN，EX-R(EX)，VOL01，DL02，SD06，DT01，WJMP，DPYG，DTP01，YU，SY2，LC01，15AY
        */
        Id:       143,
        Password: "46986414",
        Name:     "黑魔术师", // "Dark Magician"  "ブラック·マジシャン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     7,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   2500,
        Defense:  2100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR，银字R，立体UTR，平卡N
        LB，BE01，LE02，PH，EX-R(EX)，VOL01，DL02，Booster R1，15AY
        */
        Id:       144,
        Password: "06368038",
        Name:     "暗黑骑士 盖亚", // "Gaia The Fierce Knight"  "暗黒騎士ガイア"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     7,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   2300,
        Defense:  2100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，金字UR，金碎USR
        LB，BE01，LE02，EX-R(EX)，DL02，Starter Box，YAP01，15AY
        */
        Id:       145,
        Password: "91152256",
        Name:     "精灵剑士", // "Celtic Guardian"  "エルフの剣士"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1400,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，DL02，Starter Box，TP19
        */
        Id:       146,
        Password: "89091579",
        Name:     "昆虫人", // "Basic Insect"  "昆虫人間"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   500,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，BE01，EX-R(EX)，VOL01，DL02，15AY
        */
        Id:       147,
        Password: "40374923",
        Name:     "猛犸的墓场", // "Mammoth Graveyard"  "マンモスの墓場"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Dinosaur, // 恐龙
        Attack:   1200,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，BE01，EX-R(EX)，VOL01，DL02，YU，AT05
        */
        Id:       148,
        Password: "90357090",
        Name:     "银牙狼", // "Silver Fang"  "シルバー·フォング"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1200,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，金碎USR
        LB，DL02，Starter Box，Booster R1
        */
        Id:       149,
        Password: "77827521",
        Name:     "地狱的裁判", // "Trial of Nightmare"  "地獄の裁判"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1300,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE03，FET(403)，SD04
        */
        Id:       1497,
        Password: "36119641",
        Name:     "太空翻车鱼", // "Space Mambo"  "スペースマンボウ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Fish, // 鱼
        Attack:   1700,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE03，FET(403)
        */
        Id:       1498,
        Password: "62113340",
        Name:     "神龙 末日龙", // "Divine Dragon Ragnarok"  "神竜 ラグナロク"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   1500,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE03，FET(403)
        */
        Id:       1499,
        Password: "08508055",
        Name:     "格斗鼠 吱助", // "Chu-Ske the Mouse Fighter"  "格闘ねずみ チュー助"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1200,
        Defense:  0,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，金碎USR
        LB，DL02，Starter Box，Booster R1
        */
        Id:       150,
        Password: "00032864",
        Name:     "第13人的埋葬者", // "The 13th Grave"  "１３人目の埋葬者"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   1200,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE03，FET(403)
        */
        Id:       1500,
        Password: "35052053",
        Name:     "甲虫装甲骑士", // "Insect Knight"  "甲虫装甲騎士"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   1900,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE01，DL02，Starter Box
        */
        Id:       151,
        Password: "34460851",
        Name:     "火焰操纵者", // "Flame Manipulator"  "炎を操る者"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   900,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，金碎USR
        LB，BE01，DL02，Starter Box，Booster R1
        */
        Id:       152,
        Password: "53375573",
        Name:     "深渊的冥王", // "Dark King of the Abyss"  "深淵の冥王"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1200,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，DL02，Starter Box
        */
        Id:       153,
        Password: "02863439",
        Name:     "幻象鸟", // "Fiend Reflection #2"  "ミラージュ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   1100,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金碎USR，平罕NR，平卡N
        LB，BE01，DL02，Starter Box，Booster R1
        */
        Id:       154,
        Password: "85639257",
        Name:     "水魔道士", // "Aqua Madoor"  "アクア·マドール"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1200,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，DL02，Starter Box
        */
        Id:       155,
        Password: "57305373",
        Name:     "双口的暗之支配者", // "Two-Mouth Darkruler"  "二つの口を持つ闇の支配者"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Dinosaur, // 恐龙
        Attack:   900,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，DL02，Starter Box
        */
        Id:       156,
        Password: "85309439",
        Name:     "北风与太阳", // "Ray & Temperature"  "北風と太陽"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   1000,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，DL02，Starter Box
        */
        Id:       157,
        Password: "84686841",
        Name:     "烟之王", // "King Fog"  "キング·スモーク"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1000,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，斜碎SCR
        DP01，PP08，YSD01，EE03，TLM(404)，LCGX
        */
        Id:       1575,
        Password: "21844576",
        Name:     "元素英雄 羽翼侠", // "Elemental HERO Avian"  "E·HERO フェザーマン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1000,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，斜碎SCR
        DP01，PP08，YSD01，EE03，TLM(404)，LCGX
        */
        Id:       1576,
        Password: "58932615",
        Name:     "元素英雄 爆热女郎", // "Elemental HERO Burstinatrix"  "E·HERO バーストレディ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1200,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        DP01，PP08，YSD01，YSD02，EE03，TLM(404)，LCGX
        */
        Id:       1577,
        Password: "84327329",
        Name:     "元素英雄 黏土侠", // "Elemental HERO Clayman"  "E·HERO クレイマン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   800,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，BE01，DL02，Starter Box
        */
        Id:       158,
        Password: "44287299",
        Name:     "传说的剑豪 正树", // "Masaki the Legendary Swordsman"  "伝説の剣豪 MASAKI"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1100,
        Defense:  1100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        MR，DL01
        */
        Id:       16,
        Password: "54579801",
        Name:     "满潮鱼人", // "High Tide Gyojin"  "満ち潮のマーマン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1650,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        CRV(405)，EE04
        */
        Id:       1638,
        Password: "45945685",
        Name:     "单车机人", // "Cycroid"  "サイクロイド"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   800,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        CRV(405)，EE04
        */
        Id:       1639,
        Password: "60246171",
        Name:     "他者", // "Soitsu"  "ソイツ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Devine, // 天使
        Attack:   0,
        Defense:  0,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        CRV(405)，EE04
        */
        Id:       1640,
        Password: "97240270",
        Name:     "狂怒龙虾", // "Mad Lobster"  "マッド·ロブスター"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1700,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        CRV(405)，EE04，BP03
        */
        Id:       1641,
        Password: "23635815",
        Name:     "软糖豆人", // "Jerry Beans Man"  "ジェリービーンズマン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Plant, // 植物
        Attack:   1750,
        Defense:  0,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EEN(406)，EE04
        */
        Id:       1696,
        Password: "07459013",
        Name:     "暗黑界的骑士 祖尔", // "Zure, Knight of Dark World"  "暗黒界の騎士 ズール"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1800,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        DP02，YSD02，EEN(406)，EE04
        */
        Id:       1697,
        Password: "51638941",
        Name:     "V-喷气虎", // "V-Tiger Jet"  "V－タイガー·ジェット"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1600,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EEN(406)，EE04
        */
        Id:       1698,
        Password: "97023549",
        Name:     "利刃滑冰者", // "Blade Skater"  "ブレード·スケーター"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1400,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE01，MR，DL01
        */
        Id:       18,
        Password: "03797883",
        Name:     "角子老虎机AM-7", // "Slot Machine"  "スロットマシーンAM－７"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     7,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   2000,
        Defense:  2300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        [14/04/27]
此卡是通常怪物
        无限制
        平卡N
        PG，BE01，EX-R(EX)，VOL02，DL02，SY2，ST13，ST14，15AY
        */
        Id:       181,
        Password: "15025844",
        Name:     "圣精灵", // "Mystical Elf"  "ホーリー·エルフ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   800,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL02，DL02，TP17
        */
        Id:       182,
        Password: "72842870",
        Name:     "大炮鸟", // "Tyhone"  "タイホーン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   1200,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，BE01，EX-R(EX)，VOL03，DL02，YU，SY2，15AY
        */
        Id:       183,
        Password: "32452818",
        Name:     "路易斯", // "Beaver Warrior"  "ルイーズ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_BeastWarror, // 兽战士
        Attack:   1200,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，BE01，EX-R(EX)，VOL02，DL02，15AY
        */
        Id:       185,
        Password: "28279543",
        Name:     "诅咒之龙", // "Curse of Dragon"  "カース·オブ·ドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   2000,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，BE01，EX-R(EX)，YSD02，VOL03，DL02，YU，SY2，15AY
        */
        Id:       186,
        Password: "13039848",
        Name:     "岩石巨兵", // "Giant Soldier of Stone"  "岩石の巨兵"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   1300,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，BE01，EX-R(EX)，VOL02，DL02
        */
        Id:       187,
        Password: "01784619",
        Name:     "荒野盗龙", // "Uraby"  "ワイルド·ラプター"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Dinosaur, // 恐龙
        Attack:   1500,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，EX-R(EX)，VOL03，DL02
        */
        Id:       189,
        Password: "36304921",
        Name:     "魔人 死亡撒旦", // "Witty Phantom"  "魔人デスサタン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1400,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL02，DL02，YSD06，ST12
        */
        Id:       190,
        Password: "80770678",
        Name:     "竖琴精灵", // "Spirit of the Harp"  "ハープの精"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   800,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，EX-R(EX)，VOL02，DL02
        */
        Id:       191,
        Password: "63308047",
        Name:     "魔人 泰拉", // "Terra the Terrible"  "魔人 テラ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1200,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR，平卡N，银字R，爆闪PR，平爆NPR，斜碎SCR，面闪SR
        DP03，WJC，YSD02，POTD(501)，DT01，DTP01，LCGX，DE01，SD27
        */
        Id:       1913,
        Password: "89943723",
        Name:     "元素英雄 新宇侠", // "Elemental HERO Neos"  "E·HERO ネオス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     7,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   2500,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL02，DL02，TP15
        */
        Id:       192,
        Password: "75376965",
        Name:     "恍惚的人鱼", // "Enchanting Mermaid"  "恍惚の人魚"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Fish, // 鱼
        Attack:   1200,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        POTD(501)，SD09，YSD04，DB12，DE01
        */
        Id:       1929,
        Password: "37265642",
        Name:     "剑角龙", // "Sabersaurus"  "セイバーザウルス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Dinosaur, // 恐龙
        Attack:   1900,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL03，DL02，TP23
        */
        Id:       193,
        Password: "71407486",
        Name:     "炎之魔神", // "Fireyarou"  "炎の魔神"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Pyro, // 炎
        Attack:   1300,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，斜碎SCR
        DP01，YSD01，YSD02，EE03，TLM(404)，LCGX
        */
        Id:       1994,
        Password: "20721928",
        Name:     "元素英雄 电光侠", // "Elemental HERO Sparkman"  "E·HERO スパークマン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1600,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE01，MR，DL01，PE
        */
        Id:       20,
        Password: "65570596",
        Name:     "拉弓的人鱼", // "Red Archery Girl"  "弓を引くマーメイド"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1400,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL03，DL02
        */
        Id:       209,
        Password: "85326399",
        Name:     "尖刺海龙", // "Spike Seadra"  "スパイクシードラ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Seaserpent, // 海龙
        Attack:   1600,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，EX-R(EX)，VOL03，DL02
        */
        Id:       210,
        Password: "10202894",
        Name:     "天空猎手", // "Skull Red Bird"  "スカイ·ハンター"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   1550,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        面闪SR，立体UTR，银字R
        STON(503)，YSD06，DE01
        */
        Id:       2120,
        Password: "69247929",
        Name:     "基因狼人", // "Gene-Warped Warwolf"  "ジェネティック·ワーウルフ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_BeastWarror, // 兽战士
        Attack:   2000,
        Defense:  100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL03，DL02
        */
        Id:       213,
        Password: "73051941",
        Name:     "沙石怪", // "Sand Stone"  "サンド·ストーン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   1300,
        Defense:  1600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        银字R，立体UTR，平卡N
        STON(503)，ST12，DE01
        */
        Id:       2136,
        Password: "06631034",
        Name:     "冰霜腕龙", // "Frostosaurus"  "フロストザウルス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Dinosaur, // 恐龙
        Attack:   2600,
        Defense:  1700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL03，DL02
        */
        Id:       215,
        Password: "29172562",
        Name:     "钢铁巨神像", // "Steel Ogre Grotto #1"  "鋼鉄の巨神像"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1400,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        银字R，平卡N，立体UTR
        STON(503)，YSD03，TP18，DE01
        */
        Id:       2153,
        Password: "32626733",
        Name:     "螺旋龙", // "Spiral Serpent"  "スパイラルドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     8,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Seaserpent, // 海龙
        Attack:   2900,
        Defense:  2900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL03，DL02
        */
        Id:       216,
        Password: "55444629",
        Name:     "下级龙", // "Lesser Dragon"  "レッサー·ドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   1200,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL03，DL02
        */
        Id:       217,
        Password: "55291359",
        Name:     "女夜魔骑士", // "Succubus Knight"  "サキュバス·ナイト"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1650,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        限制卡
        平卡N，银字R
        PG，BE01，VOL04，DL02，15AY
        */
        Id:       222,
        Password: "08124921",
        Name:     "被封印者的右足", // "Right Leg of the Forbidden One"  "封印されし者の右足"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   200,
        Defense:  300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        限制卡
        银字R
        PG，BE01，VOL03，DL02，15AY
        */
        Id:       223,
        Password: "44519536",
        Name:     "被封印者的左足", // "Left Leg of the Forbidden One"  "封印されし者の左足"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   200,
        Defense:  300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        限制卡
        金字UR，银字R
        PG，BE01，DL02，15AY
        */
        Id:       224,
        Password: "70903634",
        Name:     "被封印者的右腕", // "Right Arm of the Forbidden One"  "封印されし者の右腕"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   200,
        Defense:  300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        限制卡
        金字UR，银字R
        PG，BE01，DL02，15AY
        */
        Id:       225,
        Password: "07902349",
        Name:     "被封印者的左腕", // "Left Arm of the Forbidden One"  "封印されし者の左腕"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   200,
        Defense:  300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，VOL01，TP15
        */
        Id:       2257,
        Password: "38142739",
        Name:     "小天使", // "Petit Angel"  "プチテンシ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   600,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，VOL01
        */
        Id:       2258,
        Password: "09159938",
        Name:     "暗黑灰羚", // "Dark Gray"  "ダーク·グレイ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   800,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，VOL01，TP15
        */
        Id:       2259,
        Password: "15401633",
        Name:     "紫炎的影武者", // "Kagemusha of the Blue Flame"  "紫炎の影武者"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   800,
        Defense:  400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，VOL01，TP14
        */
        Id:       2260,
        Password: "85705804",
        Name:     "鸭人", // "Kurama"  "ドレイク"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   800,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，VOL01
        */
        Id:       2261,
        Password: "90963488",
        Name:     "睡眠之子", // "Nemuriko"  "眠り子"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   800,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL01
        */
        Id:       2262,
        Password: "08944575",
        Name:     "死亡之足", // "The Drdek"  "デス·フット"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   700,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，VOL01
        */
        Id:       2263,
        Password: "18710707",
        Name:     "愤怒的海王", // "The Furious Sea King"  "怒りの海王"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   800,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL01，TP16
        */
        Id:       2264,
        Password: "08783685",
        Name:     "生命之沙漏", // "Hourglass of Life"  "命の砂時計"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   700,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        （这张卡在规则上也当作「恶魔/デーモン」卡使用）
        无限制
        平卡N
        VOL01
        */
        Id:       2265,
        Password: "15150371",
        Name:     "恶魔之镜", // "Wicked Mirror"  "悪魔の鏡"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   700,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        （这张卡在规则上不当作「恶魔/デーモン」卡使用）
        无限制
        平卡N
        VOL01
        */
        Id:       2266,
        Password: "53581214",
        Name:     "火焰恶魔", // "Fire Reaper"  "ファイヤー·デビル"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   700,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，VOL01，TP01
        */
        Id:       2267,
        Password: "75356564",
        Name:     "小龙", // "Petit Dragon"  "プチリュウ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   600,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL01
        */
        Id:       2268,
        Password: "53832650",
        Name:     "巴比伦", // "Meotoko"  "バビロン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   700,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，VOL01
        */
        Id:       2269,
        Password: "53293545",
        Name:     "火炎草", // "Firegrass"  "火炎草"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Plant, // 植物
        Attack:   700,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，BE01，EX-R(EX)，VOL05，DL02，SY2，15AY
        */
        Id:       227,
        Password: "41392891",
        Name:     "小精怪", // "Feral Imp"  "グレムリン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1300,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL01
        */
        Id:       2270,
        Password: "39175982",
        Name:     "暗黑杀手", // "Winged Cleaver"  "ダークキラー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   700,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL01
        */
        Id:       2271,
        Password: "15507080",
        Name:     "黑暗随从者", // "Sectarian of Secrets"  "闇にしたがう者"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   700,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL01
        */
        Id:       2272,
        Password: "46718686",
        Name:     "海星葵", // "Hitodenchak"  "ヒトデンチャク"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   600,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL01
        */
        Id:       2273,
        Password: "15510988",
        Name:     "雷电小子", // "Thunder Kid"  "サンダー·キッズ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Thunder, // 雷
        Attack:   700,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL01
        */
        Id:       2274,
        Password: "64511793",
        Name:     "复印者", // "Eyearmor"  "コピックス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   600,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL01
        */
        Id:       2275,
        Password: "09430387",
        Name:     "雷云怪", // "Lala Li-oon"  "ララ·ライウーン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Thunder, // 雷
        Attack:   600,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL01
        */
        Id:       2276,
        Password: "47695416",
        Name:     "命运之蜡烛", // "Candle of Fate"  "運命のろうそく"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   600,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL01
        */
        Id:       2277,
        Password: "22026707",
        Name:     "黑魔族的窗帘", // "Curtain of the Dark Ones"  "黒魔族のカーテン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   600,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL01
        */
        Id:       2278,
        Password: "52800428",
        Name:     "死者之腕", // "Fiend's Hand"  "死者の腕"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   600,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL01，TP06
        */
        Id:       2279,
        Password: "75889523",
        Name:     "恶魔海狸", // "Archfiend Marmot of Nefariousness"  "デーモン·ビーバー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   400,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，BE01，EX-R(EX)，VOL04，DL02，SY2，15AY
        */
        Id:       228,
        Password: "87796900",
        Name:     "守城的翼龙", // "Winged Dragon, Guardian of the Fortress #1"  "砦を守る翼竜"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   1400,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL01
        */
        Id:       2280,
        Password: "46457856",
        Name:     "小恐龙", // "Tomozaurus"  "トモザウルス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Dinosaur, // 恐龙
        Attack:   500,
        Defense:  400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL01
        */
        Id:       2281,
        Password: "84285623",
        Name:     "埴轮", // "Haniwa"  "はにわ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   500,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL02
        */
        Id:       2283,
        Password: "56283725",
        Name:     "蜘蛛男", // "Kumootoko"  "蜘蛛男"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   700,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL02，TP09
        */
        Id:       2284,
        Password: "53153481",
        Name:     "铠甲剑尾战士", // "Armaill"  "アーメイル"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   700,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL02
        */
        Id:       2286,
        Password: "33064647",
        Name:     "独眼盾龙", // "One-Eyed Shield Dragon"  "一眼の盾竜"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   700,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL02
        */
        Id:       2288,
        Password: "63432835",
        Name:     "岩石犰狳", // "Stone Armadiller"  "ストーン·アルマジラー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   800,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL02
        */
        Id:       2289,
        Password: "20060230",
        Name:     "坚硬铠甲", // "Hard Armor"  "ハードアーマー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   300,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，金字UR，面闪SR，银字R，立体UTR，爆闪PR
        RB，BE01，LE03，SC，EX-R(EX)，VOL04，DL02，Booster R3，YAP01，DPYG，DT09，15AY
        */
        Id:       229,
        Password: "70781052",
        Name:     "恶魔召唤", // "Summoned Skull"  "デーモンの召喚"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   2500,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL02
        */
        Id:       2292,
        Password: "10859908",
        Name:     "全息投影者", // "Holograh"  "ホログラー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1100,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL02
        */
        Id:       2293,
        Password: "84794011",
        Name:     "孤寂", // "Solitude"  "ソリテュード"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_BeastWarror, // 兽战士
        Attack:   1050,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL02
        */
        Id:       2294,
        Password: "94675535",
        Name:     "拉瓦斯", // "Larvas"  "ラーバス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   800,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EX-R(EX)，VOL02
        */
        Id:       2295,
        Password: "41218256",
        Name:     "利爪杀手", // "Claw Reacher"  "キラー·ザ·クロー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1000,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL02
        */
        Id:       2296,
        Password: "17733394",
        Name:     "森之尸", // "Wood Remains"  "森の屍"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   1000,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL02
        */
        Id:       2297,
        Password: "41422426",
        Name:     "阴暗处的协力者", // "Supporter in the Shelter"  "物陰の協力者"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1000,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL02
        */
        Id:       2298,
        Password: "56342351",
        Name:     "磁力战士1号", // "M-Warrior #1"  "マグネッツ１号"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1000,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL02
        */
        Id:       2299,
        Password: "92731455",
        Name:     "磁力战士2号", // "M-Warrior #2"  "マグネッツ２号"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   500,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，VOL04，DL02，JY
        */
        Id:       230,
        Password: "68846917",
        Name:     "岩窟魔人", // "Rock Ogre Grotto #1"  "岩窟魔人オーガ·ロック"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   800,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL02
        */
        Id:       2300,
        Password: "24194033",
        Name:     "泥浆兽", // "Dorover"  "ドローバ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   900,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL02
        */
        Id:       2301,
        Password: "35282433",
        Name:     "青眼银僵尸", // "Blue-Eyed Silver Zombie"  "青眼の銀ゾンビ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   900,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL02
        */
        Id:       2302,
        Password: "93553943",
        Name:     "食人花", // "Man Eater"  "マンイーター"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Plant, // 植物
        Attack:   800,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL02，TP05
        */
        Id:       2303,
        Password: "76211194",
        Name:     "暗黑拿破仑", // "Meda Bat"  "D·ナポレオン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   800,
        Defense:  400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL02
        */
        Id:       2304,
        Password: "24348204",
        Name:     "魅惑的怪盗", // "Phantom Thief"  "魅惑の怪盗"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   700,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，银字R
        EX-R(EX)，VOL03，Booster02
        */
        Id:       2305,
        Password: "48365709",
        Name:     "暗杀者", // "Ansatsu"  "アサシン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1700,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        银字R
        VOL03，Booster02
        */
        Id:       2306,
        Password: "36904469",
        Name:     "卡库塔斯", // "Akihiron"  "カクタス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1700,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        银字R
        VOL03，Booster02
        */
        Id:       2307,
        Password: "72299832",
        Name:     "机械巨兵", // "Giant Mech-soldier"  "機械の巨兵"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1750,
        Defense:  1900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        银字R
        EX-R(EX)，VOL03，Booster02，PE
        */
        Id:       2308,
        Password: "91939608",
        Name:     "神圣人偶", // "Rogue Doll"  "ホーリー·ドール"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1600,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        银字R，平卡N
        VOL03，Booster02
        */
        Id:       2309,
        Password: "98795934",
        Name:     "魔加农", // "Mabarrel"  "マキャノン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1700,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，BE01，VOL04，DL02，JY，SJ2
        */
        Id:       231,
        Password: "15480588",
        Name:     "铠蜥蜴", // "Armored Lizard"  "鎧蜥蜴"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Reptile, // 爬虫类
        Attack:   1500,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL03
        */
        Id:       2310,
        Password: "17535588",
        Name:     "甲化海星", // "Armored Starfish"  "アーマード·スターフィッシュ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   850,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL03
        */
        Id:       2311,
        Password: "16353197",
        Name:     "鲜血吮吸者", // "Drooling Lizard"  "生き血をすするもの"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Reptile, // 爬虫类
        Attack:   900,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL03
        */
        Id:       2312,
        Password: "06367785",
        Name:     "艾尔丁", // "Eldeen"  "エルディーン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   950,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL03
        */
        Id:       2313,
        Password: "34536276",
        Name:     "蟹蛛", // "Spider Crab"  "ガニグモ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   600,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL03
        */
        Id:       2314,
        Password: "68928540",
        Name:     "螳螂杀手", // "Kamakiriman"  "カマキラー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   1150,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL03
        */
        Id:       2315,
        Password: "89904598",
        Name:     "恐龙人", // "Anthrosaurus"  "恐竜人"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Dinosaur, // 恐龙
        Attack:   1000,
        Defense:  850,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL03
        */
        Id:       2316,
        Password: "98818516",
        Name:     "杀人熊猫", // "Frenzied Panda"  "キラーパンダ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1200,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL03，TP18
        */
        Id:       2317,
        Password: "93788854",
        Name:     "彷徨的亡者", // "The Wandering Doomed"  "さまよえる亡者"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   800,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL03
        */
        Id:       2318,
        Password: "45042329",
        Name:     "地雷兽", // "Tripwire Beast"  "地雷獣"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Thunder, // 雷
        Attack:   1200,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL03
        */
        Id:       2319,
        Password: "54844990",
        Name:     "死亡潜行者", // "Skull Stalker"  "デス·ストーカー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   900,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，VOL05，DL02
        */
        Id:       232,
        Password: "88979991",
        Name:     "杀人蜂", // "Killer Needle"  "キラー·ビー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   1200,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL03
        */
        Id:       2320,
        Password: "33178416",
        Name:     "奈耳", // "Misairuzame"  "ナイル"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Fish, // 鱼
        Attack:   1400,
        Defense:  1600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL03
        */
        Id:       2321,
        Password: "40200834",
        Name:     "睡狮子", // "Sleeping Lion"  "眠れる獅子"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   700,
        Defense:  1700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL03
        */
        Id:       2322,
        Password: "08058240",
        Name:     "封印之锁", // "Binding Chain"  "封印の鎖"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   1000,
        Defense:  1100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PG，VOL03
        */
        Id:       2323,
        Password: "43500484",
        Name:     "魔界之棘", // "Darkworld Thorns"  "魔界のイバラ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Plant, // 植物
        Attack:   1200,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL03
        */
        Id:       2324,
        Password: "17238333",
        Name:     "阁楼上的妖怪", // "Wretched Ghost of the Attic"  "屋根裏の物の怪"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   550,
        Defense:  400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，VOL04
        */
        Id:       2330,
        Password: "81386177",
        Name:     "神鱼", // "Bottom Dweller"  "神魚"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Fish, // 鱼
        Attack:   1650,
        Defense:  1700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，VOL04
        */
        Id:       2331,
        Password: "86088138",
        Name:     "耳天使", // "Ocubeam"  "エンゼル·イヤーズ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   1550,
        Defense:  1650,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，EX-R(EX)，VOL04
        */
        Id:       2332,
        Password: "16972957",
        Name:     "死之沉默天使 多玛", // "Doma The Angel of Silence"  "死の沈黙の天使 ドマ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Devine, // 天使
        Attack:   1600,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，VOL04
        */
        Id:       2333,
        Password: "80141480",
        Name:     "猎手蜘蛛", // "Hunter Spider"  "ハンター·スパイダー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   1600,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，VOL04
        */
        Id:       2334,
        Password: "55784832",
        Name:     "莫林芬", // "Morinphen"  "モリンフェン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1550,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，EX-R(EX)，VOL04
        */
        Id:       2335,
        Password: "93221206",
        Name:     "古代精灵", // "Ancient Elf"  "エンシェント·エルフ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1450,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EX-R(EX)，VOL04
        */
        Id:       2336,
        Password: "45121025",
        Name:     "黑影之鬼王", // "Ogre of the Black Shadow"  "黒い影の鬼王"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_BeastWarror, // 兽战士
        Attack:   1200,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL04
        */
        Id:       2337,
        Password: "68870276",
        Name:     "绝对鸟", // "Fiend Reflection#1"  "アブソリューター"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   1300,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，VOL04
        */
        Id:       2338,
        Password: "69572024",
        Name:     "舌鱼", // "Tongyo"  "舌魚"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Fish, // 鱼
        Attack:   1350,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EX-R(EX)，VOL04
        */
        Id:       2339,
        Password: "81057959",
        Name:     "龙人", // "D. Human"  "ドラゴヒューマン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1300,
        Defense:  1100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，BE01，VOL04，DL02
        */
        Id:       234,
        Password: "76812113",
        Name:     "鹰身女郎", // "Harpie Lady"  "ハーピィ·レディ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   1300,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL04
        */
        Id:       2340,
        Password: "42348802",
        Name:     "虎纹龙", // "Trakodon"  "トラコドン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Dinosaur, // 恐龙
        Attack:   1300,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL04
        */
        Id:       2341,
        Password: "55691901",
        Name:     "大嘴", // "Great Bill"  "グレード·ビル"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1250,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL04
        */
        Id:       2342,
        Password: "55014050",
        Name:     "水之少女", // "Water Girl"  "ウォーター·ガール"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1250,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL04
        */
        Id:       2343,
        Password: "25882881",
        Name:     "死神骷髅", // "Dokuroizo the Grim Reaper"  "死神のドクロイゾ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   900,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EX-R(EX)，VOL04
        */
        Id:       2344,
        Password: "41949033",
        Name:     "暗之暗杀者", // "Dark Assailant"  "闇の暗殺者"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   1200,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL04
        */
        Id:       2345,
        Password: "46247516",
        Name:     "阴阳师 道", // "Tao the Chanter"  "陰陽師 タオ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1200,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL04
        */
        Id:       2346,
        Password: "01761063",
        Name:     "猫妖精", // "Nekogal #1"  "キャッツ·フェアリー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1100,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL04
        */
        Id:       2347,
        Password: "67841515",
        Name:     "枪管百合", // "Barrel Lily"  "マグナム·リリィ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Plant, // 植物
        Attack:   1100,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，VOL04，TP12
        */
        Id:       2348,
        Password: "21817254",
        Name:     "大雷电球", // "Mega Thunderball"  "メガ·サンダーボール"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Thunder, // 雷
        Attack:   750,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL04
        */
        Id:       2349,
        Password: "02830619",
        Name:     "火焰毒蛇", // "Flame Viper"  "フレイム·ヴァイパー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Pyro, // 炎
        Attack:   400,
        Defense:  450,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL04
        */
        Id:       2350,
        Password: "81179446",
        Name:     "青虫", // "Kattapillar"  "青虫"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   250,
        Defense:  300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，VOL04
        */
        Id:       2356,
        Password: "78780140",
        Name:     "特伦特", // "Trent"  "トレント"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Plant, // 植物
        Attack:   1500,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        银字R
        VOL05，Booster05
        */
        Id:       2358,
        Password: "77998771",
        Name:     "乌鸦天狗", // "Crow Goblin"  "カラス天狗"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   1850,
        Defense:  1600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        （这张卡在规则上也当作「异虫」卡使用）
        无限制
        银字R
        VOL05，Booster05
        */
        Id:       2359,
        Password: "51228280",
        Name:     "迷宫的蠕虫", // "Dungeon Worm"  "ダンジョン·ワーム"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   1800,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，BE01，EX-R(EX)，VOL04，DL02，SJ2
        */
        Id:       236,
        Password: "01184620",
        Name:     "魔物狩人", // "Kojikocy"  "魔物の狩人"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1500,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        银字R
        EX-R(EX)，VOL05，Booster05
        */
        Id:       2360,
        Password: "26378150",
        Name:     "粗暴帝王", // "Rude Kaiser"  "ルード·カイザー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_BeastWarror, // 兽战士
        Attack:   1800,
        Defense:  1600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，银字R
        VOL05，SOVR(606)，PE，TP18
        */
        Id:       2361,
        Password: "99261403",
        Name:     "暗黑兔", // "Dark Rabbit"  "ダーク·ラビット"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1100,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，VOL05
        */
        Id:       2362,
        Password: "64501875",
        Name:     "响女", // "Hibikime"  "響女"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1450,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EX-R(EX)，VOL05
        */
        Id:       2363,
        Password: "46474915",
        Name:     "魔法幽灵", // "Magical Ghost"  "マジカル·ゴースト"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   1300,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，VOL05
        */
        Id:       2364,
        Password: "76446915",
        Name:     "圆盘魔术师", // "Disk Magician"  "ディスク·マジシャン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1350,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL05
        */
        Id:       2365,
        Password: "07670542",
        Name:     "生化植物", // "Bio Plant"  "B·プラント"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   600,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL05
        */
        Id:       2366,
        Password: "18180762",
        Name:     "泥浆潜居者", // "The Thing That Hides in the Mud"  "泥に潜み棲むもの"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   1200,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL05
        */
        Id:       2367,
        Password: "10315429",
        Name:     "切割机器人", // "Yaiba Robo"  "カッター·ロボ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1000,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL05
        */
        Id:       2368,
        Password: "98582704",
        Name:     "翼蛋精灵", // "Wing Egg Elf"  "ウィング·エッグ·エルフ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   500,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL05
        */
        Id:       2369,
        Password: "38982356",
        Name:     "冰", // "Hyo"  "氷"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   800,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EX-R(EX)，VOL05
        */
        Id:       2370,
        Password: "75499502",
        Name:     "主人与能手", // "Master & Expert"  "マスター·アン·エキスパート"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1200,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL05
        */
        Id:       2371,
        Password: "29802344",
        Name:     "蛇椰树", // "Snakeyashi"  "スネーク·パーム"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Plant, // 植物
        Attack:   1000,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL05
        */
        Id:       2372,
        Password: "16246527",
        Name:     "铠鼠", // "Armored Rat"  "鎧ネズミ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   950,
        Defense:  1100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EX-R(EX)，VOL05
        */
        Id:       2373,
        Password: "97360116",
        Name:     "恶之无名战士", // "Unknown Warrior of Fiend"  "悪の無名戦士"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1000,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL05
        */
        Id:       2374,
        Password: "62671448",
        Name:     "蛤蟆仙人", // "Toad Master"  "トードマスター"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1000,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL05
        */
        Id:       2375,
        Password: "40196604",
        Name:     "暗影", // "Dark Shade"  "ダークシェイド"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1000,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL05
        */
        Id:       2376,
        Password: "77568553",
        Name:     "酸液爬虫", // "Acid Crawler"  "アシッドクロウラー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   900,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，VOL04，DL02
        */
        Id:       238,
        Password: "67494157",
        Name:     "伏地龙", // "Crawling Dragon"  "地を這うドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   1600,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL06
        */
        Id:       2381,
        Password: "23659124",
        Name:     "海之龙王", // "Sea King Dragon"  "海の竜王"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Seaserpent, // 海龙
        Attack:   2000,
        Defense:  1700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL06
        */
        Id:       2382,
        Password: "43352213",
        Name:     "猫女郎", // "Nekogal #2"  "キャット·レディ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_BeastWarror, // 兽战士
        Attack:   1900,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL06
        */
        Id:       2383,
        Password: "69893315",
        Name:     "骰子犰狳", // "Dice Armadillo"  "ダイス·アルマジロ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1650,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        ME，VOL06
        */
        Id:       2384,
        Password: "41396436",
        Name:     "苍翼冠鸟", // "Blue-Winged Crown"  "冠を戴く蒼き翼"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   1600,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL06
        */
        Id:       2385,
        Password: "94042337",
        Name:     "暴雨怪", // "Violent Rain"  "スコール"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1550,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL06
        */
        Id:       2386,
        Password: "12146024",
        Name:     "电击蜗牛", // "Bolt Escargot"  "ボルト·エスカルゴ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Thunder, // 雷
        Attack:   1400,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL06
        */
        Id:       2387,
        Password: "69669405",
        Name:     "小鬼", // "Horn Imp"  "インプ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1300,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL06
        */
        Id:       2388,
        Password: "81618817",
        Name:     "杰米亚之神", // "Lord of Zemia"  "ゼミアの神"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1300,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL06，15AY
        */
        Id:       2389,
        Password: "53829412",
        Name:     "狮鹫兽卫", // "Griffore"  "グリフォール"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1200,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，BE01，VOL05，DL02
        */
        Id:       239,
        Password: "20277860",
        Name:     "铠武者僵尸", // "Armored Zombie"  "鎧武者ゾンビ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   1500,
        Defense:  0,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL06
        */
        Id:       2390,
        Password: "15367030",
        Name:     "蟑螂球", // "Gokibore"  "ゴキボール"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   1200,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL06
        */
        Id:       2391,
        Password: "80813021",
        Name:     "犀牛", // "Torike"  "サイガー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1200,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL06
        */
        Id:       2392,
        Password: "20848593",
        Name:     "冰水", // "Ice Water"  "氷水"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1150,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL06
        */
        Id:       2393,
        Password: "41403766",
        Name:     "冻土带的大蝎", // "Giant Scorpion of Tundra"  "ツンドラの大蠍"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   1100,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        ME，VOL06
        */
        Id:       2394,
        Password: "89272878",
        Name:     "冥界的番人", // "Guardian of the Labyrinth"  "冥界の番人"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1000,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL06
        */
        Id:       2395,
        Password: "71950093",
        Name:     "铁铲粉碎机", // "Shovel Crusher"  "シャベル·クラッシャー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   900,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        ME，VOL06
        */
        Id:       2396,
        Password: "07805359",
        Name:     "雏鸡", // "Niwatori"  "コケ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   900,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL06
        */
        Id:       2397,
        Password: "42591472",
        Name:     "笑花", // "Laughing Flower"  "笑う花"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Plant, // 植物
        Attack:   900,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL06
        */
        Id:       2398,
        Password: "06297941",
        Name:     "强盗鼠", // "Burglar"  "バーグラー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   850,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        ME，VOL06
        */
        Id:       2399,
        Password: "10071456",
        Name:     "王座守护者", // "Protector of the Throne"  "王座の守護者"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   800,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL06
        */
        Id:       2400,
        Password: "80234301",
        Name:     "棱镜人", // "Prisman"  "プリズマン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   800,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL06
        */
        Id:       2401,
        Password: "06400512",
        Name:     "突击兵", // "Cyber Commander"  "コマンダー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   750,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL06
        */
        Id:       2402,
        Password: "57935140",
        Name:     "温柔天使", // "Tenderness"  "テンダネス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   700,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL06
        */
        Id:       2403,
        Password: "84103702",
        Name:     "甲壳蟹", // "Kanikabuto"  "カニカブト"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   650,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        ME，VOL06
        */
        Id:       2406,
        Password: "60862676",
        Name:     "火焰地狱犬", // "Flame Cerebrus"  "フレイム·ケルベロス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Pyro, // 炎
        Attack:   2100,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL06
        */
        Id:       2407,
        Password: "08353769",
        Name:     "空气吸收者", // "Air Eater"  "エア·イーター"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   2100,
        Defense:  1600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL07
        */
        Id:       2408,
        Password: "39239728",
        Name:     "王家守卫", // "Royal Guard"  "ロイヤルガード"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1900,
        Defense:  2200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL07
        */
        Id:       2409,
        Password: "33621868",
        Name:     "铁球巨人", // "Giganto"  "ギガント"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1700,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL07
        */
        Id:       2410,
        Password: "49587396",
        Name:     "钢铁之心", // "Ancient Tool"  "アイアン·ハート"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1700,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL07
        */
        Id:       2411,
        Password: "32344688",
        Name:     "暗黑奇美拉", // "Dark Chimera"  "ダーク·キメラ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1610,
        Defense:  1460,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL07
        */
        Id:       2412,
        Password: "33878931",
        Name:     "撞击兽", // "Togex"  "クラッシュマン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1600,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL07
        */
        Id:       2413,
        Password: "95265975",
        Name:     "贪食的食尸鬼", // "Ghoul with an Appetite"  "大食いグール"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   1600,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL07
        */
        Id:       2414,
        Password: "53713014",
        Name:     "狂鱼", // "Crazy Fish"  "クレイジー·フィッシュ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Fish, // 鱼
        Attack:   1600,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL07
        */
        Id:       2415,
        Password: "19737320",
        Name:     "看门人", // "Gatekeeper"  "ゲート·キーパー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1500,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL07
        */
        Id:       2416,
        Password: "44865098",
        Name:     "机械士兵", // "Cyber Soldier"  "機械の兵隊"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1500,
        Defense:  1700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL07
        */
        Id:       2417,
        Password: "30090452",
        Name:     "铠武者 斩鬼", // "Zanki"  "鎧武者斬鬼"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1500,
        Defense:  1700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL07
        */
        Id:       2418,
        Password: "73911410",
        Name:     "杀人机器", // "Sword Slasher"  "キラー·マシーン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1450,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL07
        */
        Id:       2419,
        Password: "22855882",
        Name:     "被诅咒的魔剑", // "Fiend Sword"  "呪われし魔剣"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1400,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL07
        */
        Id:       2420,
        Password: "92667214",
        Name:     "杀人小丑僵尸", // "Clown Zombie"  "マーダーサーカス·ゾンビ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   1350,
        Defense:  0,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL07
        */
        Id:       2421,
        Password: "89987208",
        Name:     "东方英雄", // "Hero of the East"  "東方の英雄"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1100,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL07
        */
        Id:       2422,
        Password: "66836598",
        Name:     "温蒂妮", // "Waterdragon Fairy"  "ウンディーネ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1100,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL07
        */
        Id:       2423,
        Password: "36151751",
        Name:     "菊石骑士", // "Arma Knight"  "アンモ·ナイト"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1000,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL07
        */
        Id:       2424,
        Password: "38035986",
        Name:     "红A", // "Akakieisu"  "レッド·エース"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1000,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL07，PE
        */
        Id:       2425,
        Password: "59383041",
        Name:     "卡通鳄鱼", // "Toon Alligator"  "トゥーン·アリゲーター"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Reptile, // 爬虫类
        Attack:   800,
        Defense:  1600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL07
        */
        Id:       2426,
        Password: "14708569",
        Name:     "阿罗妮", // "Arlownay"  "アルラウネ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Plant, // 植物
        Attack:   800,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        ME，VOL07
        */
        Id:       2427,
        Password: "00549481",
        Name:     "硬毛鼠", // "Prevent Rat"  "プリヴェント·ラット"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   500,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL07
        */
        Id:       2429,
        Password: "89832901",
        Name:     "密林的黑龙王", // "B. Dragon Jungle King"  "密林の黒竜王"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   2100,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        VOL07
        */
        Id:       2430,
        Password: "35712107",
        Name:     "巨大怪鸟", // "Monstrous Bird"  "巨大な怪鳥"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   2000,
        Defense:  1900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01，Booster R1
        */
        Id:       2445,
        Password: "78556320",
        Name:     "女夜魔战士", // "Vishwar Randi"  "ヴィシュワ·ランディー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   900,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01
        */
        Id:       2446,
        Password: "37243151",
        Name:     "气象控制员", // "Weather Control"  "ウェザー·コントロール"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   600,
        Defense:  400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01，Booster R1
        */
        Id:       2447,
        Password: "03732747",
        Name:     "水元素", // "Water Element"  "ウォーター·エレメント"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   900,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EX-R(EX)，Booster01，Booster R1
        */
        Id:       2448,
        Password: "15303296",
        Name:     "石像怪", // "Ryu-Kishin"  "ガーゴイル"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1000,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，面闪SR
        Booster01，Booster R1，JY，TP16，NUMH
        */
        Id:       2449,
        Password: "55550921",
        Name:     "格斗战士 阿提米特", // "Battle Warrior"  "格闘戦士アルティメーター"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   700,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01
        */
        Id:       2450,
        Password: "97843505",
        Name:     "风之番人 精", // "Djinn the Watcher of the Wind"  "風の番人 ジン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   700,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01，Booster R1
        */
        Id:       2451,
        Password: "60589682",
        Name:     "格洛斯", // "Twin Long Rods #1"  "グロス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   900,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01
        */
        Id:       2452,
        Password: "61201220",
        Name:     "幽灵", // "Phantom Ghost"  "ゴースト"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   600,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01
        */
        Id:       2453,
        Password: "77603950",
        Name:     "萨塔那", // "Phantom Dewan"  "サターナ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   700,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01，TP17
        */
        Id:       2454,
        Password: "92944626",
        Name:     "邪炎之翼", // "Wings of Wicked Flame"  "邪炎の翼"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Pyro, // 炎
        Attack:   700,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01
        */
        Id:       2455,
        Password: "97973387",
        Name:     "大嘴鸟", // "Droll Bird"  "スピック"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   600,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01
        */
        Id:       2456,
        Password: "63515678",
        Name:     "神圣之锁", // "Mystical Capture Chain"  "聖なる鎖"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   700,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01
        */
        Id:       2457,
        Password: "63545455",
        Name:     "僵尸鬼灯", // "Mech Mole Zombie"  "ゾンビランプ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   500,
        Defense:  400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01
        */
        Id:       2458,
        Password: "13193642",
        Name:     "暗黑植物", // "Dark Plant"  "ダーク·プラント"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Plant, // 植物
        Attack:   300,
        Defense:  400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01
        */
        Id:       2459,
        Password: "81492226",
        Name:     "太古之壶", // "Ancient Jar"  "太古の壺"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   400,
        Defense:  200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，VOL04，DL02
        */
        Id:       246,
        Password: "02483611",
        Name:     "水之舞女", // "Water Omotics"  "水の踊り子"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1400,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01
        */
        Id:       2460,
        Password: "29948642",
        Name:     "鸟嘴兽", // "Dig Beak"  "ディッグ·ビーク"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   500,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01
        */
        Id:       2461,
        Password: "27094595",
        Name:     "招手的墓场", // "Graveyard and the Hand of Invitation"  "手招きする墓場"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   700,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01，Booster R1
        */
        Id:       2462,
        Password: "00756652",
        Name:     "多隆", // "Doron"  "ドローン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   900,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01
        */
        Id:       2463,
        Password: "98898173",
        Name:     "溶化的赤影", // "The Melting Red Shadow"  "とろける赤き影"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   500,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01，Booster R1
        */
        Id:       2464,
        Password: "88643173",
        Name:     "梦魇蝎", // "Nightmare Scorpion"  "ナイトメア·スコーピオン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   900,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01，TP12
        */
        Id:       2465,
        Password: "99030164",
        Name:     "幸福恋人", // "Happy Lover"  "ハッピー·ラヴァー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   800,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01，Booster R1
        */
        Id:       2466,
        Password: "15042735",
        Name:     "飓风怪", // "Hurricail"  "ハリケル"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   900,
        Defense:  200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01
        */
        Id:       2467,
        Password: "49127943",
        Name:     "食人植物", // "Man-eating Plant"  "人喰い植物"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Plant, // 植物
        Attack:   800,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01，Booster R1
        */
        Id:       2468,
        Password: "88435542",
        Name:     "火眼", // "Fire Eye"  "ファイヤー·アイ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Pyro, // 炎
        Attack:   800,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01，Booster R1
        */
        Id:       2469,
        Password: "75646173",
        Name:     "法兰克斯", // "Synchar"  "ファランクス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   800,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，VOL04，DL02
        */
        Id:       247,
        Password: "58314394",
        Name:     "陆战型战斗艇", // "Ground Attacker Bugroth"  "陸戦型 バグロス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1500,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01
        */
        Id:       2470,
        Password: "49258578",
        Name:     "地狱之门", // "Gate Deeg"  "ヘルゲート·ディーグ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Beast, // 兽
        Attack:   700,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01
        */
        Id:       2471,
        Password: "03985011",
        Name:     "神圣能量", // "Lucky Trinket"  "ホーリー·パワー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   600,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01，Booster R1
        */
        Id:       2472,
        Password: "02957055",
        Name:     "魔头邪龙", // "Wicked Dragon with the Ersatz Head"  "魔頭を持つ邪竜"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   900,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        （这张卡在规则上不当作「恶魔/デーモン」卡使用）
        无限制
        平卡N
        Booster01
        */
        Id:       2473,
        Password: "64154377",
        Name:     "未熟的恶魔", // "Embryonic Beast"  "未熟な悪魔"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   500,
        Defense:  750,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        （这张卡在规则上不当作「恶魔/デーモン」卡使用）
        无限制
        平卡N
        Booster01
        */
        Id:       2474,
        Password: "83678433",
        Name:     "午夜恶魔", // "Midnight Fiend"  "ミッドナイト·デビル"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   800,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01，Booster R1
        */
        Id:       2475,
        Password: "76704943",
        Name:     "八俣龙绘卷", // "Yamatano Dragon Scroll"  "ヤマタノ竜絵巻"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   900,
        Defense:  300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster01，Booster R1
        */
        Id:       2476,
        Password: "63125616",
        Name:     "司暗的影子", // "The Shadow Who Controls the Dark"  "闇を司る影"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   800,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02
        */
        Id:       2477,
        Password: "56713552",
        Name:     "邪恶老鼠", // "Obese Marmot of Nefariousness"  "イビル·ラット"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   750,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02，Booster R1
        */
        Id:       2478,
        Password: "70924884",
        Name:     "锯足锹形虫", // "Alinsection"  "インセクション"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   950,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02，Booster R1
        */
        Id:       2479,
        Password: "92391084",
        Name:     "威尔米", // "Wilmee"  "ウィルミー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1000,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，BE01，VOL04，DL02
        */
        Id:       248,
        Password: "58192742",
        Name:     "幼虫宝宝", // "Petit Moth"  "プチモス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   300,
        Defense:  200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02，Booster R1
        */
        Id:       2480,
        Password: "17511156",
        Name:     "木灵小丑", // "Wood Clown"  "ウッド·ジョーカー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   800,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02，Booster R1
        */
        Id:       2481,
        Password: "37160778",
        Name:     "天使魔女", // "Angelwitch"  "エンジェル·魔女"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   800,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02，Booster Chronicle，Booster R1，TU04
        */
        Id:       2482,
        Password: "82065276",
        Name:     "电波英雄", // "Oscillo Hero"  "オシロ·ヒーロー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1250,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02，Booster Chronicle，Booster R1
        */
        Id:       2483,
        Password: "69750536",
        Name:     "鱼战士", // "Wow Warrior"  "魚ギョ戦士"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Fish, // 鱼
        Attack:   1250,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02
        */
        Id:       2484,
        Password: "10598400",
        Name:     "龙虾怪", // "Zarigun"  "ザリガン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   600,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02
        */
        Id:       2485,
        Password: "40387124",
        Name:     "开在深渊的花", // "Abyss Flower"  "深淵に咲く花"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Plant, // 植物
        Attack:   750,
        Defense:  400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02，Booster R1
        */
        Id:       2486,
        Password: "72269672",
        Name:     "岩石幽灵", // "Stone Ghost"  "ストーン·ゴースト"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   1200,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02，Booster R1
        */
        Id:       2487,
        Password: "03606209",
        Name:     "狩魂者", // "Person who hunts soul"  "魂を狩る者"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_BeastWarror, // 兽战士
        Attack:   1100,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02，Booster Chronicle，Booster R1
        */
        Id:       2488,
        Password: "34290067",
        Name:     "死亡鲨", // "Corroding Shark"  "デッド·シャーク"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   1100,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02
        */
        Id:       2489,
        Password: "72076281",
        Name:     "机械蝙蝠", // "Bat"  "バット"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Machine, // 机械
        Attack:   300,
        Defense:  350,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02，Booster Chronicle，Booster R1
        */
        Id:       2490,
        Password: "84990171",
        Name:     "豌豆战士", // "Bean Soldier"  "ビーン·ソルジャー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Plant, // 植物
        Attack:   1400,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02，TP21
        */
        Id:       2491,
        Password: "68963107",
        Name:     "书之精灵", // "Boo Koo"  "ブークー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   650,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02，Booster R1
        */
        Id:       2492,
        Password: "42625254",
        Name:     "小霸王龙", // "Little D"  "ベビー·ティーレックス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Dinosaur, // 恐龙
        Attack:   1100,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02
        */
        Id:       2493,
        Password: "21239280",
        Name:     "骨鼠", // "Bone Mouse"  "骨ネズミ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   400,
        Defense:  300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02，Booster R1，TP22
        */
        Id:       2494,
        Password: "48531733",
        Name:     "电击企鹅", // "Bolt Penguin"  "ボルト·ペンギン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Thunder, // 雷
        Attack:   1100,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02
        */
        Id:       2495,
        Password: "92409659",
        Name:     "白海豚", // "White Dolphin"  "ホワイト·ドルフィン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Fish, // 鱼
        Attack:   500,
        Defense:  400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02，Booster R1
        */
        Id:       2496,
        Password: "23032273",
        Name:     "梅基拉斯之光", // "Megirus Light"  "メギラス·ライト"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   900,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02，Booster R1
        */
        Id:       2497,
        Password: "20541432",
        Name:     "心锁妖精", // "Key Mace #2"  "ロックメイス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1050,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster02，Booster Chronicle，Booster R1
        */
        Id:       2498,
        Password: "57405307",
        Name:     "翼龙", // "Winged Dragon, Guardian of the Fortress #2"  "ワイバーン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   1200,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03，TP23
        */
        Id:       2502,
        Password: "62403074",
        Name:     "赤剑之莱蒙多斯", // "Rhaimundos of the Red Sword"  "赤き剣のライムンドス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1200,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03
        */
        Id:       2503,
        Password: "34320307",
        Name:     "有生命的花瓶", // "Living Vase"  "命ある花瓶"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Plant, // 植物
        Attack:   900,
        Defense:  1100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03
        */
        Id:       2504,
        Password: "52367652",
        Name:     "食命者", // "That Which Feeds on Life"  "命を食する者"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1200,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03，Booster R2
        */
        Id:       2505,
        Password: "46864967",
        Name:     "岩之战士", // "Minomushi Warrior"  "岩の戦士"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   1300,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03，Booster Chronicle，Booster R2
        */
        Id:       2506,
        Password: "38942059",
        Name:     "音女", // "Sonic Maid"  "音女"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1200,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EX-R(EX)，Booster03，Booster R2
        */
        Id:       2507,
        Password: "13429800",
        Name:     "大白鲨", // "Great White"  "グレート·ホワイト"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Fish, // 鱼
        Attack:   1600,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03，Booster Chronicle，Booster R2，TP20
        */
        Id:       2508,
        Password: "60802233",
        Name:     "锹甲阿尔法", // "Kuwagata α"  "クワガタ·アルファ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   1250,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03
        */
        Id:       2509,
        Password: "07526150",
        Name:     "钢铁魔人", // "Golgoil"  "ゴルゴイル"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   900,
        Defense:  1600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03，Booster R2
        */
        Id:       2510,
        Password: "55210709",
        Name:     "青玉眼怪", // "Lisark"  "サファイヤ·リサーク"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1300,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03，Booster Chronicle，Booster R2
        */
        Id:       2511,
        Password: "16899564",
        Name:     "斩首的美女", // "Beautiful Headhuntress"  "斬首の美女"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1600,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03，Booster R2
        */
        Id:       2512,
        Password: "47922711",
        Name:     "海马兽", // "Tatsunootoshigo"  "シーホース"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1350,
        Defense:  1600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03，Booster Chronicle，Booster R2
        */
        Id:       2513,
        Password: "28003512",
        Name:     "审判之手", // "The Judgement Hand"  "ジャジメント·ザ·ハンド"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1400,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03
        */
        Id:       2514,
        Password: "00732302",
        Name:     "骷髅寺院", // "Temple of Skulls"  "髑髏の寺院"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   900,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03，TP07
        */
        Id:       2515,
        Password: "84916669",
        Name:     "树精", // "Dryad"  "ドリアード"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1200,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EX-R(EX)，Booster03，Booster R2
        */
        Id:       2516,
        Password: "47060154",
        Name:     "狂战士", // "Mystic Clown"  "バーサーカー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1500,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03
        */
        Id:       2517,
        Password: "48109103",
        Name:     "复仇的河童", // "Kappa Avenger"  "復讐のカッパ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1200,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03，Booster Chronicle，Booster R2，TP20
        */
        Id:       2518,
        Password: "52584282",
        Name:     "大力独角仙", // "Hercules Beetle"  "ヘラクレス·ビートル"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   1500,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03，Booster Chronicle，Booster R2
        */
        Id:       2519,
        Password: "14037717",
        Name:     "书之精灵 飞鹰主教", // "Spirit of the Books"  "本の精霊 ホーク·ビショップ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   1400,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03，Booster Chronicle，Booster R2
        */
        Id:       2520,
        Password: "75559356",
        Name:     "魔界的机械兵", // "Cyber Soldier of Darkworld"  "魔界の機械兵"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1400,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03，TP16
        */
        Id:       2521,
        Password: "69992868",
        Name:     "缪斯天使", // "Muse-A"  "ミューズの天使"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   850,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03
        */
        Id:       2522,
        Password: "07225792",
        Name:     "蒙·拉瓦斯", // "Mon Larvas"  "モン·ラーバス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1300,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03
        */
        Id:       2523,
        Password: "41061625",
        Name:     "椰树", // "Yashinoki"  "ヤシの木"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Plant, // 植物
        Attack:   800,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03
        */
        Id:       2524,
        Password: "17115745",
        Name:     "尤蒙刚德", // "Yormungarde"  "ヨルムンガルド"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Reptile, // 爬虫类
        Attack:   1200,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03
        */
        Id:       2525,
        Password: "75850803",
        Name:     "月之魔法师", // "La Moon"  "ラムーン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1200,
        Defense:  1700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster03，Booster R2
        */
        Id:       2526,
        Password: "04392470",
        Name:     "狮子男巫", // "Leo Wizard"  "レオ·ウィザード"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1350,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster04，Booster Chronicle，Booster R2
        */
        Id:       2533,
        Password: "72929454",
        Name:     "龟鸟", // "Turtle Bird"  "タートル·バード"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1900,
        Defense:  1700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster04，Booster R2
        */
        Id:       2534,
        Password: "50176820",
        Name:     "电子鱼", // "Mech Bass"  "サイボーグ·バス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1800,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster04，Booster R2
        */
        Id:       2535,
        Password: "11250655",
        Name:     "水陆的帝王", // "Emperor of the Land and Sea"  "水陸の帝王"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Reptile, // 爬虫类
        Attack:   1800,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster04，Booster06，Booster Chronicle，Booster R2
        */
        Id:       2536,
        Password: "04179849",
        Name:     "红叶之女王", // "Queen of Autumn Leaves"  "紅葉の女王"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Plant, // 植物
        Attack:   1800,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster04，Booster R2
        */
        Id:       2537,
        Password: "02971090",
        Name:     "战神 奥利安", // "Orion the Battle Kami"  "戦いの神 オリオン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   1800,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster04，Booster R2
        */
        Id:       2538,
        Password: "34690519",
        Name:     "山之精灵", // "Spirit of the Mountain"  "山の精霊"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1300,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster04，Booster R2
        */
        Id:       2539,
        Password: "42418084",
        Name:     "诞生的天使", // "Winged Egg of New Life"  "誕生の天使"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   1400,
        Defense:  1700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster04
        */
        Id:       2540,
        Password: "53606874",
        Name:     "巨大蚂蚁", // "Big Insect"  "ビック·アント"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   1200,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster04，Booster Chronicle，Booster R2
        */
        Id:       2541,
        Password: "75582395",
        Name:     "圣鸟", // "Faith Bird"  "セイント·バード"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   1500,
        Defense:  1100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        （这张卡在规则上不当作「恶魔/デーモン」卡使用）
        无限制
        平卡N
        Booster04
        */
        Id:       2542,
        Password: "77456781",
        Name:     "恶魔乌贼", // "Fiend Kraken"  "デビル·クラーケン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1200,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster04
        */
        Id:       2543,
        Password: "17968114",
        Name:     "海原的女战士", // "Amazon of the Seas"  "海原の女戦士"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Fish, // 鱼
        Attack:   1300,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster04
        */
        Id:       2544,
        Password: "65623423",
        Name:     "杀人污泥", // "Gruesome Goo"  "キラー·ブロッブ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1300,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster04
        */
        Id:       2545,
        Password: "74277583",
        Name:     "剪刀勇士", // "Brave Scizzar"  "ブレイブ·シザー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1300,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster04
        */
        Id:       2546,
        Password: "71746462",
        Name:     "深海切割手", // "Sea Kamen"  "シーカーメン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1100,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster04
        */
        Id:       2547,
        Password: "10476868",
        Name:     "机枪岩", // "Barrel Rock"  "ガンロック"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   1000,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EX-R(EX)，Booster04
        */
        Id:       2548,
        Password: "89494469",
        Name:     "暗黑魔神 梦魇", // "Dark Titan of Terror"  "暗黒魔神 ナイトメア"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1300,
        Defense:  1100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster04
        */
        Id:       2549,
        Password: "11793047",
        Name:     "戈耳工蛋", // "Gorgon Egg"  "ゴーゴン·エッグ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   300,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster04，TP02
        */
        Id:       2550,
        Password: "20315854",
        Name:     "妖精龙", // "Fairy Dragon"  "フェアリー·ドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   1100,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster05，Booster Chronicle，Booster R3
        */
        Id:       2557,
        Password: "47319141",
        Name:     "飞鹰", // "Wing Eagle"  "ウイング·イーグル"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   1800,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster05
        */
        Id:       2558,
        Password: "38116136",
        Name:     "机器攻击者", // "Machine Attacker"  "マシン·アタッカー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1600,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster05，Booster Chronicle，Booster R3
        */
        Id:       2559,
        Password: "38289717",
        Name:     "贪尸龙", // "Crawling Dragon #2"  "屍を貪る竜"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Dinosaur, // 恐龙
        Attack:   1600,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster05
        */
        Id:       2560,
        Password: "29402771",
        Name:     "彩虹人鱼", // "Rainbow Marine Mermaid"  "レインボー·マリン·マーメイド"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Fish, // 鱼
        Attack:   1550,
        Defense:  1700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster05，Booster R3
        */
        Id:       2561,
        Password: "29491031",
        Name:     "美杜莎的亡灵", // "The Snake Hair"  "メデューサの亡霊"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   1500,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster05，Booster Chronicle，Booster R3
        */
        Id:       2562,
        Password: "93343894",
        Name:     "水之魔导师", // "Water Magician"  "水の魔導師"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1400,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster05，Booster R3
        */
        Id:       2563,
        Password: "94022093",
        Name:     "大肚海蛇", // "Behegon"  "ベヒゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1350,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster05
        */
        Id:       2564,
        Password: "11714098",
        Name:     "3万年的白龟", // "30,000-Year White Turtle"  "３万年の白亀"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1250,
        Defense:  2100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster05
        */
        Id:       2565,
        Password: "49417509",
        Name:     "狼", // "Wolf"  "オオカミ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1200,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster05
        */
        Id:       2566,
        Password: "76512652",
        Name:     "鳄鱼人", // "Krokodilus"  "クロコダイラス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Reptile, // 爬虫类
        Attack:   1100,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster05
        */
        Id:       2567,
        Password: "12436646",
        Name:     "水蛇", // "Aqua Snake"  "アクア·スネーク"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1050,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster05
        */
        Id:       2568,
        Password: "96643568",
        Name:     "气象精灵", // "Wetha"  "ウェザ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1000,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster05
        */
        Id:       2569,
        Password: "28450915",
        Name:     "来自异次元的侵略者", // "Invader from Another Dimension"  "異次元からの侵略者"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   950,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，VOL05，DL02
        */
        Id:       257,
        Password: "10538007",
        Name:     "狮子王", // "Leogun"  "レオグン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1750,
        Defense:  1550,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster05
        */
        Id:       2570,
        Password: "06103114",
        Name:     "鸟嘴蛇", // "Beaked Snake"  "くちばしヘビ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Reptile, // 爬虫类
        Attack:   800,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster05
        */
        Id:       2571,
        Password: "17441953",
        Name:     "龟狸", // "Turtle Raccoon"  "タートル·狸"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   700,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster05
        */
        Id:       2572,
        Password: "32569498",
        Name:     "屎壳郎", // "Korogashi"  "コロガーシ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   550,
        Defense:  400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster05
        */
        Id:       2573,
        Password: "60715406",
        Name:     "藤蔓植物", // "Tentacle Plant"  "テンタクル·プラント"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Plant, // 植物
        Attack:   500,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster05
        */
        Id:       2574,
        Password: "55567161",
        Name:     "戏法壶", // "Pot the Trick"  "ポット·ザ·トリック"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   400,
        Defense:  400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster05
        */
        Id:       2575,
        Password: "53830602",
        Name:     "魔界植物", // "Fungi of the Musk"  "魔界植物"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   400,
        Defense:  300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster06，Booster Chronicle，Booster R3，TP11
        */
        Id:       2580,
        Password: "48649353",
        Name:     "牛鬼", // "Ushi Oni"  "牛鬼"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   2150,
        Defense:  1950,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster06
        */
        Id:       2581,
        Password: "69780745",
        Name:     "加尔瓦斯", // "Garvas"  "ガルヴァス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   2000,
        Defense:  1700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster06，Booster Chronicle，Booster R3，PE
        */
        Id:       2582,
        Password: "62762898",
        Name:     "鹦鹉龙", // "Parrot Dragon"  "パロット·ドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   2000,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster06，Booster Chronicle，Booster R3
        */
        Id:       2583,
        Password: "95288024",
        Name:     "天空龙", // "Sky Dragon"  "天空竜"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   1900,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster06，Booster Chronicle，Booster R3，JY
        */
        Id:       2584,
        Password: "14977074",
        Name:     "加鲁撒斯", // "Garoozis"  "ガルーザス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_BeastWarror, // 兽战士
        Attack:   1800,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster06，Booster R3
        */
        Id:       2585,
        Password: "13069066",
        Name:     "剑龙", // "Sword Arm of Dragon"  "剣竜"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Dinosaur, // 恐龙
        Attack:   1750,
        Defense:  2030,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster06
        */
        Id:       2586,
        Password: "70084224",
        Name:     "猎头魔人", // "Neck Hunter"  "首狩り魔人"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1750,
        Defense:  1900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster06
        */
        Id:       2587,
        Password: "54615781",
        Name:     "风之精灵", // "Spirit of the Wind"  "風の精霊"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1700,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EX-R(EX)，Booster06，Booster R3，KA，SK2
        */
        Id:       2588,
        Password: "05053103",
        Name:     "牛头人", // "Battle Ox"  "ミノタウルス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_BeastWarror, // 兽战士
        Attack:   1700,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster06
        */
        Id:       2589,
        Password: "82818645",
        Name:     "岩石之精灵", // "Rock Spirit"  "岩石の精霊"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1650,
        Defense:  1900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster06
        */
        Id:       2590,
        Password: "81686058",
        Name:     "塞壬", // "Ill Witch"  "セイレーン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1600,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster06
        */
        Id:       2591,
        Password: "09540040",
        Name:     "岩石龟", // "Boulder Tortoise"  "岩石カメッター"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1450,
        Defense:  2200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster06，Booster R3
        */
        Id:       2592,
        Password: "85448931",
        Name:     "守卫海洋的战士", // "Sentinel of the Seas"  "海を守る戦士"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1300,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster06
        */
        Id:       2593,
        Password: "09197735",
        Name:     "龙魂之石像", // "Dragon Statue"  "竜魂の石像"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1100,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster06
        */
        Id:       2594,
        Password: "94412545",
        Name:     "机械变色龙", // "Mechaleon"  "メカレオン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Reptile, // 爬虫类
        Attack:   800,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster07
        */
        Id:       2601,
        Password: "68339286",
        Name:     "金属守护者", // "Metal Guardian"  "メタル·ガーディアン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1150,
        Defense:  2150,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster07
        */
        Id:       2602,
        Password: "36821538",
        Name:     "古代魔导士", // "Ancient Sorcerer"  "古代魔導士"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1000,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        （这张卡在规则上不当作「恶魔/デーモン」卡使用）
        无限制
        平卡N
        EX-R(EX)，Booster07
        */
        Id:       2603,
        Password: "67724379",
        Name:     "恶魔龙", // "Koumori Dragon"  "デビル·ドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   1500,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster07，TP20
        */
        Id:       2604,
        Password: "34743446",
        Name:     "碎块人", // "Blocker"  "ブロッカー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   850,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster07
        */
        Id:       2605,
        Password: "55337339",
        Name:     "转职的魔镜", // "Job-Change Mirror"  "転職の魔鏡"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   800,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        ME，Booster07
        */
        Id:       2606,
        Password: "42431843",
        Name:     "魔天老", // "Ancient Brain"  "魔天老"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1000,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster07
        */
        Id:       2607,
        Password: "68171737",
        Name:     "岩石龙", // "Stone Dragon"  "ストーン·ドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     7,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   2000,
        Defense:  2300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster07
        */
        Id:       2608,
        Password: "47986555",
        Name:     "千年石人", // "Millennium Golem"  "千年ゴーレム"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   2000,
        Defense:  2200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster07
        */
        Id:       2609,
        Password: "75390004",
        Name:     "角龙", // "Megazowler"  "メガザウラー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Dinosaur, // 恐龙
        Attack:   1800,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，EX-R(EX)，VOL04，DL02
        */
        Id:       261,
        Password: "73481154",
        Name:     "破坏的石人", // "Destroyer Golem"  "破壊のゴーレム"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   1500,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster07
        */
        Id:       2610,
        Password: "29616941",
        Name:     "美丽的魔物使", // "Beautiful Beast Trainer"  "美しき魔物使い"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1750,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster07
        */
        Id:       2611,
        Password: "97612389",
        Name:     "地狱的魔物使", // "Monster Tamer"  "地獄の魔物使い"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1800,
        Defense:  1600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster07
        */
        Id:       2612,
        Password: "14181608",
        Name:     "蘑菇人", // "Mushroom Man"  "きのこマン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Plant, // 植物
        Attack:   800,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster07
        */
        Id:       2613,
        Password: "02906250",
        Name:     "捕猎蛇", // "Grappler"  "グラップラー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Reptile, // 爬虫类
        Attack:   1300,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster07
        */
        Id:       2614,
        Password: "78402798",
        Name:     "蜥蜴骑士", // "Night Lizard"  "ナイト·リザード"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1150,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Booster07
        */
        Id:       2615,
        Password: "45909477",
        Name:     "月之使者", // "Moon Envoy"  "月の使者"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1100,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，金碎USR
        LB，Starter Box，Booster R1，TP03
        */
        Id:       2619,
        Password: "37313348",
        Name:     "龟虎", // "Turtle Tiger"  "タートル·タイガー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1000,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Starter Box
        */
        Id:       2620,
        Password: "80600490",
        Name:     "暗影战士", // "Kageningen"  "シャドウ·ファイター"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   800,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，Starter Box，TP13
        */
        Id:       2621,
        Password: "36121917",
        Name:     "怪兽蛋", // "Monster Egg"  "モンスター·エッグ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   600,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Starter Box
        */
        Id:       2622,
        Password: "77581312",
        Name:     "假面道化", // "Masked Clown"  "仮面道化"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   500,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Starter Box
        */
        Id:       2623,
        Password: "04931562",
        Name:     "山岳斗士", // "Mountain Warrior"  "マウンテン·ウォーリアー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_BeastWarror, // 兽战士
        Attack:   600,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Starter Box
        */
        Id:       2624,
        Password: "61454890",
        Name:     "时间魔人 独眼死灵师", // "Necrolancer the Timelord"  "時の魔人 ネクロランサ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   800,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Starter Box
        */
        Id:       2625,
        Password: "49370026",
        Name:     "戏法师", // "Genin"  "ジャグラー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   600,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Starter Box
        */
        Id:       2626,
        Password: "62210247",
        Name:     "月之女神 艾露塞姆", // "Lunar Queen Elzaim"  "月の女神 エルザェム"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   750,
        Defense:  1100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Starter Box，TP14，TU04
        */
        Id:       2627,
        Password: "01929294",
        Name:     "心钥妖精", // "Key Mace"  "キーメイス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   400,
        Defense:  300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Starter Box
        */
        Id:       2628,
        Password: "89558090",
        Name:     "暗黑囚犯", // "Dark Prisoner"  "ダーク·プリズナー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   600,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Starter Box
        */
        Id:       2629,
        Password: "43905751",
        Name:     "魔人铳", // "Madjinn Gunn"  "魔人銃"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   600,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，EX-R(EX)，VOL05，DL02
        */
        Id:       263,
        Password: "21263083",
        Name:     "苍白兽", // "Pale Beast"  "ペイルビースト"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1500,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Starter Box
        */
        Id:       2630,
        Password: "62793020",
        Name:     "谜之手", // "Mystery Hand"  "なぞの手"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   500,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Starter Box
        */
        Id:       2631,
        Password: "40575313",
        Name:     "地狱亡灵", // "Shadow Specter"  "ヘルバウンド"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   500,
        Defense:  200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，Starter Box，TP15
        */
        Id:       2632,
        Password: "83464209",
        Name:     "催眠羊", // "Mystical Sheep #2"  "スリーピィ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   800,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Starter Box
        */
        Id:       2633,
        Password: "86421986",
        Name:     "悟道的老树", // "Ancient Tree of Enlightenment"  "悟りの老樹"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Plant, // 植物
        Attack:   600,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，Starter Box
        */
        Id:       2634,
        Password: "22910685",
        Name:     "绿树之灵王", // "Green Phantom King"  "緑樹の霊王"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Plant, // 植物
        Attack:   500,
        Defense:  1600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        （这张卡在规则上不当作「恶魔/デーモン」卡使用）
        无限制
        平卡N
        Starter Box
        */
        Id:       2635,
        Password: "98075147",
        Name:     "恶魔蜗牛", // "Spiked Snail"  "デビルツムリ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   700,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        （这张卡在规则上不当作「恶魔/デーモン」卡使用）
        无限制
        平卡N
        Starter Box
        */
        Id:       2636,
        Password: "82742611",
        Name:     "恶魔蛇", // "Serpent Marauder"  "デビル·スネーク"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Reptile, // 爬虫类
        Attack:   700,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，Starter Box
        */
        Id:       2637,
        Password: "39004808",
        Name:     "根潭鱼人", // "Root Water"  "ルート·ウォーター"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Fish, // 鱼
        Attack:   900,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Starter Box
        */
        Id:       2638,
        Password: "15820147",
        Name:     "怪龟", // "Monsturtle"  "モンスタートル"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   800,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Starter Box
        */
        Id:       2639,
        Password: "18914778",
        Name:     "变形史莱姆", // "Change Slime"  "チェンジ·スライム"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   400,
        Defense:  300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，Starter Box，TP13
        */
        Id:       2640,
        Password: "96851799",
        Name:     "史汀", // "Hinotama Soul"  "スティング"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Pyro, // 炎
        Attack:   600,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LB，Starter Box
        */
        Id:       2641,
        Password: "40826495",
        Name:     "岩浆人", // "Dissolverock"  "マグマン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   900,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        Starter Box
        */
        Id:       2642,
        Password: "62193699",
        Name:     "大地战士", // "Rock Ogre Grotto #2"  "ウォー·アース"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   700,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EX-R(EX)，TP21，ST13
        */
        Id:       2648,
        Password: "50930991",
        Name:     "魔法剑士 尼奥", // "Neo the Magic Swordsman"  "魔法剣士ネオ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1700,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，金字UR
        EX-R(EX)，YAP01，TP16
        */
        Id:       2649,
        Password: "31122090",
        Name:     "逆转之女神", // "Gyakutenno Megami"  "逆転の女神"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   1800,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，VOL04，DL02
        */
        Id:       265,
        Password: "43230671",
        Name:     "古代的蜥蜴战士", // "Ancient Lizard Warrior"  "古代のトカゲ戦士"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Reptile, // 爬虫类
        Attack:   1400,
        Defense:  1100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EX-R(EX)
        */
        Id:       2650,
        Password: "86325596",
        Name:     "邪剑男爵", // "Baron of the Fiend Sword"  "邪剣男爵"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1550,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EX-R(EX)，DPKB
        */
        Id:       2651,
        Password: "30113682",
        Name:     "审判者", // "Judge Man"  "ジャッジ·マン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   2200,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EX-R(EX)
        */
        Id:       2652,
        Password: "49218300",
        Name:     "死亡巫师", // "Sorcerer of the Doomed"  "デス·ソーサラー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1450,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EX-R(EX)
        */
        Id:       2653,
        Password: "13723605",
        Name:     "食人宝石箱", // "Man-Eating Treasure Chest"  "人喰い宝石箱"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1600,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EX-R(EX)，DPKB
        */
        Id:       2654,
        Password: "50005633",
        Name:     "复仇的潜行剑士", // "Swordstalker"  "復讐のソード·ストーカー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   2000,
        Defense:  1600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        MA，YSD02
        */
        Id:       2655,
        Password: "44203504",
        Name:     "机械军曹", // "Robotic Knight"  "機械軍曹"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1600,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        TAEV(505)，DE02
        */
        Id:       2657,
        Password: "97127906",
        Name:     "外星人士兵", // "Alien Shocktrooper"  "エーリアン·ソルジャー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Reptile, // 爬虫类
        Attack:   1900,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        TAEV(505)
        */
        Id:       2658,
        Password: "33112041",
        Name:     "火山鼠", // "Volcanic Rat"  "ヴォルカニック·ラット"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Pyro, // 炎
        Attack:   500,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        TAEV(505)，YSD03，DE02
        */
        Id:       2659,
        Password: "60606759",
        Name:     "暗黑界的番兵 兰治", // "Renge, Gatekeeper of Dark World"  "暗黒界の番兵 レンジ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   100,
        Defense:  2100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        [14/04/27]
此卡是通常怪物。
        无限制
        银字R，平卡N
        TAEV(505)，DE02，ST14
        */
        Id:       2660,
        Password: "96005454",
        Name:     "猎龙", // "Hunter Dragon"  "ハウンド·ドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   1700,
        Defense:  100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        TAEV(505)
        */
        Id:       2661,
        Password: "22499463",
        Name:     "蛇毒眼镜蛇", // "Venom Cobra"  "ヴェノム·コブラ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Reptile, // 爬虫类
        Attack:   100,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        RB，VOL05，DL02
        */
        Id:       267,
        Password: "02118022",
        Name:     "兵主部", // "Hyosube"  "ひょうすべ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1500,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        MR
        */
        Id:       2739,
        Password: "73081602",
        Name:     "女王鸟", // "Queen Bird"  "クイーン·バード"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   1200,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        MR
        */
        Id:       2740,
        Password: "20624263",
        Name:     "孔雀", // "Peacock"  "クジャック"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   1700,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        MR
        */
        Id:       2741,
        Password: "46534755",
        Name:     "火焰乌贼", // "Fire Kraken"  "ファイヤー·クラーケン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_None, // 水
        Attack:   1600,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        MR
        */
        Id:       2742,
        Password: "56789759",
        Name:     "红龙", // "Tyhone #2"  "レッド·ドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   1700,
        Defense:  1900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        MR
        */
        Id:       2743,
        Password: "14015067",
        Name:     "大森林的长老", // "Ancient One of the Deep Forest"  "深き森の長老"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1800,
        Defense:  1900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        MR
        */
        Id:       2744,
        Password: "35565537",
        Name:     "女武神", // "Dark Witch"  "ヴァルキリー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   1800,
        Defense:  1700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        MR
        */
        Id:       2745,
        Password: "34442949",
        Name:     "机械蜗牛", // "Mechanical Snail"  "メカニカルスネイル"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   800,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        MR
        */
        Id:       2746,
        Password: "96981563",
        Name:     "食火的大龟", // "Giant Turtle Who Feeds on Flames"  "炎を食らう大亀"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1400,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        MR
        */
        Id:       2747,
        Password: "93108297",
        Name:     "液体兽", // "Liquid Beast"  "リクイド·ビースト"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   950,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        MR
        */
        Id:       2748,
        Password: "91996584",
        Name:     "鞭尾石像怪", // "Whiptail Crow"  "ウィップテイル·ガーゴイル"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1650,
        Defense:  1600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PS
        */
        Id:       2749,
        Password: "74637266",
        Name:     "章鱼怪", // "Octoberser"  "オクトバーサー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1600,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PS
        */
        Id:       2750,
        Password: "07892180",
        Name:     "念力河童", // "Psychic Kappa"  "サイコ·カッパー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   400,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PS
        */
        Id:       2751,
        Password: "29692206",
        Name:     "双尾", // "Twin Long Rods #2"  "ツインテール"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   850,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PS
        */
        Id:       2752,
        Password: "15023985",
        Name:     "岩石巨人", // "Stone Ogre Grotto"  "ストーンジャイアント"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   1600,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        CA
        */
        Id:       2753,
        Password: "42599677",
        Name:     "火焰杀手", // "Flame Champion"  "フレイムキラー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Pyro, // 炎
        Attack:   1900,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        CA
        */
        Id:       2754,
        Password: "78984772",
        Name:     "爆炎龙", // "Twin-Headed Fire Dragon"  "ビッグバンドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Pyro, // 炎
        Attack:   2200,
        Defense:  1700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        CA
        */
        Id:       2755,
        Password: "05388481",
        Name:     "燃烧士兵", // "Darkfire Soldier #1"  "バーニングソルジャー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Pyro, // 炎
        Attack:   1700,
        Defense:  1150,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        （这张卡在规则上不当作「火山/ヴォルカニック」卡使用）
        无限制
        平卡N
        CA
        */
        Id:       2756,
        Password: "31477025",
        Name:     "火山先生", // "Mr. Volcano"  "ミスターボルケーノ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Pyro, // 炎
        Attack:   2100,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        TB
        */
        Id:       2757,
        Password: "67049542",
        Name:     "暗黑蝙蝠", // "Dark Bat"  "ダークバット"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   1000,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        TB
        */
        Id:       2758,
        Password: "66927994",
        Name:     "鬼坦克T-34", // "Oni Tank T-34"  "鬼タンクT－３４"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1400,
        Defense:  1700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        TB
        */
        Id:       2759,
        Password: "02311603",
        Name:     "机关炮装甲车", // "Overdrive"  "ガトリングバギー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1600,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        TB
        */
        Id:       2760,
        Password: "87511987",
        Name:     "尖钉头", // "Spikebot"  "スパイク·ヘッド"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1800,
        Defense:  1700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        TB
        */
        Id:       2761,
        Password: "84620194",
        Name:     "断头台锹甲", // "Girochin Kuwagata"  "ギロチン·クワガタ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   1700,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        TB，TP17
        */
        Id:       2762,
        Password: "05265750",
        Name:     "海贼船 血骸骨号", // "Skull Mariner"  "海賊船スカルブラッド号"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1600,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        TB
        */
        Id:       2763,
        Password: "32269855",
        Name:     "独眼白虎", // "The All-Seeing White Tiger"  "隻眼のホワイトタイガー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1300,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        TB
        */
        Id:       2764,
        Password: "04042268",
        Name:     "岛龟", // "Island Turtle"  "島亀"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1100,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        TB
        */
        Id:       2765,
        Password: "31447217",
        Name:     "织翼者", // "Wingweaver"  "翼を織りなす者"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     7,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   2750,
        Defense:  2400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        TB
        */
        Id:       2766,
        Password: "67532912",
        Name:     "科学特殊兵", // "Science Soldier"  "科学特殊兵"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   800,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        TB
        */
        Id:       2767,
        Password: "04920010",
        Name:     "怨念集合体", // "Souls of the Forgotten"  "怨念集合体"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   900,
        Defense:  200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        TB
        */
        Id:       2768,
        Password: "30325729",
        Name:     "死神回力镖", // "Dokuroyaiba"  "死神ブーメラン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1000,
        Defense:  400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        SM，TP17
        */
        Id:       2769,
        Password: "31987274",
        Name:     "飞鱼", // "Flying Fish"  "フライング·フィッシュ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Fish, // 鱼
        Attack:   800,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        SM，KA，SK2
        */
        Id:       2770,
        Password: "86281779",
        Name:     "齿轮战士", // "Gadget Soldier"  "ガジェット·ソルジャー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1800,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        SM
        */
        Id:       2771,
        Password: "58818411",
        Name:     "女帝螳螂", // "Empress Mantis"  "女帝カマキリ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   2200,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        SM
        */
        Id:       2772,
        Password: "58696829",
        Name:     "生化僧侣", // "Bio-Mage"  "バイオ僧侶"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   1150,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LN
        */
        Id:       2773,
        Password: "98456117",
        Name:     "骨海马", // "Boneheimer"  "ボーンハイマー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   850,
        Defense:  400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LN
        */
        Id:       2774,
        Password: "12883044",
        Name:     "火焰舞者", // "Flame Dancer"  "フレイムダンサー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Pyro, // 炎
        Attack:   550,
        Defense:  450,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LN
        */
        Id:       2775,
        Password: "52121290",
        Name:     "蛇女郎", // "Spherous Lady"  "スフィラスレディ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   400,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        LN
        */
        Id:       2776,
        Password: "27671321",
        Name:     "雷鳗", // "Lightning Conger"  "雷ウナギ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Thunder, // 雷
        Attack:   350,
        Defense:  750,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        SC，YSD04
        */
        Id:       2779,
        Password: "93346024",
        Name:     "洞穴的潜龙", // "The Dragon Dwelling in the Cave"  "洞窟に潜む竜"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   1300,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        SC
        */
        Id:       2780,
        Password: "20831168",
        Name:     "蜥蜴兵", // "Lizard Soldier"  "リザード兵"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   1100,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        MA
        */
        Id:       2781,
        Password: "56369281",
        Name:     "贝奥武夫狼", // "Wolf Axwielder"  "ベイオウルフ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_BeastWarror, // 兽战士
        Attack:   1650,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        MA
        */
        Id:       2782,
        Password: "83764996",
        Name:     "魔导绅士-J", // "The Illusory Gentleman"  "魔導紳士－J"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1500,
        Defense:  1600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        MA
        */
        Id:       2783,
        Password: "92421852",
        Name:     "稀有金属女郎", // "Robolady"  "レアメタル·レディ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   450,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        MA
        */
        Id:       2784,
        Password: "38916461",
        Name:     "稀有金属战士", // "Roboyarou"  "レアメタル·ソルジャー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   900,
        Defense:  450,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        （这张卡在规则上不当作「熔岩/ラヴァル」卡使用）
        无限制
        平卡N
        PH
        */
        Id:       2787,
        Password: "17192817",
        Name:     "熔岩大巨人", // "Molten Behemoth"  "溶岩大巨人"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Pyro, // 炎
        Attack:   1000,
        Defense:  2200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PH
        */
        Id:       2788,
        Password: "04035199",
        Name:     "形体抓取者", // "Shapesnatch"  "シェイプ·スナッチ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1200,
        Defense:  1700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PH
        */
        Id:       2789,
        Password: "31242786",
        Name:     "食魂鱼", // "Souleater"  "魂喰らい"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Fish, // 鱼
        Attack:   1200,
        Defense:  0,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，金碎USR
        WCS2003，其他
        */
        Id:       2819,
        Password: "12829151",
        Name:     "女剑士 迦南", // "Kanan the Swordmistress"  "女剣士カナン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1400,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        [14/04/27]
此卡是通常怪物。
        无限制
        平卡N
        GLAS(506)，LCGX，DE02，SDWA，ST13，ST14
        */
        Id:       2830,
        Password: "44430454",
        Name:     "六武众的侍从", // "Chamberlain of the Six Samurai"  "六武衆の侍従"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   200,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        GLAS(506)，TP22
        */
        Id:       2831,
        Password: "80825553",
        Name:     "云魔物-小烟球", // "Cloudian - Smoke Ball"  "雲魔物－スモークボール"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Devine, // 天使
        Attack:   200,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PTDN(507)，LCGX，DE02
        */
        Id:       2969,
        Password: "90582719",
        Name:     "剑斗兽 独眼斗", // "Gladiator Beast Andal"  "剣闘獣アンダル"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_BeastWarror, // 兽战士
        Attack:   1900,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PTDN(507)，YSD03，DE02，SD23
        */
        Id:       2970,
        Password: "26976414",
        Name:     "海皇的长枪兵", // "Atlantean Pikeman"  "海皇の長槍兵"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Seaserpent, // 海龙
        Attack:   1400,
        Defense:  0,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE01，TB ，DL03，PE
        */
        Id:       301,
        Password: "27125110",
        Name:     "千眼邪教神", // "Thousand-Eyes Idol"  "千眼の邪教神"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   0,
        Defense:  0,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平罕NR，平卡N
        BE01，TB ，DL03
        */
        Id:       306,
        Password: "10992251",
        Name:     "超时空战斗机 V形蛇", // "Gradius"  "超時空戦闘機ビック·バイパー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1200,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，面闪SR
        BE01，SM，DL03，JY，SJ2
        */
        Id:       313,
        Password: "03573512",
        Name:     "地星剑士", // "Swordsman of Landstar"  "ランドスターの剣士"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   500,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE01，SM，DL03，SDM
        */
        Id:       314,
        Password: "46821314",
        Name:     "人形史莱姆", // "Humanoid Slime"  "ヒューマノイド·スライム"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   800,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        （这张卡在规则上也当作「异虫」卡使用）
        无限制
        平卡N
        BE01，SM，DL03，SDM
        */
        Id:       315,
        Password: "73216412",
        Name:     "鳞虫", // "Worm Drake"  "ワームドレイク"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Reptile, // 爬虫类
        Attack:   1400,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        YSD03，YSD05，AT02
        */
        Id:       3162,
        Password: "74093656",
        Name:     "调整战士", // "Tune Warrior"  "チューン·ウォリアー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1600,
        Defense:  200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        YSD03，YSD05
        */
        Id:       3163,
        Password: "00487395",
        Name:     "冰水精灵", // "Water Spirit"  "ウォーター·スピリット"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   400,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE01，SM，DL03
        */
        Id:       318,
        Password: "67371383",
        Name:     "半鱼兽", // "Amphibian Beast"  "半魚獣·フィッシャービースト"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Fish, // 鱼
        Attack:   2400,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，面闪SR
        DT01，HA01，SD25
        */
        Id:       3187,
        Password: "21615956",
        Name:     "炎狱护卫龙", // "Flamvell Guard"  "ガード·オブ·フレムベル"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   100,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE01，SM，DL03
        */
        Id:       319,
        Password: "87303357",
        Name:     "光辉深渊", // "Shining Abyss"  "シャイン·アビス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   1600,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，面闪SR
        DT01，HA01
        */
        Id:       3192,
        Password: "89386122",
        Name:     "正义盟军 辉剑鸟", // "Ally of Justice Clausolas"  "A·O·J クラウソラス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   2300,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE01，SM，DL03
        */
        Id:       320,
        Password: "13676474",
        Name:     "假面咒术师 诅咒之喉", // "Grand Tiki Elder"  "仮面呪術師カースド·ギュラ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1500,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE01，SM，DL03
        */
        Id:       322,
        Password: "86569121",
        Name:     "梅尔基多四面兽", // "Melchid the Four-Face Beast"  "メルキド四面獣"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1500,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        银字R，平卡N
        DT02，TSHD(608)
        */
        Id:       3315,
        Password: "68505803",
        Name:     "次世代控制员", // "Genex Controller"  "ジェネクス·コントローラー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1400,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        DT02，GLD03
        */
        Id:       3323,
        Password: "29054481",
        Name:     "霞之谷的看守者", // "Mist Valley Watcher"  "霞の谷の見張り番"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1500,
        Defense:  1900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        DT02，YSD04，DB12
        */
        Id:       3329,
        Password: "23115241",
        Name:     "X-剑士 安娜佩瑞拉", // "X-Saber Anu Piranha"  "X－セイバー アナペレラ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1800,
        Defense:  1100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平罕NR，面闪SR
        DT03，HA02
        */
        Id:       3461,
        Password: "40155554",
        Name:     "盟军·头脑", // "Ally Mind"  "A·マインド"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1800,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE01，LN，DL03
        */
        Id:       360,
        Password: "32541773",
        Name:     "绘画中的潜藏者", // "The Portrait's Secret"  "絵画に潜む者"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1200,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE01，LN，DL03
        */
        Id:       361,
        Password: "68049471",
        Name:     "梦魔的亡灵", // "The Gross Ghost of Fled Dreams"  "夢魔の亡霊"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1300,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE01，LN，DL03
        */
        Id:       362,
        Password: "05434080",
        Name:     "无头骑士", // "Headless Knight"  "首なし騎士"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1450,
        Defense:  1700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE01，LN，DL03
        */
        Id:       363,
        Password: "67105242",
        Name:     "地缚灵", // "Earthbound Spirit"  "地縛霊"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   500,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE01，LN，DL03
        */
        Id:       364,
        Password: "66989694",
        Name:     "死灵伯爵", // "The Earl of Demise"  "死霊伯爵"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   2000,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        ABPF(607)
        */
        Id:       4043,
        Password: "57308711",
        Name:     "单轮车人", // "Unicycular"  "アンサイクラー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   100,
        Defense:  100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        ME，BE02，VOL06，DL04，JY，SJ2
        */
        Id:       408,
        Password: "88819587",
        Name:     "宝贝龙", // "Baby Dragon"  "ベビードラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   1200,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        ME，BE02，VOL06，DL04
        */
        Id:       409,
        Password: "87564352",
        Name:     "暗黑之龙王", // "Blackland Fire Dragon"  "暗黒の竜王"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   1500,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        ME，BE02，VOL06，DL04
        */
        Id:       411,
        Password: "18246479",
        Name:     "牛魔人", // "Battle Steer"  "牛魔人"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_BeastWarror, // 兽战士
        Attack:   1800,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        ME，BE02，VOL07，DL04，DPKB，KA，SK2
        */
        Id:       413,
        Password: "66602787",
        Name:     "暗道化师 扎奇", // "Saggi the Dark Clown"  "闇·道化師のサギー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   600,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        ME，BE02，VOL07，DL04，PE
        */
        Id:       415,
        Password: "28546905",
        Name:     "无脸幻想师", // "Illusionist Faceless Mage"  "幻想師·ノー·フェイス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1200,
        Defense:  2200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE01，PS，DL01
        */
        Id:       42,
        Password: "67284908",
        Name:     "迷宫壁", // "Labyrinth Wall"  "迷宮壁－ラビリンス·ウォール－"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   0,
        Defense:  3000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        ME，DL04，Booster07
        */
        Id:       420,
        Password: "14851496",
        Name:     "水母", // "Jellyfish"  "海月－ジェリーフィッシュ－"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1200,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        ME，BE02，DL04，Booster07
        */
        Id:       422,
        Password: "69455834",
        Name:     "暗魔界的霸王", // "King of Yamimakai"  "闇魔界の覇王"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   2000,
        Defense:  1530,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        ME，BE02，EX-R(EX)，VOL07，DL04
        */
        Id:       424,
        Password: "68516705",
        Name:     "人马兽", // "Mystic Horseman"  "ケンタウロス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1300,
        Defense:  1550,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        ME，BE02，EX-R(EX)，DL04，Booster07，KA，SK2
        */
        Id:       435,
        Password: "24611934",
        Name:     "强化石像怪", // "Ryu-Kishin Powered"  "ガーゴイル·パワード"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1600,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，金字UR
        ME，BE02，LE02，VOL07，DL04
        */
        Id:       436,
        Password: "87322377",
        Name:     "TM-1火箭炮蜘蛛", // "Launcher Spider"  "TM－１ランチャースパイダー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     7,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Machine, // 机械
        Attack:   2200,
        Defense:  2500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，面闪SR
        DT09，HA05，SPRG
        */
        Id:       4368,
        Password: "91731841",
        Name:     "宝石骑士·红榴", // "Gem-Knight Garnet"  "ジェムナイト·ガネット"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Pyro, // 炎
        Attack:   1900,
        Defense:  0,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，面闪SR
        DT09，HA05，SPRG
        */
        Id:       4369,
        Password: "27126980",
        Name:     "宝石骑士·青玉", // "Gem-Knight Sapphire"  "ジェムナイト·サフィア"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_None, // 水
        Attack:   0,
        Defense:  2100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        ME，VOL06，DL04，TP11
        */
        Id:       437,
        Password: "08471389",
        Name:     "高科技狼", // "Giga-Tech Wolf"  "ギガテック·ウルフ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1200,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，面闪SR
        DT09，HA05，SPRG
        */
        Id:       4370,
        Password: "54620698",
        Name:     "宝石骑士·黄碧", // "Gem-Knight Tourmaline"  "ジェムナイト·ルマリン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Thunder, // 雷
        Attack:   1600,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        ME，VOL07，DL04，SD04，GLD01
        */
        Id:       439,
        Password: "23771716",
        Name:     "彩虹鱼", // "7 Colored Fish"  "レインボー·フィッシュ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Fish, // 鱼
        Attack:   1800,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        ME，VOL06，DL04，TP10
        */
        Id:       457,
        Password: "17358176",
        Name:     "高等女祭司", // "Lady of Faith"  "ハイ·プリーステス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1100,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，金字UR
        ME，BE02，LE02，YSD01，DL04，YU，SY2，15AY
        */
        Id:       460,
        Password: "05818798",
        Name:     "幻兽王 加泽尔", // "Gazelle the King of Mythical Beasts"  "幻獣王ガゼル"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1500,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，金字UR
        ME，DL04，VJ
        */
        Id:       461,
        Password: "49888191",
        Name:     "甘尼许巨象", // "Garnecia Elefantis"  "ガーネシア·エレファンティス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     7,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_BeastWarror, // 兽战士
        Attack:   2400,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        银字R，面闪SR
        DT11，HA06，SPRG
        */
        Id:       4666,
        Password: "76908448",
        Name:     "宝石骑士·白晶", // "Gem-Knight Crystal"  "ジェムナイト·クリスタ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     7,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   2450,
        Defense:  1950,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE02，DL04，Booster01，Booster Chronicle，Booster R1，TU04
        */
        Id:       479,
        Password: "27324313",
        Name:     "电气小子", // "Wattkid"  "エレキッズ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Thunder, // 雷
        Attack:   1000,
        Defense:  500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        DL04，Booster02，Booster Chronicle，Booster R1
        */
        Id:       480,
        Password: "41762634",
        Name:     "吸血跳蚤", // "Giant Flea"  "吸血ノミ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Insect, // 昆虫
        Attack:   1500,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE02，DL04，Booster02，Booster Chronicle，Booster R1
        */
        Id:       485,
        Password: "10262698",
        Name:     "复活节岛的摩艾石像", // "The Statue of Easter Island"  "イースター島のモアイ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   1100,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE02，DL04，Booster02，Booster Chronicle，Booster R1
        */
        Id:       486,
        Password: "82085619",
        Name:     "友谊天使", // "Shining Friendship"  "フレンドシップ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Devine, // 天使
        Attack:   1300,
        Defense:  1100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，金字UR
        BE02，LE02，DL04，Booster03，Booster Chronicle，Booster R2，JY
        */
        Id:       492,
        Password: "49791927",
        Name:     "老虎斧战士", // "Tiger Axe"  "タイガー·アックス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_BeastWarror, // 兽战士
        Attack:   1300,
        Defense:  1100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE02，DL04，Booster04，Booster Chronicle，Booster R2，YSD03，JY，SJ2
        */
        Id:       493,
        Password: "48305365",
        Name:     "巨斧袭击者", // "Axe Raider"  "アックス·レイダー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1700,
        Defense:  1150,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE02，DL04，SD10，Booster04，Booster Chronicle，Booster R2
        */
        Id:       495,
        Password: "07359741",
        Name:     "机械猎手", // "Mechanicalchaser"  "メカ·ハンター"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1850,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        DL04，Booster04，Booster Chronicle，Booster R2
        */
        Id:       496,
        Password: "58831685",
        Name:     "海鳞蛇", // "Giant Red Seasnake"  "シーザリオン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1800,
        Defense:  800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE02，DL04，SD06，Booster04，Booster Chronicle，Booster R2
        */
        Id:       498,
        Password: "69140098",
        Name:     "双生精灵", // "Gemini Elf"  "ヂェミナイ·エルフ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1900,
        Defense:  900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PHSW(706)
        */
        Id:       4982,
        Password: "69380702",
        Name:     "小白兔", // "Bunilla"  "バニーラ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   150,
        Defense:  2050,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        银字R，平卡N
        PHSW(706)，SD25
        */
        Id:       4983,
        Password: "95788410",
        Name:     "兔龙", // "Rabidragon"  "ラビードラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     8,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   2950,
        Defense:  2900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        DL04，Booster03，Booster Chronicle，Booster R2
        */
        Id:       500,
        Password: "03170832",
        Name:     "橐蜚", // "Takuhee"  "タクヒ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   1450,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        DL04，Booster04，Booster Chronicle，Booster R2
        */
        Id:       502,
        Password: "81563416",
        Name:     "泉之妖精", // "Fairy of the Fountain"  "泉の妖精"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   1600,
        Defense:  1100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        DL04，Booster04，Booster Chronicle，Booster R2
        */
        Id:       503,
        Password: "79629370",
        Name:     "月光少女", // "Maiden of the Moonlight"  "月明かりの乙女"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   1500,
        Defense:  1300,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N，金字UR，爆闪PR
        BE02，DL04，Booster05，Booster Chronicle，Booster R3，TP19
        */
        Id:       506,
        Password: "94119974",
        Name:     "双头恐龙王", // "Two-Headed King Rex"  "二頭を持つキング·レックス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Dinosaur, // 恐龙
        Attack:   1600,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        DL04，Booster06，Booster Chronicle，Booster R3
        */
        Id:       507,
        Password: "76634149",
        Name:     "海龙神", // "Kairyu-Shin"  "海竜神"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Seaserpent, // 海龙
        Attack:   1800,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EX-R(EX)，DL04，Booster05，Booster Chronicle，Booster R3
        */
        Id:       508,
        Password: "66672569",
        Name:     "龙僵尸", // "Dragon Zombie"  "ドラゴン·ゾンビ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   1600,
        Defense:  0,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE02，DL04，Booster05，Booster Chronicle，Booster R3
        */
        Id:       510,
        Password: "99510761",
        Name:     "神灯魔人", // "Lord of the Lamp"  "ランプの魔人"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1400,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        DT13，DS13
        */
        Id:       5106,
        Password: "77542832",
        Name:     "入魔鬼·血石", // "Evilswarm Heliotrope"  "ヴェルズ·ヘリオロープ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   1950,
        Defense:  650,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        面闪SR，金字UR，平卡N
        PHSW(706)，ST12，EP12，SD25
        */
        Id:       5126,
        Password: "43096270",
        Name:     "紫翠玉龙", // "Alexandrite Dragon"  "アレキサンドライドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   2000,
        Defense:  100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        [14/04/27]
此卡是通常怪物。
        无限制
        平卡N
        GAOV(708)，ST14
        */
        Id:       5300,
        Password: "87151205",
        Name:     "电气尾龙", // "Wattaildragon"  "エレキテルドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   2500,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        GAOV(708)
        */
        Id:       5301,
        Password: "13140300",
        Name:     "神龙之圣刻印", // "Hieratic Seal of the Sun Dragon Overlord"  "神龍の聖刻印"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     8,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   0,
        Defense:  0,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        银字R，平卡N
        REDU(801)，ST13
        */
        Id:       5388,
        Password: "14214060",
        Name:     "魔法剑士 特兰斯", // "Trance the Magic Swordsman"  "魔法剣士トランス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   2600,
        Defense:  200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        面闪SR，金字UR，银字R
        GAOV(708)，EP13
        */
        Id:       5470,
        Password: "92125819",
        Name:     "圣骑士 阿托利斯", // "Noble Knight Artorigus"  "聖騎士アルトリウス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1800,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        PS，DL01，PE
        */
        Id:       55,
        Password: "02964201",
        Name:     "蛋龙", // "Ryu-Ran"  "ドラゴン·エッガー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     7,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   2200,
        Defense:  2600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE02，SC，DL05
        */
        Id:       551,
        Password: "14531242",
        Name:     "红色独眼巨人", // "Opticlops"  "レッド·サイクロプス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1800,
        Defense:  1700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        [14/04/27]
此卡是通常怪物。
        无限制
        平卡N
        BE02，SC，YSD01，DL05，ST14
        */
        Id:       562,
        Password: "75953262",
        Name:     "战士 戴·格雷法", // "Warrior Dai Grepher"  "戦士ダイ·グレファー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1700,
        Defense:  1600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        银字R，平卡N
        BE02，SC，YSD01，DL05，SK2，YSD06
        */
        Id:       589,
        Password: "17658803",
        Name:     "绿宝石龙", // "Luster Dragon #2"  "エメラルド·ドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   2400,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        [拉长石龙]
<ラブラドライドラゴン>
[14/07/19]

◇此卡是通常怪兽，是调整，不是效果怪兽
        无限制
        平卡N，面闪SR
        SHSP(806)
        */
        Id:       5994,
        Password: "62514770",
        Name:     "拉长石龙", // "Labradorite Dragon"  "ラブラドライドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   0,
        Defense:  2400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        BE01，PS，DL01
        */
        Id:       60,
        Password: "62397231",
        Name:     "钻石龙", // "Hyozanryu"  "ダイヤモンド·ドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     7,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   2100,
        Defense:  2800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        [银河星蛇]
<Galaxy Serpent>
[14/05/21]

◇此卡是通常怪兽、是调整
◇此卡视为「银河/ギャラクシー」怪兽
        无限制
        面闪SR，金字UR，平卡N
        JOTL(805)，EP14
        */
        Id:       6079,
        Password: "11066358",
        Name:     "银河海蛇", // "Galaxy Serpent"  "ギャラクシーサーペント"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   1000,
        Defense:  0,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平罕NR
        LVAL(807)
        */
        Id:       6131,
        Password: "03557275",
        Name:     "白尘妖", // "White Duston"  "ホワイト·ダストン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   0,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        [闪光之骑士]
<閃光の騎士>
[14/07/21]

◇此卡是没有怪兽效果的灵摆怪兽，是通常怪兽
◇此卡没有灵摆效果，灵摆刻度不是效果
        无限制
        银字R
        DUEA(901)
        */
        Id:       6376,
        Password: "17390179",
        Name:     "闪光之骑士", // "Flash Knight"  "閃光の騎士"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1800,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        [傅科魔炮石]
<フーコーの魔砲石>
[14/04/24]

◇此卡是没有怪物效果的灵摆怪兽，是通常怪兽
●①：这张卡发动的回合的结束阶段，以场上1张表侧表示的魔法·陷阱卡为对象才能发动。那张卡破坏。
◇魔法卡的效果，任意发动，开连锁，以场上表侧表示的1张魔法·陷阱卡为对象
◇可以取灵摆区域内的卡为对象，可以取此卡自身为对象
        无限制
        银字R
        DUEA(901)
        */
        Id:       6377,
        Password: "43785278",
        Name:     "傅科魔炮石", // "Foucault's Cannon"  "フーコーの魔砲石"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   2200,
        Defense:  1200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        [玄化武装龙]
<メタファイズ·アームド·ドラゴン>
[14/04/24]

◇此卡是通常怪兽
        无限制
        平卡N
        DUEA(901)
        */
        Id:       6378,
        Password: "89189982",
        Name:     "玄化武装龙", // "Metaphys Armed Dragon"  "メタファイズ·アームド·ドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     7,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_MagicDragon, // 幻龙
        Attack:   2800,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        [机壳工具 丑恶]
<クリフォート·ツール>
[14/07/19]

◎此卡是没有怪兽效果的灵摆怪兽，是通常怪兽
●①：自己不是「机壳/クリフォート」怪兽不能特殊召唤。这个效果不会被无效化。
◇魔法卡的效果，不开连锁，不取对象，这个效果无法无效，此卡的其他效果可以被无效
●②：1回合1次，支付800基本分才能发动。从卡组把「机壳工具 丑恶/クリフォート·ツール」以外的1张「机壳/クリフォート」卡加入手卡。
◇魔法卡的效果，开连锁，不取对象
◇在自己的主要阶段才能发动这个效果
◇支付800基本分（LP）是效果发动COST
        无限制
        银字R
        NECH(902)
        */
        Id:       6597,
        Password: "65518099",
        Name:     "机壳工具 丑恶", // "-"  "クリフォート·ツール"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1000,
        Defense:  2800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        ◇此卡是观赏卡，不能在决斗中使用
        观赏卡
        金字UR
        WCP2014
        */
        Id:       6668,
        Password: "-",
        Name:     "列奥纳多的白银飞船", // "Leonardo's Silver Skyship"  "Leonardo's Silver Skyship"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     7,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Machine, // 机械
        Attack:   0,
        Defense:  3000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR，面闪SR
        DUEA(901)
        */
        Id:       6704,
        Password: "21970285",
        Name:     "龙角猎人", // "Dragon Horn Hunter"  "-"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   2300,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        SECE(903)
        */
        Id:       6760,
        Password: "99645428",
        Name:     "宝石骑士·小琉", // "-"  "ジェムナイト・ラピス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Rock, // 岩石
        Attack:   1200,
        Defense:  100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        [机壳汇编器 不安定]
<クリフォート・アセンブラ>
[14/11/16]

◇不持有怪兽效果的灵摆怪兽，即通常怪兽
====
●①：自己不是「机壳」怪兽不能特殊召唤。这个效果不会被无效化。
◇不进入连锁
●②：自己上级召唤成功的回合的结束阶段才能发动。自己从卡组抽出这个回合自己为上级召唤而解放的「机壳」怪兽的数量。
◇进入连锁
◇回合结束时发动一次
◇解放后才发动这张卡也可以发动这个效果，会计算这张卡发动前解放的数目
◇解放装备有「機殻の生贄/机壳的牲祭」的1体怪兽上级召唤7星以上怪兽，只抽1张
◇上级召唤被无效不能发动
◇灵摆区有2张同名卡的场合，各自开连锁发动
◇里侧表示上级召唤也发动
◇解放得到「机壳」怪兽名字的「ファントムオブカオス/混沌幻影」也发动
◇解放里侧表示「机壳」怪兽也发动
◇解放原卡名是「机壳」，现卡名变成不带「机壳」的怪兽不能发动
        无限制
        银字R
        SECE(903)
        */
        Id:       6768,
        Password: "51194046",
        Name:     "机壳汇编器 不安定", // "-"  "クリフォート・アセンブラ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   2400,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE01，301，SD02
        */
        Id:       680,
        Password: "24530661",
        Name:     "达人僵尸", // "Master Kyonshee"  "達人キョンシー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Zombie, // 不死
        Attack:   1750,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE01，301，SD09
        */
        Id:       681,
        Password: "51934376",
        Name:     "打喷嚏的河马龙", // "Kabazauls"  "大くしゃみのカバザウルス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Dinosaur, // 恐龙
        Attack:   1700,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE01，301
        */
        Id:       682,
        Password: "97923414",
        Name:     "大木人18", // "Inpachi"  "大木人１８"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1600,
        Defense:  1900,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平爆NPR
        JF15
        */
        Id:       6850,
        Password: "83980492",
        Name:     "铜锣龙", // "-"  "銅鑼ドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   500,
        Defense:  2100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平爆NPR
        JF15
        */
        Id:       6851,
        Password: "19474136",
        Name:     "曼陀林草龙", // "-"  "マンドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Plant, // 植物
        Attack:   2500,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平罕NR
        CROS(904)
        */
        Id:       6891,
        Password: "74852097",
        Name:     "幻之狮鹫", // "-"  "幻のグリフォン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   2000,
        Defense:  0,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        [索菲娅之影灵衣]
<ｓｏｐｈｉａの影霊衣>
[2015/02/22]

●「影灵衣」仪式魔法卡降临。不是使用各自种族不同的自己场上3只怪兽作从手卡的仪式召唤不能特殊召唤。
◇只能从手牌进行仪式召唤
◇只能使用各自种族不同的自己场上3只怪兽
不能
·适用《影霊衣の術士 シュリット/影灵衣术士 施里特》
·使用仪式魔人的效果
·因《影霊衣の降魔鏡/影灵衣的降魔镜》将手牌解放或除外墓地
·因《クロス・ソウル/灵魂交错》使用对方怪兽
·因《影霊衣の反魂術/影灵衣的返魂术》使用墓地的这张卡仪式召唤
·因《高等儀式術/高等仪式术》使用卡组的通常怪兽
·因《影霊衣の万華鏡/影灵衣的万华镜》将额外卡组怪兽送去墓地
●自己·对方的主要阶段1从手卡把这张卡和1张「影灵衣」魔法卡丢弃才能发动。那次阶段内，对方不能从额外卡组把怪兽特殊召唤。
◇诱发即时效果
◇不取对象
◇丢弃是COST
◇连锁《融合/融合》发动的场合，其效果不适用
◇适用后不进入连锁的特殊召唤也不能进行
◇《マクロコスモス/大宇宙》适用中可以发动并正常适用
◇效果适用后这张卡不在墓地效果仍适用
●这张卡仪式召唤成功时才能发动。这张卡以外的双方的场上·墓地的卡全部除外。这个效果发动的回合，自己不能把其他怪兽通常召唤·特殊召唤。
◇诱发效果
◇不取对象
◇已经召唤·特殊召唤其他怪兽的回合不能发动
◇双方的场上·墓地有一个区域没有卡也可以发动
◇连锁《王宮の鉄壁/王宫的铁壁》的场合，除外不处理，不能召唤仍适用
◇发动被无效的场合，这回合可以召唤
◇仅仅效果被无效的场合，这回合不能召唤
◇《王家の眠る谷－ネクロバレー/王家长眠之谷》适用中，若双方墓地没有卡可以发动，不然不能发动
        无限制
        金闪UR，银碎SER，立体UTR
        CROS(904)
        */
        Id:       6928,
        Password: "21105106",
        Name:     "索菲娅之影灵衣", // "-"  "ｓｏｐｈｉａの影霊衣"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     11,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Spellcaster, // 魔法师
        Attack:   3600,
        Defense:  3400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        金字UR，爆闪PR，立体UTR，银字R，平卡N，面闪SR
        PG，BE01，PP05，VOL03，DL02，301，SD01，DT01，YAP01，DTP01，JY，SJ2，LC01，SD22
        */
        Id:       735,
        Password: "74677422",
        Name:     "真红眼黑龙", // "Red-Eyes B. Dragon"  "真紅眼の黒竜"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     7,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   2400,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE01，302，SY2
        */
        Id:       736,
        Password: "12143771",
        Name:     "流离的饥民", // "People Running About"  "逃げまどう民"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Pyro, // 炎
        Attack:   600,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE01，302，SY2
        */
        Id:       737,
        Password: "58538870",
        Name:     "受弹压的民众", // "Oppressed People"  "弾圧される民"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_None, // 水
        Attack:   400,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE01，302，SY2
        */
        Id:       738,
        Password: "85936485",
        Name:     "团结的反抗军", // "United Resistance"  "団結するレジスタンス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Thunder, // 雷
        Attack:   1000,
        Defense:  400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        面闪SR，平卡N，银字R
        DP02，EE01，302，DPKB，SK2
        */
        Id:       739,
        Password: "62651957",
        Name:     "X-首领加农", // "X-Head Cannon"  "X－ヘッド·キャノン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1800,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        [14/04/27]
此卡是通常怪物。　
        无限制
        面闪SR，平卡N
        YSD01，EE01，302，SD05，YSD03，SY2，YSD06，ST14
        */
        Id:       742,
        Password: "11321183",
        Name:     "暗魔界的战士 暗黑之剑", // "Dark Blade"  "闇魔界の戦士 ダークソード"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Warrior, // 战士
        Attack:   1800,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE01，303，SD03，SJ2
        */
        Id:       790,
        Password: "11813953",
        Name:     "大型肉畜", // "Great Angus"  "グレート·アンガス"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Beast, // 兽
        Attack:   1800,
        Defense:  600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE01，303
        */
        Id:       791,
        Password: "48202661",
        Name:     "彼者", // "Aitsu"  "アイツ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     5,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Devine, // 天使
        Attack:   100,
        Defense:  100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE01，303，SD08
        */
        Id:       792,
        Password: "84696266",
        Name:     "音速鸭", // "Sonic Duck"  "音速ダック"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_WindBeast, // 鸟兽
        Attack:   1700,
        Defense:  700,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        [14/04/27]
此卡是通常怪物。
        无限制
        平卡N
        YSD01，EE01，303，SD01，SK2，ST14
        */
        Id:       793,
        Password: "11091375",
        Name:     "蓝宝石龙", // "Luster Dragon"  "サファイアドラゴン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_Dragon, // 龙
        Attack:   1900,
        Defense:  1600,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE01，304，YSD05，SJ2
        */
        Id:       842,
        Password: "48094997",
        Name:     "战斗橄榄球员", // "Battle Footballer"  "バトルフットボーラー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_None, // 炎
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1000,
        Defense:  2100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        YSD02，EE01，304
        */
        Id:       843,
        Password: "11987744",
        Name:     "忍犬 奇狗", // "Nin-Ken Dog"  "忍犬ワンダードッグ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 风
        Lr:       ygo.LR_BeastWarror, // 兽战士
        Attack:   1800,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE01，304，SJ2
        */
        Id:       844,
        Password: "47372349",
        Name:     "杂技猴", // "Acrobat Monkey"  "アクロバットモンキー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     3,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1000,
        Defense:  1800,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        CA，DL01，KA，SK2
        */
        Id:       89,
        Password: "90908427",
        Name:     "铁腕巨人", // "Steel Ogre Grotto #2"  "鉄腕ゴーレム"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     6,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Machine, // 机械
        Attack:   1900,
        Defense:  2200,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE01，305
        */
        Id:       896,
        Password: "49003308",
        Name:     "魔蜥义豪", // "Gagagigo"  "ガガギゴ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 水
        Lr:       ygo.LR_Reptile, // 爬虫类
        Attack:   1850,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE01，305
        */
        Id:       897,
        Password: "86498013",
        Name:     "异次元的驯兽师", // "D.D. Trainer"  "異次元トレーナー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     1,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   100,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        DP02，EE01，305
        */
        Id:       898,
        Password: "12482652",
        Name:     "扰乱·绿", // "Ojama Green"  "おジャマ·グリーン"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Beast, // 兽
        Attack:   0,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE01，305
        */
        Id:       899,
        Password: "49881766",
        Name:     "恶魔士兵", // "Archfiend Soldier"  "デーモン·ソルジャー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1900,
        Defense:  1500,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        CA，DL01
        */
        Id:       90,
        Password: "78423643",
        Name:     "三头基多", // "Three-Headed Geedo"  "三ツ首のギドー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 暗
        Lr:       ygo.LR_Fiend, // 恶魔
        Attack:   1200,
        Defense:  1400,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        DP02，EE02，306，DT01，DTP01
        */
        Id:       949,
        Password: "42941100",
        Name:     "扰乱·黄", // "Ojama Yellow"  "おジャマ·イエロー"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Beast, // 兽
        Attack:   0,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        DP02，EE02，306
        */
        Id:       950,
        Password: "79335209",
        Name:     "扰乱·黑", // "Ojama Black"  "おジャマ·ブラック"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     2,
        La:       ygo.LA_Earth, // 光
        Lr:       ygo.LR_Beast, // 兽
        Attack:   0,
        Defense:  1000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE02，306
        */
        Id:       951,
        Password: "15734813",
        Name:     "魂虎", // "Soul Tiger"  "魂虎"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     4,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   0,
        Defense:  2100,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)


    co = &ygo.CardOriginal{
        /*  
        
        无限制
        平卡N
        EE02，306
        */
        Id:       952,
        Password: "42129512",
        Name:     "巨大树熊", // "Big Koala"  "ビッグ·コアラ"
        Describe: "",
        Lc:       ygo.LC_OrdinaryMonster, // 通常怪兽
        Star:     7,
        La:       ygo.LA_Earth, // 地
        Lr:       ygo.LR_Beast, // 兽
        Attack:   2700,
        Defense:  2000,
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)

}