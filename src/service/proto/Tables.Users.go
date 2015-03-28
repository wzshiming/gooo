package proto

import (
	"rego/dbs"
	"time"
)

type User struct {
	Id        uint64
	Username  string `sql:"type:varchar(64);not null;"`
	Password  string `sql:"type:varchar(64);not null;"`
	Email     Email
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *User) CreateTable(db *dbs.DB) {
	db.CreateTable(s)
	db.Model(s).AddUniqueIndex("idx_user_username", "username")
}

func (s *User) DropTable(db *dbs.DB) {
	db.DropTable(s)
}

type Email struct {
	Id         uint64
	UserId     uint64
	Email      string `sql:"type:varchar(64);"` // Set field's type
	Subscribed bool
}