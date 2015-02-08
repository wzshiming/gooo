package main

import (
	"gooo"
	"service/connect/route"
)

type Status struct {
	gooo.Status
	Hand           *gooo.Handeln
	Conf           *gooo.Configs
	ServiceConnect *Connect
	CilentHandeln  *route.Handel
}

func (r *Status) Init(args gooo.InitRequest, reply *int) (err error) {
	if args.State == 1 {
		r.Conf = &args.Conf
		if r.ServiceConnect == nil {
			r.ServiceConnect = NewConnect(r)
			r.Hand.Register(r.ServiceConnect)
		}
		if r.CilentHandeln == nil {
			r.CilentHandeln = route.NewHandel(r.Conf, 4096)
		}
	}
	return nil
}

func main() {
	defer gooo.Recover()
	h := gooo.NewHandeln()
	h.Register(&Status{
		Hand:   h,
		Status: *gooo.NewStatus(h),
	})
	gooo.EchoPortInfo(gooo.Name, gooo.Port)
	h.Start(gooo.Port)

}
