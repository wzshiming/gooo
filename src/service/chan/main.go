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
	Hand          *handeln.Handeln
	Conf          *configs.Configs
	ServiceChan   *Chan
	ServiceInChan *InChan
	ServiceInfo   *Info
}

func (r *Status) Init(args protocol.InitRequest, reply *int) (err error) {
	if args.State == 1 {
		r.Conf = &args.Conf
		if r.ServiceChan == nil {
			r.ServiceChan = NewChan(r)
			r.Hand.Register(r.ServiceChan)
		}
		if r.ServiceInChan == nil {
			r.ServiceInChan = NewInChan(r)
			r.Hand.Register(r.ServiceInChan)
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
