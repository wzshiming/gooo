package main

import (
	"gooo"
)

type Status struct {
	gooo.Methods
	gooo.Status
	Conf          *gooo.Configs
	ServiceChan   *Chan
	ServiceInChan *InChan
	ServiceInfo   *Info
	ServiceUse    *Use
}

func (r *Status) Init(args gooo.InitRequest, reply *int) (err error) {
	if args.State == 1 {
		r.Conf = &args.Conf

		room := gooo.NewGameRooms(100)
		if r.ServiceChan == nil {
			r.ServiceChan = NewChan(r, room)
			r.Hand.Register(r.ServiceChan)
		}
		if r.ServiceInChan == nil {
			r.ServiceInChan = NewInChan(r, room)
			r.Hand.Register(r.ServiceInChan)
		}
		if r.ServiceInfo == nil {
			r.ServiceInfo = NewInfo(r, room)
			r.Hand.Register(r.ServiceInfo)
		}
		if r.ServiceUse == nil {
			r.ServiceUse = NewUse(r)
			r.Hand.Register(r.ServiceUse)
		}
	}
	return nil
}

func main() {
	defer gooo.Recover()
	h := gooo.NewHandeln()
	h.Register(&Status{
		Status: *gooo.NewStatus(h),
		Methods: *gooo.NewMethods(
			gooo.FlagLogin,
			gooo.FlagNone,
			"Chan",
			"InChan",
			"Info",
			"Use",
		),
	})
	gooo.EchoPortInfo(gooo.Name, gooo.Port)
	h.Start(gooo.Port)

}
