package main

import (
	"rego"
	//"rego/server/proto"
)

type Auth struct {
}

func NewAuth() *Auth {
	r := Auth{}
	return &r
}
func (r *Auth) Register(args rego.Request, reply *rego.Response) error {
	rego.INFO(string(args.Request.Bytes()))
	es := rego.NewEncodeStream()
	es.EnJson(map[string]string{
		"111": "222",
	})
	reply.Response = es
	return nil
}

//func (r *Auth) Register(args rego.Request, reply *rego.Response) error {
//	var p proto.RegisterRequest

//	rego.Decode(args.Request, &p)
//	var d struct {
//		Language string
//	}
//	rego.Decode(args.Session.Data, &d)
//	Trans := I18n.TranslationsForLocale(d.Language)
//	if len(p.Password) < 6 {
//		return errors.New(Trans.Value("auth.pwdshort"))
//	}

//	var ouser proto.User
//	if err := db.Where(&proto.User{Username: p.Username}).First(&ouser).Error; err == nil {
//		return errors.New(Trans.Value("auth.userexists"))
//	}

//	if err := db.Create(&proto.User{
//		Username: p.Username,
//		Password: p.Password,
//	}).Error; err != nil {
//		return errors.New(Trans.Value("auth.userexists"))
//	}
//	return nil
//}

//func (r *Auth) LogIn(args rego.Request, reply *rego.Response) error {
//	var p proto.LogInRequest
//	rego.Decode(args.Request, &p)
//	var d struct {
//		Flag     uint32 `json:"flag"`
//		Language string
//		UserId   uint64 `json:"userId"`
//	}
//	rego.Decode(args.Session.Data, &d)
//	Trans := I18n.TranslationsForLocale(d.Language)
//	if d.UserId != 0 {
//		return errors.New(Trans.Value("auth.islogin"))
//	}
//	var ouser proto.User
//	if err := db.Where(&proto.User{Username: p.Username}).First(&ouser).Error; err != nil {
//		return errors.New(Trans.Value("auth.usernotexists"))
//	}

//	if ouser.Username != p.Username {
//		return errors.New(Trans.Value("auth.usernotexists"))
//	}

//	if ouser.Password != p.Password {
//		return errors.New(Trans.Value("auth.pwderr"))
//	}

//	var rr proto.ReconnectionResponse
//	if err := r.status.ServiceOffline.Reconnection(proto.ReconnectionRequest{
//		UserId: ouser.Id,
//		Unique: args.Session.Uniq,
//	}, &rr); err != nil {
//		return err
//	}

//	*reply = rego.Response{
//		Coverage: rr.Data,
//		Data: &map[string]interface{}{
//			"userId": ouser.Id,
//			"flag":   d.Flag | rego.FlagLogin,
//		},
//	}
//	return nil
//}
