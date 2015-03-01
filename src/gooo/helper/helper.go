package helper

import (
	"fmt"
	"gooo"
	"gooo/route"
	ht "html/template"
	tt "text/template"
)

var Msf *route.Remap
var Conf *gooo.Configs

func Init(conf *gooo.Configs, file string) {
	Conf = conf
	Msf = route.NewRemapFrom(file)
}

var helpfunc = map[string]interface{}{
	"key": func(key string) string {
		b := Msf.Get(key)
		return fmt.Sprintf("%d,%d,%d", b[0], b[1], b[2])
	},
	"servers": func() []gooo.ServerConfig {
		return Conf.Servers().Get("Chan")
	},
}

var FuncMapsText = []tt.FuncMap{
	tt.FuncMap(helpfunc),
}

var FuncMapsHtml = []ht.FuncMap{
	ht.FuncMap(helpfunc),
}
