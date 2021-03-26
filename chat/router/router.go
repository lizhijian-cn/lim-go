package router

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
	"github.com/golang/protobuf/proto"
	pb "lim/proto"
	"lim/userstatus"
)

type Router struct {
	znet.BaseRouter
}

func (r *Router) Handle(req ziface.IRequest) {
	zlog.Infof("router handle ConnId = %d", req.GetConnection().GetConnID())
	msg := pb.Msg{}
	proto.Unmarshal(req.GetData(), &msg)

	switch msg.Type {
	case pb.Msg_Greet:
		userstatus.Online(msg.SenderId, req.GetConnection())
		req.GetConnection().SetProperty("userId", msg.SenderId)
	case pb.Msg_Chat:
		receiverConn := userstatus.GetConn(msg.ReceiverId)
		receiverConn.SendMsg(0, req.GetData())
	}
}
