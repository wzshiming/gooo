package protocol

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	Id        int64
	Username  string `sql:"type:varchar(100);not null;"`
	Password  string `sql:"type:varchar(100);not null;"`
	Email     Email
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *User) CreateTable(db *gorm.DB) {
	db.CreateTable(s)
	db.Model(s).AddUniqueIndex("idx_user_username", "username")
}

func (s *User) DropTable(db *gorm.DB) {
	db.DropTable(s)
}

type Email struct {
	Id         int64
	UserId     int64
	Email      string `sql:"type:varchar(100);"` // Set field's type
	Subscribed bool
}
