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
