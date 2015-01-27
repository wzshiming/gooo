package main

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gooo/configs"
	"gooo/encoder"
	"gooo/protocol"
	"gooo/router"
	authprtc "service/auth/protocol"
	connprtc "service/connect/protocol"
)

type PassAuth struct {
	status   *Status
	db       gorm.DB
	callconn *router.CallServer
	calloffl *router.CallServer
}

func NewPassAuth(m *Status) *PassAuth {
	r := PassAuth{
		status: m,
	}
	us := m.Conf.Dc["Users"]
	r.db, _ = gorm.Open(us.Dialect, us.Source)
	r.callconn = router.NewCallServer("Connect", m.Conf)
	r.calloffl = router.NewCallServer("Offline", m.Conf)
	return &r
}

func (r *PassAuth) ChangePwd(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var p authprtc.ChangePwdRequest
	encoder.Decode(args.Request, &p)
	var d struct {
		UserId   int64 `json:"userId"`
		Language string
	}

	encoder.Decode(args.Session.Data, &d)
	Trans := configs.I18n.TranslationsForLocale(d.Language)

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

func (r *PassAuth) Unregister(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var p authprtc.LogInRequest
	encoder.Decode(args.Request, &p)
	var d struct {
		Language string
		UserId   int64 `json:"userId"`
	}
	encoder.Decode(args.Session.Data, &d)
	Trans := configs.I18n.TranslationsForLocale(d.Language)
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

func (r *PassAuth) LogOut(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var d struct {
		UserId   uint64 `json:"userId"`
		Language string
	}
	encoder.Decode(args.Session.Data, &d)
	Trans := configs.I18n.TranslationsForLocale(d.Language)
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
			"userId": 0,
		},
	}
	return nil
}
