package main

import (
	"gooo"
	"gooo/route"
	"time"
)

type Status struct {
	gooo.Status
	Hand           *gooo.Handeln
	Conf           *gooo.Configs
	ServiceConnect *route.ConnRpc
	CilentHandeln  *route.Handel
}

func (r *Status) Init(args gooo.InitRequest, reply *int) (err error) {
	if args.State == 1 {
		r.Conf = &args.Conf
		if r.CilentHandeln == nil {
			go func() {
				time.Sleep(time.Second)
				if r.CilentHandeln == nil {
					r.CilentHandeln = route.NewHandel(r.Conf, 4096,
						"",
						"Auth",
						"Chan",
					)
					ser := gooo.NewServer(r.CilentHandeln, route.NewIORange(1024))
					Run(r.Conf, ser.WebsocketHandler().ServeHTTP)
				}
				if r.ServiceConnect == nil {
					r.ServiceConnect = route.NewConnRpc(r.CilentHandeln)
					r.Hand.Register(r.ServiceConnect)
				}
			}()
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
