package models

import (
	"errors"
	"github.com/go-xorm/xorm"
	"lim/db"
)

type Relation struct {
	Id  int64 `xorm:"pk autoincr"`
	Aid int64
	Bid int64
}

func AddFriend(aid, bid int64) error {
	if aid == bid {
		return errors.New("can't add yourself")
	}
	if aid > bid {
		return AddFriend(bid, aid)
	}
	relation := &Relation{
		Aid: aid,
		Bid: bid,
	}
	_, err := relationEngine().Insert(relation)
	return err
}

func GetFriends(userId int64) []int64 {
	a, b := make([]int64, 0), make([]int64, 0)
	err := relationEngine().Where("aid = ?", userId).Cols("bid").Find(&a)
	if err != nil {
		db.Engine.Logger().Errorf("get friends err: %s", err)
		return []int64{}
	}
	err = relationEngine().Where("bid = ?", userId).Cols("aid").Find(&b)
	if err != nil {
		db.Engine.Logger().Errorf("get friends err: %s", err)
		return []int64{}
	}
	return append(a, b...)
}

func IsFriend(aid, bid int64) bool {
	if aid > bid {
		return IsFriend(bid, aid)
	}
	has, err := relationEngine().Where("aid = ? and bid = ?", aid, bid).Exist()
	if err != nil {
		db.Engine.Logger().Errorf("is friend err: %s", err)
		return false
	}
	return has
}

func relationEngine() *xorm.Session {
	return db.Engine.Table("Relation")
}
