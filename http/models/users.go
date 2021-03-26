package models

import (
	"github.com/go-xorm/xorm"
	"lim/db"
)

//type User struct {
//	Id       int64  `json:"id"`
//	Username string `json:"username"`
//	Password string `json:"password"`
//}

type User struct {
	Id       int64  `xorm:"pk autoincr"`
	Username string `xorm:"varchar(25) not null unique 'username'"`
	Password string `xorm:"varchar(25) not null 'password'"`
}

func FindUser(username string) User {
	user := User{}
	userEngine().Where("User.username = ?", username).Get(&user)
	return user
}

func FindUserById(id int64) User {
	user := User{}
	userEngine().ID(id).Get(&user)
	return user
}

//func CreateUser(username, password string) int64 {
//	Users[UserId] = User{UserId, username, password}
//	oldId := UserId
//	UserId++
//	return oldId
//}

func userEngine() *xorm.Session {
	return db.Engine.Table("User")
}
