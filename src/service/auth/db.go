package main

import (
	"service/proto"

	"github.com/wzshiming/dbs"
	"github.com/wzshiming/server/cfg"
)

var db dbs.DB

func dbconn(us cfg.DbConfig) (err error) {
	db, err = dbs.Conn(us)
	var users proto.User
	users.CreateTable(&db)
	return
}
