package chatserver

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
	"lim/chat/router"
	"lim/userstatus"
)

func StartChatServer() {
	server := znet.NewServer()
	server.AddRouter(0, &router.Router{})

	server.SetOnConnStop(func(connection ziface.IConnection) {
		userId, err := connection.GetProperty("userId")
		if err != nil {
			zlog.Fatalf("failed to get property userId: %s", err)
		}
		userstatus.Offline(userId.(int64))
	})
	server.Serve()
}
