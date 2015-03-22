package main

import (
	"rego/cfg"
	"rego/dbs"
	"service/proto"
)

var db dbs.DB

func dbconn(us cfg.DbConfig) (err error) {
	db, err = dbs.Conn(us)
	var users proto.User
	users.CreateTable(&db)
	return
}
