package configs

import (
	"gooo/addr"
	"gooo/helper"
)

var Type, Port, Name, Id = helper.FlagVar()
var IP = addr.GetIP()
