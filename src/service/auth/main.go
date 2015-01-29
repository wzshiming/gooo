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
	Hand            *handeln.Handeln
	Conf            *configs.Configs
	ServiceAuth     *Auth
	ServicePassAuth *PassAuth
	ServiceOffline  *Offline
	ServiceInfo     *Info
	ServiceUse      *Use
}

func (r *Status) Init(args protocol.InitRequest, reply *int) (err error) {
	if args.State == 1 {
		r.Conf = &args.Conf
		if r.ServiceAuth == nil {
			r.ServiceAuth = NewAuth(r)
			r.Hand.Register(r.ServiceAuth)
		}
		if r.ServicePassAuth == nil {
			r.ServicePassAuth = NewPassAuth(r)
			r.Hand.Register(r.ServicePassAuth)
		}
		if r.ServiceInfo == nil {
			r.ServiceInfo = NewInfo(r)
			r.Hand.Register(r.ServiceInfo)
		}
		if r.ServiceUse == nil {
			r.ServiceUse = NewUse(r)
			r.Hand.Register(r.ServiceUse)
		}
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
