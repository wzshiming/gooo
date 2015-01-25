package main

import (
	i18n "github.com/kortem/lingo"
	"gooo/handeln"
	//_ "github.com/lib/pq"
	"gooo/configs"
	"gooo/helper"
	"gooo/protocol"
)

type Status struct {
	handeln.Status
	Hand           *handeln.Handeln
	Conf           *configs.Configs
	I18n           *i18n.L
	ServiceConnect *Connect
}

func (r *Status) Init(args protocol.InitRequest, reply *int) (err error) {
	if args.State == 1 {
		r.Conf = &args.Conf
		if r.ServiceConnect == nil {
			r.ServiceConnect = NewConnect(r)
			r.Hand.Register(r.ServiceConnect)
		}
	}
	return nil
}

func main() {
	defer helper.Recover()
	h := handeln.NewHandeln()
	h.Register(&Status{
		Hand:   h,
		Status: *handeln.NewStatus(h),
		I18n:   i18n.New("zh_CN", "i18n"),
	})
	helper.EchoPortInfo(configs.Name, configs.Port)
	h.Start(configs.Port)

}
