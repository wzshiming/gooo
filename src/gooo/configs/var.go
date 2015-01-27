package configs

import (
	i18n "github.com/kortem/lingo"
	"gooo/addr"
	"gooo/helper"
)

var Type, Port, Name, Id = helper.FlagVar()
var IP = addr.GetIP()
var I18n = i18n.New("zh_CN", "i18n")
