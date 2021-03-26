package userstatus

import (
	"github.com/aceld/zinx/ziface"
	"sync"
)

var (
	userStatus = sync.Map{}
)

func Online(userId int64, conn ziface.IConnection) {
	userStatus.Store(userId, conn)
}

func Offline(userId int64) {
	userStatus.Delete(userId)
}

func GetConn(userId int64) ziface.IConnection {
	if conn, ok := userStatus.Load(userId); ok {
		return conn.(ziface.IConnection)
	}
	panic("invalid userid")
}
