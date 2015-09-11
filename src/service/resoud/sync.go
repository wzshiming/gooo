package main

import (
	"fmt"
	"strings"

	"github.com/wzshiming/base"
	"github.com/wzshiming/server/cfg"
)

func runSync(cfg cfg.ServerConfig) {
	url := fmt.Sprintf("http://%s:%d", cfg.Host, cfg.ClientPort)
	b := base.OpenUrl(url + "/versions.json")
	nh := map[string]string{}
	base.NewEncodeBytes(b).DeJson(&nh)
	h := base.FilesHashIndexing("../static/web")

	base.SaveLocal("../static/web/versions.json", b)
	for k, v := range nh {
		if h[k] != v {
			path := strings.Replace(k, "../static/", "/", 1)
			file := base.OpenUrl(url + path)
			base.SaveLocal(k, file)
		}
	}
}
