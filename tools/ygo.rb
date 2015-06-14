#!/usr/bin/env ruby
#-*- coding:utf-8 -*-


require 'nokogiri'
require './puts'
require './num_cn'
require 'json'

class String
  @@Encodes = %w(gbk gb18030 gb2312 cp936 ucs-bom shift-jis ascii-8bit)
#   @@Encodes = Encoding.list
  def to_utf8!
    @@Encodes.each do |coding|
      begin
        self.force_encoding Encoding::UTF_8
        return self if self.valid_encoding?
        self.encode! Encoding::UTF_8, coding
      rescue
      end
    end
    self
  end
end
def files_list path, &block
  Dir.entries(path).each do |sub|   
    if sub != '.' && sub != '..'
      if File.directory?("#{path}/#{sub}")
        get_file_list("#{path}/#{sub}")
      else
        block.call "#{path}/#{sub}"
      end
    end
  end
end

allcard = {}




$las = {
    "地"=>"Earth",
    "水"=>"Earth",
    "火"=>"Earth",
    "风"=>"Earth",
    "光"=>"Earth",
    "暗"=>"Earth",
    "神"=>"Earth"
}

def la n
    "ygo.LA_" + if $las[n] != nil 
        $las[n]
    else
        "None"
    end
end

$lcs = {
    "通常怪兽"=>"OrdinaryMonster",
    "效果怪兽"=>"EffectMonster",
    "融合怪兽"=>"FusionMonster",
    "XYZ怪兽"=>"ExcessMonster",
    "同调怪兽"=>"HomologyMonster",
    "仪式怪兽"=>"RiteMonster",
    "通常魔法"=>"OrdinaryMagic",
    "仪式魔法"=>"RiteMagic",
    "永续魔法"=>"SustainsMagic",
    "装备魔法"=>"EquipMagic",
    "场地魔法"=>"PlaceMagic",
    "速攻魔法"=>"RushMagic",
    "通常陷阱"=>"OrdinaryTrap",
    "永续陷阱"=>"SustainsTrap",
    "反击陷阱"=>"ReactionTrap"
}
def lc n
    "ygo.LC_" + if $lcs[n] != nil 
        $lcs[n]
    else
        "None"
    end
end
$lrs = {
    "魔法师"=>"SpellCaster",
    "龙"=>"Dragon",
    "爬虫类"=>"Reptile",
    "兽"=>"Beast",
    "昆虫"=>"Insect",
    "海龙"=>"Seaserpent",
    "鱼"=>"Fish",
    "炎"=>"Pyro",
    "兽战士"=>"BeastWarror",
    "鸟兽"=>"WindBeast",
    "恐龙"=>"Dinosaur",
    "植物"=>"Plant",
    "机械"=>"Machine",
    "战士"=>"Warrior",
    "天使"=>"Devine",
    "不死"=>"Zombie",
    "恶魔"=>"Fiend",
    "岩石"=>"Rock",
    "雷"=>"Thunder",
    "幻神兽"=>"DivineBeast",
    "念动力"=>"Psycho",
    "创造神"=>"CreatorGod",
    "幻龙"=>"PhantomDragon"
}

def lr n
    "ygo.LR_" + if $lrs[n] != nil 
        $lrs[n]
    else
        "None"
    end
end

def OrdinaryMonster allcard
    print '
package cards

import (
    "ygo"
)


