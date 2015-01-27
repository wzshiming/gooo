package main

import (
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
	ServiceOffline *Offline
}

func (r *Status) Init(args protocol.InitRequest, reply *int) (err error) {
	if args.State == 1 {
		r.Conf = &args.Conf
		if r.ServiceOffline == nil {
			r.ServiceOffline = NewOffline(5000)
			r.Hand.Register(r.ServiceOffline)
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
	})
	helper.EchoPortInfo(configs.Name, configs.Port)
	h.Start(configs.Port)

}
