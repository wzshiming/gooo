package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"log"
	"time"
)

func CheckErr(i interface{}) {
	if i != nil {
		log.Println(i)
	}
}

func main2() {

	//db := sqldb.NewSqlDB("yugioh", "YuGiOh2015", "mohegame.com:39979", "YuGiOh")
	db, _ := gorm.Open("mysql", "yugioh:YuGiOh2015@tcp(mohegame.com:39979)/YuGiOh?charset=utf8&parseTime=True")
	//log.Println(db.Find(&User{}))
	//db.CreateTable(&User{})
	//db.Model(&User{}).AddUniqueIndex("idx_user_username", "username")
	user := User{
		Username: "hh1",
		Password: "pss2",
	}

	//db.NewRecord(user) // => true
	err := db.Create(&user).Error
	CheckErr(err)
	db.Find(&user)
	log.Println(user)
	db.First(&user)
	log.Println(user)

	var users []User

	db.Find(&users)

	log.Println(users)
}

type User struct {
	Id        int64
	Username  string `sql:"type:varchar(100);not null;"`
	Password  string `sql:"type:varchar(100);not null;"`
	Email     string `sql:"type:varchar(100);not null;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	main2()
}
