package protocol

import (
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

type Email struct {
	Id         int64
	UserId     int64
	Email      string `sql:"type:varchar(100);"` // Set field's type
	Subscribed bool
}
