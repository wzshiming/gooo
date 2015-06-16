package main

import (
	"github.com/wzshiming/rego/cfg"
	"github.com/wzshiming/rego/dbs"
	"service/proto"
)

var db dbs.DB

func dbconn(us cfg.DbConfig) (err error) {
	db, err = dbs.Conn(us)
	var users proto.User
	users.CreateTable(&db)
	return
}
