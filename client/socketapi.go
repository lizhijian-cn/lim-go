package main

import (
	pb "lim/proto"
	"net"
)

type SocketApi struct {
	conn net.Conn
}

func (api *SocketApi) Connect() error {
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		return err
	}

	api.conn = conn
	msg := pb.Msg{
		SenderId: UserId,
		Type:     pb.Msg_Greet,
		Content:  "",
		Ok:       false,
	}
	SendMsg(api.conn, 0, &msg)
	return nil
}

func (api *SocketApi) Close() {
	api.conn.Close()
}

func (api *SocketApi) Chat(receiverId int64, content string) {
	msg := pb.Msg{
		SenderId:   UserId,
		ReceiverId: receiverId,
		Type:       pb.Msg_Chat,
		Content:    content,
		Ok:         false,
	}
	SendMsg(api.conn, 0, &msg)
}