func original(cardBag *ygo.CardVersion) {
    var co *ygo.CardOriginal
    '
    allcard.each do|k,v|
        if v["卡片种类"] == "通常怪兽"
            print "
    co = &ygo.CardOriginal{
        /*#{_puts v}*/
        Id:       #{v["id"]},
        Password: \"#{v["卡片密码"]}\",
        Name:     \"#{v["中文名"]}\", // \"#{v["英文名"]}\"  \"#{v["日文名"]}\"
        Lc:       #{lc(v["卡片种类"])}, // #{v["卡片种类"]}
        Level:    #{v["星级"]},
        La:       #{la(v["属性"])}, // #{v["属性"]}
        Lr:       #{lr(v["种族"])}, // #{v["种族"]}
        Attack:   #{v["攻击力"]},
        Defense:  #{v["防御力"]},
        Summon:   ygo.SuperiorCall, // 召唤
        IsValid:  true,
    }
    cardBag.Register(co)

"
        end
    end
    print '}'
end


def elses allcard
    i = 0
    j = 0
    print '
package cards

import (
    "ygo"
)
'
    allcard.keys().each do|k|
        v = allcard[k]
        if v["卡片种类"] != "通常怪兽"
    
        if i == 0 
            j = j +1
            print "

func elses#{j}(cardBag *ygo.CardVersion) {
    var co *ygo.CardOriginal
    "
        end

        
            print "
    co = &ygo.CardOriginal{
        /*#{_puts v}*/
        Id:       #{v["id"]},
        Password: \"#{v["卡片密码"]}\",
        Name:     \"#{v["中文名"]}\", // \"#{v["英文名"]}\"  \"#{v["日文名"]}\"
        Lc:       #{lc(v["卡片种类"])}, // #{v["卡片种类"]}
        Level:     #{ if v["星级"] != nil then v["星级"] else 0 end},
        La:       #{la(v["属性"])}, // #{v["属性"]}
        Lr:       #{lr(v["种族"])}, // #{v["种族"]}
        Attack:   #{ if v["攻击力"] != nil && v["攻击力"] != "？" && v["攻击力"] != "?" then v["攻击力"] else 0 end},
        Defense:  #{ if v["防御力"] != nil && v["防御力"] != "？" && v["防御力"] != "?" then v["防御力"] else 0 end},
        //Initiative:    func(ca *Card)) {}, // 主动发动
        //Declaration:   func(ca *Card)) {}, // 攻击宣言
        //Damage:        func(ca *Card)) {}, // 伤害计算
        //Freedom:       func(ca *Card)) {}, // 解放    送去墓地
        //Destroy:       func(ca *Card)) {}, // 战斗破坏 送去墓地
        //Flip:          func(ca *Card)) {}, // 反转
        Summon:          ygo.SuperiorCall, // 召唤
        IsValid:       false,
}
    cardBag.Register(co)

"

        i = i + 1
        if i == 1000
            i = 0
            print '}'
        end
    end
end
print '}'
end

def changename allcard 
    sor = {
        "通常怪兽"=> 11,
        "效果怪兽"=> 12,
        "融合怪兽"=> 13,
        "仪式怪兽"=> 14,
        "同调怪兽"=> 15,
        "XYZ怪兽"=> 16,
        "通常魔法"=> 21,
        "速攻魔法"=> 22,
        "永续魔法"=> 23,
        "场地魔法"=> 24,
        "装备魔法"=> 25,
        "仪式魔法"=> 26,
        "通常陷阱"=> 31,
        "永续陷阱"=> 32,
        "反击陷阱"=> 33
    }
    types={}
    newcard = {}
    allcard.each do|k,v|
        types[v["卡片种类"]] = 0 if types[v["卡片种类"]] == nil
        types[v["卡片种类"]] += 1
        id = types[v["卡片种类"]] + sor[v["卡片种类"]]*10000
        v["id"]=id
        newcard[id]=v
        #{}`cp ./img/#{k}.jpg ./tmp/#{id}.jpg`
        #print v.to_json
        #{}`echo '#{v.to_json.to_utf8!}' > ./tmp/#{id}.json`
        b = File.new("./tmp/#{id}.json","w")
        b.puts v.to_json
        b.close
    end
    print newcard.to_json
end

if $0 == __FILE__


  allcard = JSON.parse File.read("card.json").to_utf8!
  #types={}
  #卡片种类
  elses allcard

end