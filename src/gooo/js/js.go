package js

import (
	"fmt"
	"gooo/route"
	"text/template"
)

var Msf = route.NewRemapFrom("./conf/Connect_1_map.json")

var FuncMaps = []template.FuncMap{
	template.FuncMap{
		"key": func(key string) string {
			b := Msf.Get(key)
			return fmt.Sprintf("%d,%d,%d", b[0], b[1], b[2])
		},
	},
}
