package service

import (
	"errors"
	"lim/http/models"
)

func Register(username, password string) error {
	//user := models.FindUser(username)
	//
	//if user != nil {
	//	return errors.New(fmt.Sprintf("%s has been registered", username))
	//}
	//_ = models.CreateUser(username, password)
	return nil
}

func Login(username, password string) (*models.User, error) {
	user := models.FindUser(username)
	if user.Id == 0 || user.Password != password {
		return nil, errors.New("password incorrect")
	}
	return &user, nil
}

func GetUser(userId int64) (*models.User, error) {
	user := models.FindUserById(userId)
	if user.Id == 0 {
		return nil, errors.New("invalid user id")
	}
	return &user, nil
}
