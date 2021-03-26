package service

import (
	"errors"
	"lim/http/models"
)

func GetFriends(userId int64) []int64 {
	return models.GetFriends(userId)
}

func AddFriend(aid, bid int64) error {
	auser, buser := models.FindUserById(aid), models.FindUserById(bid)
	if auser.Id == 0 || buser.Id == 0 {
		return errors.New("no such id")
	}
	if models.IsFriend(aid, bid) {
		return errors.New("add friend repeatedly")
	}
	return models.AddFriend(aid, bid)
}
