package main

import (
	i18n "github.com/kortem/lingo"
	"gooo"
)

var I18n = i18n.New("zh_CN", "i18n")

type Status struct {
	gooo.Status
	Hand            *gooo.Handeln
	Conf            *gooo.Configs
	ServiceAuth     *Auth
	ServicePassAuth *PassAuth
	ServiceOffline  *Offline
}

func (r *Status) Init(args gooo.InitRequest, reply *int) (err error) {
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

		if r.ServiceOffline == nil {
			r.ServiceOffline = NewOffline(r, 5000)
			r.Hand.Register(r.ServiceOffline)
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
