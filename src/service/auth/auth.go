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

type Auth struct {
	status *Status
	db     gorm.DB
}

func NewAuth(m *Status) *Auth {
	r := Auth{
		status: m,
	}
	us := m.Conf.Dc["Users"]
	r.db, _ = gorm.Open(us.Dialect, us.Source)
	return &r
}

func (r *Auth) Register(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var p authprot.RegisterRequest
	encoder.Decode(args.Request, &p)
	var d struct {
		Language string
	}
	encoder.Decode(args.Session.Data, &d)
	Trans := configs.I18n.TranslationsForLocale(d.Language)
	if len(p.Password) <= 6 {
		return errors.New(Trans.Value("auth.pwdshort"))
	}

	var ouser authprot.User
	if err := r.db.Where(&authprot.User{Username: p.Username}).First(&ouser).Error; err == nil {
		return errors.New(Trans.Value("auth.userexists"))
	}

	if err := r.db.Create(authprot.User{
		Username: p.Username,
		Password: p.Password,
	}).Error; err != nil {
		return errors.New(Trans.Value("auth.userexists"))
	}
	return nil
}

func (r *Auth) LogIn(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var p authprot.LogInRequest
	encoder.Decode(args.Request, &p)
	var d struct {
		Flag     uint32 `json:"flag"`
		Language string
		UserId   uint64 `json:"userId"`
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

	var rr authprot.ReconnectionResponse
	if err := r.status.ServiceOffline.Reconnection(authprot.ReconnectionRequest{
		UserId: ouser.Id,
		Unique: args.Session.Uniq,
	}, &rr); err != nil {
		return err
	}

	*reply = protocol.RpcResponse{
		Coverage: rr.Data,
		Data: &map[string]interface{}{
			"userId": ouser.Id,
			"flag":   d.Flag | configs.FlagLogin,
		},
	}
	return nil
}
