package main

import (
	"gooo/handeln"

	"gooo/configs"
	"gooo/helper"
	"gooo/protocol"
)

type Status struct {
	handeln.Status
	Hand        *handeln.Handeln
	Conf        *configs.Configs
	ServiceUse  *Use
	ServiceInfo *Info
}

func (r *Status) Init(args protocol.InitRequest, reply *int) (err error) {
	if args.State == 1 {
		r.Conf = &args.Conf
		if r.ServiceUse == nil {
			r.ServiceUse = NewUse(r)
			r.Hand.Register(r.ServiceUse)
		}
		if r.ServiceInfo == nil {
			r.ServiceInfo = NewInfo(r)
			r.Hand.Register(r.ServiceInfo)
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
