package main

import (
	"errors"
	"gooo"
	"gooo/protocol"
)

type Auth struct {
	gooo.Methods
	status *Status
}

func NewAuth(m *Status) *Auth {
	r := Auth{
		status: m,
	}
	r.Methods = *gooo.NewMethods(
		gooo.FlagNone,
		gooo.FlagLogin,
		"Register",
		"LogIn",
	)
	return &r
}

func (r *Auth) Register(args gooo.RpcRequest, reply *gooo.RpcResponse) error {
	var p protocol.RegisterRequest
	gooo.Decode(args.Request, &p)
	var d struct {
		Language string
	}
	gooo.Decode(args.Session.Data, &d)
	Trans := I18n.TranslationsForLocale(d.Language)
	if len(p.Password) < 6 {
		return errors.New(Trans.Value("auth.pwdshort"))
	}

	var ouser protocol.User
	if err := db.Where(&protocol.User{Username: p.Username}).First(&ouser).Error; err == nil {
		return errors.New(Trans.Value("auth.userexists"))
	}

	if err := db.Create(&protocol.User{
		Username: p.Username,
		Password: p.Password,
	}).Error; err != nil {
		return errors.New(Trans.Value("auth.userexists"))
	}
	return nil
}

func (r *Auth) LogIn(args gooo.RpcRequest, reply *gooo.RpcResponse) error {
	var p protocol.LogInRequest
	gooo.Decode(args.Request, &p)
	var d struct {
		Flag     uint32 `json:"flag"`
		Language string
		UserId   uint64 `json:"userId"`
	}
	gooo.Decode(args.Session.Data, &d)
	Trans := I18n.TranslationsForLocale(d.Language)
	if d.UserId != 0 {
		return errors.New(Trans.Value("auth.islogin"))
	}
	var ouser protocol.User
	if err := db.Where(&protocol.User{Username: p.Username}).First(&ouser).Error; err != nil {
		return errors.New(Trans.Value("auth.usernotexists"))
	}

	if ouser.Username != p.Username {
		return errors.New(Trans.Value("auth.usernotexists"))
	}

	if ouser.Password != p.Password {
		return errors.New(Trans.Value("auth.pwderr"))
	}

	var rr protocol.ReconnectionResponse
	if err := r.status.ServiceOffline.Reconnection(protocol.ReconnectionRequest{
		UserId: ouser.Id,
		Unique: args.Session.Uniq,
	}, &rr); err != nil {
		return err
	}

	*reply = gooo.RpcResponse{
		Coverage: rr.Data,
		Data: &map[string]interface{}{
			"userId": ouser.Id,
			"flag":   d.Flag | gooo.FlagLogin,
		},
	}
	return nil
}
