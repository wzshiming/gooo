package main

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gooo/configs"
	"gooo/encoder"
	"gooo/protocol"

	authprot "service/auth/protocol"
)

type PassAuth struct {
	status *Status
	db     gorm.DB
}

func NewPassAuth(m *Status) *PassAuth {
	r := PassAuth{
		status: m,
	}
	us := m.Conf.Dc["Users"]
	r.db, _ = gorm.Open(us.Dialect, us.Source)
	return &r
}

func (r *PassAuth) ChangePwd(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var p authprot.ChangePwdRequest
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
	var ouser authprot.User
	if err := r.db.Where(&authprot.User{
		Username: p.Username,
	}).First(&ouser).Error; err != nil {
		return errors.New(Trans.Value("auth.usernotexists"))
	}
	if ouser.Password != p.Password {
		return errors.New(Trans.Value("auth.pwderr"))
	}
	r.db.Model(&ouser).Update(&authprot.User{Password: p.NewPassword})
	return nil
}

func (r *PassAuth) Unregister(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var p authprot.LogInRequest
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
	var ouser authprot.User
	if err := r.db.Where(&authprot.User{Username: p.Username}).First(&ouser).Error; err != nil {
		return errors.New(Trans.Value("auth.usernotexists"))
	}
	if ouser.Password != p.Password {
		return errors.New(Trans.Value("auth.pwderr"))
	}

	var gonli authprot.GetOnlineResponse
	if r.status.ServiceOffline.GetOnline(authprot.GetOnlineRequest{
		UserId: ouser.Id,
	}, &gonli); gonli.Online {
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
	var sonli authprot.InterruptResponse
	r.status.ServiceOffline.Interrupt(authprot.InterruptRequest{
		Data: args.Session.Data,
	}, &sonli)

	*reply = protocol.RpcResponse{
		Data: &map[string]interface{}{
			"userId": 0,
		},
	}
	return nil
}
