package main

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gooo"
	"gooo/protocol"
)

type PassAuth struct {
	status *Status
	db     gorm.DB
}

func NewPassAuth(m *Status) *PassAuth {
	r := PassAuth{
		status: m,
	}
	us := m.Conf.DataBase("Users")
	r.db, _ = gorm.Open(us.Dialect, us.Source)
	return &r
}

func (r *PassAuth) ChangePwd(args gooo.RpcRequest, reply *gooo.RpcResponse) error {
	var p protocol.ChangePwdRequest
	gooo.Decode(args.Request, &p)
	var d struct {
		UserId   int64 `json:"userId"`
		Language string
	}

	gooo.Decode(args.Session.Data, &d)
	Trans := I18n.TranslationsForLocale(d.Language)

	if len(p.NewPassword) <= 6 {
		return errors.New(Trans.Value("auth.newpwdshort"))
	}
	var ouser protocol.User
	if err := r.db.Where(&protocol.User{
		Username: p.Username,
	}).First(&ouser).Error; err != nil {
		return errors.New(Trans.Value("auth.usernotexists"))
	}
	if ouser.Password != p.Password {
		return errors.New(Trans.Value("auth.pwderr"))
	}
	r.db.Model(&ouser).Update(&protocol.User{Password: p.NewPassword})
	return nil
}

func (r *PassAuth) Unregister(args gooo.RpcRequest, reply *gooo.RpcResponse) error {
	var p protocol.LogInRequest
	gooo.Decode(args.Request, &p)
	var d struct {
		Language string
		UserId   int64 `json:"userId"`
	}
	gooo.Decode(args.Session.Data, &d)
	Trans := I18n.TranslationsForLocale(d.Language)
	if d.UserId != 0 {
		return errors.New(Trans.Value("auth.islogin"))
	}
	var ouser protocol.User
	if err := r.db.Where(&protocol.User{Username: p.Username}).First(&ouser).Error; err != nil {
		return errors.New(Trans.Value("auth.usernotexists"))
	}
	if ouser.Password != p.Password {
		return errors.New(Trans.Value("auth.pwderr"))
	}

	var gonli protocol.GetOnlineResponse
	if r.status.ServiceOffline.GetOnline(protocol.GetOnlineRequest{
		UserId: ouser.Id,
	}, &gonli); gonli.Online {
		return errors.New(Trans.Value("auth.inlogin"))
	}

	r.db.Delete(ouser)
	return nil
}

func (r *PassAuth) LogOut(args gooo.RpcRequest, reply *gooo.RpcResponse) error {
	var d struct {
		UserId   uint64 `json:"userId"`
		Language string
	}
	gooo.Decode(args.Session.Data, &d)
	Trans := I18n.TranslationsForLocale(d.Language)
	if d.UserId == 0 {
		return errors.New(Trans.Value("auth.nologin"))
	}
	var sonli protocol.InterruptResponse
	r.status.ServiceOffline.Interrupt(protocol.InterruptRequest{
		Data: args.Session.Data,
	}, &sonli)
	*reply = gooo.RpcResponse{
		Coverage: []byte("{}"),
	}
	return nil
}
