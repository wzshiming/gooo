package proto

import "time"

type User struct {
	Id        uint64 `gorm:"primary_key"`
	Username  string `sql:"type:varchar(64);not null;unique_index"`
	Password  string `sql:"type:varchar(64);not null;"`
	Email     Email
	CreatedAt time.Time
	UpdatedAt time.Time
	//DeletedAt *time.Time
}

type Email struct {
	Id         uint64 `gorm:"primary_key"`
	UserId     uint64 `sql:"index"`
	Email      string `sql:"type:varchar(100);unique_index"` // Set field's type
	Subscribed bool
}
