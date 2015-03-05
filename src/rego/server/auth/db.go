package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"rego/cfg"
	"rego/server/proto"
)

var db gorm.DB

func dbconn(us cfg.DbConfig) (err error) {
	db, err = gorm.Open(us.Dialect, us.Source)
	var users proto.User
	users.CreateTable(&db)
	return
}
