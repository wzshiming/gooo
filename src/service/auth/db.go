package main

import (
	"service/proto"

	"github.com/wzshiming/server/cfg"
	"github.com/wzshiming/server/dbs"
)

var db dbs.DB

func dbconn(us cfg.DbConfig) (err error) {
	db, err = dbs.Conn(us)
	db.CreateTable(&proto.User{})
	return
}
