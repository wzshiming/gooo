package main

import (
	"gooo"
)

type Status struct {
	gooo.Status
	Hand          *gooo.Handeln
	Conf          *gooo.Configs
	ServiceChan   *Chan
	ServiceInChan *InChan
	ServiceInfo   *Info
}

func (r *Status) Init(args gooo.InitRequest, reply *int) (err error) {
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
	defer gooo.Recover()
	h := gooo.NewHandeln()
	h.Register(&Status{
		Hand:   h,
		Status: *gooo.NewStatus(h),
	})
	gooo.EchoPortInfo(gooo.Name, gooo.Port)
	h.Start(gooo.Port)

}
