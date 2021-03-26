package router

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/golang/protobuf/proto"
)

func SendMsg(conn ziface.IConnection, msgId uint32, data proto.Message) {
	msg, err := proto.Marshal(data)
	if err != nil {
		zlog.Fatalf("failed to marshal: %s", err)
		return
	}

	if err := conn.SendMsg(msgId, msg); err != nil {
		zlog.Fatalf("failed to send msg to %s: %s", conn.RemoteAddr(), err)
	}
}
