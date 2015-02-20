package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"gooo"
	"gooo/protocol"
)

var db gorm.DB

func dbconn(us gooo.DataBaseConfig) (err error) {
	db, err = gorm.Open(us.Dialect, us.Source)
	var users protocol.User
	users.CreateTable(&db)
	return
}
