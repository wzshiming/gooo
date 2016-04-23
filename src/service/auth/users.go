package main

import (
	"service/proto"
	"time"

	"github.com/wzshiming/base"
	"github.com/wzshiming/server/agent"
)

type Users struct {
	room *agent.Room
}

func NewUsers() *Users {
	r := Users{
		room: agent.NewRoom("Users"),
	}
	//i := 0
	o := 0
	tick := time.Tick(time.Second * 5)
	go func() {
		for {
			<-tick
			if o != r.room.Len() {
				base.NOTICE("User size:", r.room.Len())
				o = r.room.Len()
			}
			//			i++
			//			r.room.ForEach(func(sess *agent.Session) {
			//				sess.Mutex(func() {

			//					i := 0
			//					sess.Data.Get("size", &i)
			//					i++
			//					sess.Data.Set("size", i)

			//					//base.INFO(sess.Rooms.Data())
			//					//sess.PushForm(map[string]int{"a": d.Size}, r.room.Head(sess))
			//				})
			//			})
		}
	}()
	return &r
}

func (r *Users) Register(args agent.Request, reply *agent.Response) error {
	args.Mutex(reply, func() {
		var p proto.RegisterRequest
		args.Request.DeJson(&p)

		// 判断密码是否符合要求
		if len(p.Password) < 6 {
			reply.ReplyError("auth.pwdshort")
			return
		}

		// 检查该用户是否存在
		var ouser proto.User
		if err := db.Where(&proto.User{Username: p.Username}).First(&ouser).Error; err == nil {
			reply.ReplyError("auth.userexists")
			return
		}

		// 创建用户
		if err := db.Create(&proto.User{
			Username: p.Username,
			Password: p.Password,
		}).Error; err != nil {
			reply.ReplyError("auth.userexists")
			return
		}
		reply.ReplyError("")
	})

	return nil
}

func (r *Users) LogIn(args agent.Request, reply *agent.Response) error {

	args.Mutex(reply, func() {
		var p proto.RegisterRequest
		args.Request.DeJson(&p)

		// 判断用户是否登入
		if r.room.IsExist(args.Session) {
			reply.ReplyError("auth.islogin")
			return
		}

		// 向数据库查询用户
		var ouser proto.User
		if err := db.Where(&proto.User{Username: p.Username}).First(&ouser).Error; err != nil {
			reply.ReplyError("auth.pwderr")
			return

		}

		// 判断用户 账号密码是否正确
		if ouser.Username != p.Username {
			reply.ReplyError("auth.pwderr")
			return
		}
		if ouser.Password != p.Password {
			reply.ReplyError("auth.pwderr")
			return
		}
		if r.room.Get(uint(ouser.Id)) != nil {
			reply.ReplyError("auth.useruseing")
			return
		}

		args.Session.Data.Set("username", ouser.Username)
		args.Session.Data.Set("userid", ouser.Id)

		// 用户登入
		r.room.JoinFrom(uint(ouser.Id), args.Session, args.Head)

		reply.ReplyError("")
	})

	return nil
}

func (r *Users) ChangePwd(args agent.Request, reply *agent.Response) error {
	args.Mutex(reply, func() {
		var p proto.ChangePwdRequest
		args.Request.DeJson(&p)

		// 判断用户是否登入
		if !r.room.IsExist(args.Session) {
			reply.ReplyError("auth.nologin")
			return
		}

		// 判断新密码是否符合要求
		if len(p.NewPassword) < 6 {
			reply.ReplyError("auth.newpwdshort")
			return
		}

		// 检查用户是否存在
		var ouser proto.User
		if err := db.Where(&proto.User{Username: p.Username}).First(&ouser).Error; err != nil {
			reply.ReplyError("auth.pwderr")
			return
		}

		// 判断用户 账号密码是否正确
		if ouser.Username != p.Username {
			reply.ReplyError("auth.pwderr")
			return
		}

		if ouser.Password != p.Password {
			reply.ReplyError("auth.pwderr")
			return
		}

		// 修改用户密码
		db.Model(&ouser).Update(&proto.User{Password: p.NewPassword})
		reply.ReplyError("")
	})
	return nil
}

//func (r *Users) Unregister(args agent.Request, reply *agent.Response) error {
//	args.Mutex(reply, func() {
//		var p proto.UnregisterRequest
//		args.Request.DeJson(&p)

//		// 判断用户是否登入
//		if !r.room.IsExist(args.Session) {
//			reply.ReplyError("auth.nologin")
//			return
//		}

//		// 检查用户是否存在
//		var ouser proto.User
//		if err := db.Where(&proto.User{Username: p.Username}).First(&ouser).Error; err != nil {
//			reply.ReplyError("auth.pwderr")
//			return
//		}

//		// 判断用户 账号密码是否正确
//		if ouser.Username != p.Username {
//			reply.ReplyError("auth.pwderr")
//			return
//		}
//		if ouser.Password != p.Password {
//			reply.ReplyError("auth.pwderr")
//			return
//		}

//		// 删除用户
//		db.Delete(ouser)
//		reply.ReplyError("")
//	})
//	return nil
//}

func (r *Users) LogOut(args agent.Request, reply *agent.Response) error {
	args.Mutex(reply, func() {
		// 判断用户是否登入
		if !r.room.IsExist(args.Session) {
			reply.ReplyError("auth.nologin")
			return
		}
		// 用户登出
		r.room.Leave(args.Session)
		reply.ReplyError("")
	})
	return nil
}
