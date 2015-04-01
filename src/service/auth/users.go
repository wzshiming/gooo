package main

import (
	"errors"
	"rego"
	"rego/agent"
	"rego/misc"
	"service/proto"
	"time"
)

type Users struct {
	room *misc.Rooms
}

func NewUsers() *Users {
	r := Users{
		room: misc.NewRooms("Users"),
	}
	i := 0
	tick := time.Tick(time.Second * 10)
	go func() {
		for {
			<-tick
			rego.NOTICE(r.room.Len())
			i++
			r.room.Broadcast(&agent.Response{
				Response: rego.EnJson(map[string]interface{}{
					"hello": "world",
					"index": i,
				}),
			}, nil)
		}
	}()
	return &r
}

func (r *Users) Register(args agent.Request, reply *agent.Response) error {
	var p proto.RegisterRequest
	args.Request.DeJson(&p)
	var d struct {
		Language string
	}
	args.Session.Data.DeJson(&d)
	Trans := i18n.TranslationsForLocale(d.Language)

	// 判断密码是否符合要求
	if len(p.Password) < 6 {
		return errors.New(Trans.Value("auth.pwdshort"))
	}

	// 检查该用户是否存在
	var ouser proto.User
	if err := db.Where(&proto.User{Username: p.Username}).First(&ouser).Error; err == nil {
		return errors.New(Trans.Value("auth.userexists"))
	}

	// 创建用户
	if err := db.Create(&proto.User{
		Username: p.Username,
		Password: p.Password,
	}).Error; err != nil {
		return errors.New(Trans.Value("auth.userexists"))
	}
	return nil
}

func (r *Users) LogIn(args agent.Request, reply *agent.Response) error {
	var p proto.RegisterRequest
	args.Request.DeJson(&p)
	var d struct {
		Language string
	}
	args.Session.Data.DeJson(&d)
	Trans := i18n.TranslationsForLocale(d.Language)

	// 判断用户是否登入
	if sess := r.room.Sync(args.Session); sess != nil {
		return errors.New(Trans.Value("auth.islogin"))
	}

	// 向数据库查询用户
	var ouser proto.User
	if err := db.Where(&proto.User{Username: p.Username}).First(&ouser).Error; err != nil {
		return errors.New(Trans.Value("auth.usernotexists"))
	}

	// 判断用户 账号密码是否正确
	if ouser.Username != p.Username {
		return errors.New(Trans.Value("auth.usernotexists"))
	}
	if ouser.Password != p.Password {
		return errors.New(Trans.Value("auth.pwderr"))
	}

	// 用户登入
	if ret, err := r.room.JoinFrom(uint(ouser.Id), args.Session, args.Head); err != nil {
		return errors.New(Trans.Value("auth.islogin"))
	} else {
		reply.Data = ret
	}
	return nil
}

func (r *Users) ChangePwd(args agent.Request, reply *agent.Response) error {
	var p proto.ChangePwdRequest
	args.Request.DeJson(&p)
	var d struct {
		Language string
	}
	args.Session.Data.DeJson(&d)
	Trans := i18n.TranslationsForLocale(d.Language)

	// 判断用户是否登入
	if sess := r.room.Sync(args.Session); sess == nil {
		return errors.New(Trans.Value("auth.nologin"))
	}

	// 判断新密码是否符合要求
	if len(p.NewPassword) < 6 {
		return errors.New(Trans.Value("auth.newpwdshort"))
	}

	// 检查用户是否存在
	var ouser proto.User
	if err := db.Where(&proto.User{Username: p.Username}).First(&ouser).Error; err != nil {
		return errors.New(Trans.Value("auth.usernotexists"))
	}

	// 判断用户 账号密码是否正确
	if ouser.Username != p.Username {
		return errors.New(Trans.Value("auth.usernotexists"))
	}
	if ouser.Password != p.Password {
		return errors.New(Trans.Value("auth.pwderr"))
	}

	// 修改用户密码
	db.Model(&ouser).Update(&proto.User{Password: p.NewPassword})
	return nil
}

func (r *Users) Unregister(args agent.Request, reply *agent.Response) error {
	var p proto.UnregisterRequest
	args.Request.DeJson(&p)
	var d struct {
		Language string
	}
	args.Session.Data.DeJson(&d)
	Trans := i18n.TranslationsForLocale(d.Language)

	// 判断用户是否登入
	if sess := r.room.Sync(args.Session); sess == nil {
		return errors.New(Trans.Value("auth.nologin"))
	}

	// 检查用户是否存在
	var ouser proto.User
	if err := db.Where(&proto.User{Username: p.Username}).First(&ouser).Error; err != nil {
		return errors.New(Trans.Value("auth.usernotexists"))
	}

	// 判断用户 账号密码是否正确
	if ouser.Username != p.Username {
		return errors.New(Trans.Value("auth.usernotexists"))
	}
	if ouser.Password != p.Password {
		return errors.New(Trans.Value("auth.pwderr"))
	}

	// 删除用户
	db.Delete(ouser)
	return nil
}

func (r *Users) LogOut(args agent.Request, reply *agent.Response) error {
	var d struct {
		Language string
	}
	args.Session.Data.DeJson(&d)
	Trans := i18n.TranslationsForLocale(d.Language)

	// 判断用户是否登入
	sess := r.room.Sync(args.Session)
	if sess == nil {
		return errors.New(Trans.Value("auth.nologin"))
	}

	// 用户登出
	reply.Data = r.room.Leave(sess)
	return nil
}
