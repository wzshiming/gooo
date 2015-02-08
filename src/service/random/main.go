package main

import (
	"gooo"
)

type Status struct {
	gooo.Status
	Hand          *gooo.Handeln
	Conf          *gooo.Configs
	ServiceRandom *Random
}

func (r *Status) Init(args gooo.InitRequest, reply *int) (err error) {
	if args.State == 1 {
		r.Conf = &args.Conf
		if r.ServiceRandom == nil {
			r.ServiceRandom = NewRandom(r)
			r.Hand.Register(r.ServiceRandom)
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
