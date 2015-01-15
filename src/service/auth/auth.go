package main

import (
	//"database/sql"
	//"fmt"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"gooo/configs"
	"gooo/encoder"
	"gooo/helper"
	"gooo/protocol"
	//"log"
	authprtc "service/auth/protocol"
)

type Auth struct {
	conf *configs.Configs
	db   gorm.DB
}

func NewAuth() *Auth {
	return &Auth{}
}

func (r *Auth) Register(args protocol.RpcRequest, reply *protocol.RpcResponse) (err error) {
	var p authprtc.RegisterRequest
	encoder.Decode(args.Request, &p)

	if len(p.Password) <= 6 {
		return errors.New("密码太短")
	}

	err = r.db.Create(authprtc.User{
		Username: p.Username,
		Password: p.Password,
	}).Error
	if err != nil {
		return errors.New("用户名已存在")
	}

	ret := authprtc.OkResponse{true}
	res, _ := encoder.Encode(ret)
	*reply = protocol.RpcResponse{
		Response: res,
	}
	return nil
}

func (r *Auth) ChangePwd(args protocol.RpcRequest, reply *protocol.RpcResponse) (err error) {
	//var p authprtc.RegisterRequest
	//encoder.Decode(args.Request, &p)

	//if len(p.NewPassword) <= 6 {
	//	return errors.New("新密码太短")
	//}
	//var ouser authprtc.User

	////err = r.db.Where("Username = ?", p.Username).First(&ouser)
	//err = r.db.Updta(authprtc.User{
	//	Username: p.Username,
	//	Password: p.Password,
	//}).Error
	//if err != nil {
	//	return errors.New("用户名已存在")
	//}

	//ret := authprtc.OkResponse{true}
	//res, _ := encoder.Encode(ret)
	//*reply = protocol.RpcResponse{
	//	Response: res,
	//}
	return nil
}

func (r *Auth) Unregister(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	return nil
}

func (r *Auth) LogIn(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	return nil
}

func (r *Auth) LogOut(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	return nil
}

func (r *Auth) Init(args protocol.InitRequest, reply *int) (err error) {
	if args.State == 1 {
		r.conf = &args.Conf
		us := r.conf.Dc["Users"]
		r.db, err = gorm.Open(us.Dialect, us.Source)
		//r.db.DropTable(&authprtc.User{})
		//r.db.CreateTable(&authprtc.User{})
		//r.db.Model(&authprtc.User{}).AddUniqueIndex("idx_user_username", "username")
	}
	return nil
}
func main() {
	defer helper.Recover()
	h := helper.NewHandeln()
	c := NewAuth()
	h.Register(c)
	h.Register(helper.NewStatus(h))
	helper.EchoPortInfo(configs.Name, configs.Port)
	h.Start(configs.Port)
}
