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
    "效果怪兽"=>"FusionMonster",
    "XYZ怪兽"=>"ExcessMonster",
    "同调怪兽"=>"HomologyMonster",
    "仪式怪兽"=>"RiteMonster",
    "通常魔法"=>"OrdinaryMagic",
    "仪式魔法"=>"RiteMagic",
    "永续魔法"=>"SustainsMagic",
    "装备魔法"=>"EquipMagic",
    "场地魔法"=>"PlaceMagic",
    "速攻魔法"=>"RushMagic",
    "普通陷阱"=>"OrdinaryTrap",
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
    "魔法师"=>"Spellcaster",
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
    "幻龙"=>"MagicDragon"
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


func original(CardBag *ygo.CardVersion) {
    var co *ygo.CardOriginal
    '
    allcard.each do|k,v|
        if v["卡片种类"] == "通常怪兽"
            print "
    co = &ygo.CardOriginal{
        /*  
        #{v["调整"]}
        #{v["使用限制"]}
        #{v["罕见度"]}
        #{v["卡包"]}
        */
        Id:       #{v["id"]},
        Password: \"#{v["卡片密码"]}\",
        Name:     \"#{v["中文名"]}\", // \"#{v["英文名"]}\"  \"#{v["日文名"]}\"
        Describe: \"#{v["效果"].gsub!("$\n","")}\",
        Lc:       #{lc(v["卡片种类"])}, // #{v["卡片种类"]}
        Star:     #{v["星级"]},
        La:       #{la(v["属性"])}, // #{v["属性"]}
        Lr:       #{lr(v["种族"])}, // #{v["种族"]}
        Attack:   #{v["攻击力"]},
        Defense:  #{v["防御力"]},
    }
    CardBag.Register(co)

"
        end
    end
    print '}'
end
if $0 == __FILE__


  t = File.read("CardBag.go").to_utf8!

  allcard = JSON.parse File.read("card.json").to_utf8!
  types={}
  #卡片种类
  OrdinaryMonster allcard

end