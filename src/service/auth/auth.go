package main

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	i18n "github.com/kortem/lingo"
	"gooo/handeln"
	//_ "github.com/lib/pq"
	"gooo/configs"
	"gooo/encoder"
	"gooo/helper"
	"gooo/protocol"
	"gooo/router"
	authprtc "service/auth/protocol"
	connprtc "service/connect/protocol"
	offlprot "service/offline/protocol"
)

type Auth struct {
	conf     *configs.Configs
	db       gorm.DB
	i18n     *i18n.L
	callconn *router.CallServer
	calloffl *router.CallServer
}

func NewAuth() *Auth {
	return &Auth{i18n: i18n.New("zh_CN", "i18n")}
}

func (r *Auth) Register(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var p authprtc.RegisterRequest
	encoder.Decode(args.Request, &p)
	var d struct {
		Language string
	}
	encoder.Decode(args.Session.Data, &d)
	Trans := r.i18n.TranslationsForLocale(d.Language)
	if len(p.Password) <= 6 {
		return errors.New(Trans.Value("auth.pwdshort"))
	}
	if err := r.db.Create(authprtc.User{
		Username: p.Username,
		Password: p.Password,
	}).Error; err != nil {
		return errors.New(Trans.Value("auth.userexists"))
	}
	return nil
}

func (r *Auth) ChangePwd(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var p authprtc.ChangePwdRequest
	encoder.Decode(args.Request, &p)
	var d struct {
		UserId   int64
		Language string
	}

	encoder.Decode(args.Session.Data, &d)
	Trans := r.i18n.TranslationsForLocale(d.Language)

	if len(p.NewPassword) <= 6 {
		return errors.New(Trans.Value("auth.newpwdshort"))
	}
	var ouser authprtc.User
	if err := r.db.Where(&authprtc.User{
		Username: p.Username,
	}).First(&ouser).Error; err != nil {
		return errors.New(Trans.Value("auth.usernotexists"))
	}
	if ouser.Password != p.Password {
		return errors.New(Trans.Value("auth.pwderr"))
	}
	r.db.Model(&ouser).Update(&authprtc.User{Password: p.NewPassword})
	return nil
}

func (r *Auth) Unregister(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var p authprtc.LogInRequest
	encoder.Decode(args.Request, &p)
	var d struct {
		Language string
		UserId   int64
	}
	encoder.Decode(args.Session.Data, &d)
	Trans := r.i18n.TranslationsForLocale(d.Language)
	if d.UserId != 0 {
		return errors.New(Trans.Value("auth.islogin"))
	}
	var ouser authprtc.User
	if err := r.db.Where(&authprtc.User{Username: p.Username}).First(&ouser).Error; err != nil {
		return errors.New(Trans.Value("auth.usernotexists"))
	}
	if ouser.Password != p.Password {
		return errors.New(Trans.Value("auth.pwderr"))
	}
	var gonli connprtc.GetOnlineResponse
	r.callconn.CallBySession(args.Session, "Connect.GetOnline", connprtc.GetOnlineRequest{
		UserId: ouser.Id,
	}, &gonli)
	if gonli.Online {
		return errors.New(Trans.Value("auth.inlogin"))
	}
	r.db.Delete(ouser)
	return nil
}

func (r *Auth) LogIn(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var p authprtc.LogInRequest
	encoder.Decode(args.Request, &p)
	var d struct {
		Language string
		UserId   int64
	}
	encoder.Decode(args.Session.Data, &d)
	Trans := r.i18n.TranslationsForLocale(d.Language)
	if d.UserId != 0 {
		return errors.New(Trans.Value("auth.islogin"))
	}
	var ouser authprtc.User
	if err := r.db.Where(&authprtc.User{Username: p.Username}).First(&ouser).Error; err != nil {
		return errors.New(Trans.Value("auth.usernotexists"))
	}
	if ouser.Password != p.Password {
		return errors.New(Trans.Value("auth.pwderr"))
	}
	var gonli connprtc.GetOnlineResponse
	r.callconn.CallBySession(args.Session, "Connect.GetOnline", connprtc.GetOnlineRequest{
		UserId: ouser.Id,
	}, &gonli)
	if gonli.Online {
		return errors.New(Trans.Value("auth.inlogin"))
	}
	var sonli connprtc.SetOnlineResponse
	r.callconn.CallBySession(args.Session, "Connect.SetOnline", connprtc.SetOnlineRequest{
		UserId: ouser.Id,
		Online: true,
	}, &sonli)
	var rr offlprot.ReconnectionResponse
	err := r.calloffl.Call("Offline.Reconnection", offlprot.ReconnectionRequest{
		UserId: ouser.Id,
	}, &rr)
	if err == nil {
		*reply = protocol.RpcResponse{
			Coverage: rr.Data,
			Response: []byte("{\"s\":\"Reconnection\"}"),
		}
	} else {
		*reply = protocol.RpcResponse{
			Data: &map[string]interface{}{
				"UserId": ouser.Id,
			},
		}
	}
	return nil
}

func (r *Auth) LogOut(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var d struct {
		UserId   uint64
		Language string
	}
	encoder.Decode(args.Session.Data, &d)
	Trans := r.i18n.TranslationsForLocale(d.Language)
	if d.UserId == 0 {
		return errors.New(Trans.Value("auth.nologin"))
	}
	var sonli connprtc.SetOnlineResponse
	r.callconn.CallBySession(args.Session, "Connect.SetOnline", connprtc.SetOnlineRequest{
		UserId: d.UserId,
		Online: false,
	}, &sonli)
	*reply = protocol.RpcResponse{
		Data: &map[string]interface{}{
			"UserId": 0,
		},
	}
	return nil
}

func (r *Auth) Init(args protocol.InitRequest, reply *int) (err error) {
	if args.State == 1 {
		r.conf = &args.Conf
		us := r.conf.Dc["Users"]
		r.db, err = gorm.Open(us.Dialect, us.Source)
		r.callconn = router.NewCallServer("Connect", r.conf)
		r.calloffl = router.NewCallServer("Offline", r.conf)
	}
	return nil
}
func main() {
	defer helper.Recover()
	h := handeln.NewHandeln()
	c := NewAuth()
	h.Register(c)
	h.Register(handeln.NewStatus(h))
	helper.EchoPortInfo(configs.Name, configs.Port)
	h.Start(configs.Port)
}
