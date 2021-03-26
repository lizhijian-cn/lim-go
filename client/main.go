package main

import (
	"errors"
	"fmt"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
	"github.com/golang/protobuf/proto"
	"io"
	pb "lim/proto"
	"net"
	"os"
)

//func getOffline(conn net.Conn) error {
//	msg, err := GetMsg(conn)
//	if err != nil {
//		return err
//	}
//	offlineMsg := msg.(*pb.OfflineMsg)
//	for _, addRes := range offlineMsg.FriendAddingRes {
//		if addRes.Ok {
//			fmt.Printf("%s agree to add you as friend\n", addRes.SenderName)
//		} else {
//			fmt.Printf("%s disagree to add you as friend\n", addRes.SenderName)
//		}
//	}
//	fmt.Println()
//
//	fmt.Println("here are you friends now:")
//	for _, friend := range offlineMsg.Friend {
//		fmt.Println(friend.Username)
//	}
//	fmt.Println()
//
//	for _, add := range offlineMsg.FriendAdding {
//		fmt.Printf("%s want to add you as friend\n", add.SenderName)
//	}
//	fmt.Println()
//
//	fmt.Println("here are offline messages")
//	for _, msg := range offlineMsg.ChatMsg {
//		fmt.Printf("%s: %s\n", msg.SenderId)
//	}
//}
func main() {
	args := os.Args
	httpApi := &HttpApi{}
	socketApi := &SocketApi{}

	_, err := httpApi.Login(args[1], args[2])
	if err != nil {
		zlog.Fatalf("failed to login: %s", err)
		return
	}

	httpApi.SyncMsg()

	if err = socketApi.Connect(); err != nil {
		zlog.Fatal("failed to connect chat server: %s", err)
		return
	}
	defer socketApi.Close()

	// reader
	go func() {
		for {
			msg, err := GetMsg(socketApi.conn)
			if err != nil {
				zlog.Fatalf("failed to get message: %s", err)
				break
			}
			switch msg.Type {
			case pb.Msg_Chat:
				fmt.Printf("%s: %s\n", IdFriendMap[msg.SenderId], msg.Content)
			case pb.Msg_Adding:
				// TODO
			}
		}
	}()
	for {
		var in string
		fmt.Scan(&in)
		if in == "quit" {
			break
		} else if in == "send" {
			var name, content string
			fmt.Scanln(&name)
			fmt.Scanf("%q", &content)
			if receiverId, ok := FriendIdMap[name]; ok {
				fmt.Println(content)
				socketApi.Chat(receiverId, content)
			} else {
				fmt.Printf("%s is not your friend\n", name)
			}
		} else if in == "reply" {
			//var name, ok string
			// TODO
		}
	}
}

func SendMsg(conn net.Conn, msgId uint32, data proto.Message) {
	msg, err := proto.Marshal(data)
	if err != nil {
		zlog.Fatalf("failed to marshal: %v", data)
		return
	}
	dp := znet.NewDataPack()
	pack, _ := dp.Pack(znet.NewMsgPackage(msgId, msg))
	if _, err = conn.Write(pack); err != nil {
		zlog.Fatalf("failed to write to conn %s: %s\n", conn.LocalAddr(), err)
		return
	}
}

//
//func GetMsg(conn net.Conn) (proto.Message, error) {
//	data, err := GetData(conn)
//	if err != nil {
//		return nil, err
//	}
//
//	switch data.Id {
//	case router.FriendAddingId:
//		msg := &pb.FriendAdding{}
//		if err := proto.Unmarshal(data.Data, msg); err != nil {
//			return nil, err
//		}
//		return msg, nil
//	case router.FriendAddingResId:
//		msg := &pb.FriendAddingRes{}
//		if err := proto.Unmarshal(data.Data, msg); err != nil {
//			return nil, err
//		}
//		return msg, nil
//	case router.ChatId:
//		msg := &pb.ChatMsg{}
//		if err := proto.Unmarshal(data.Data, msg); err != nil {
//			return nil, err
//		}
//		return msg, nil
//	case router.OfflineId:
//		msg := &pb.OfflineMsg{}
//		if err := proto.Unmarshal(data.Data, msg); err != nil {
//			return nil, err
//		}
//		return msg, nil
//	default:
//		return nil, errors.New("invalid msgId")
//	}
//}

func GetMsg(conn net.Conn) (*pb.Msg, error) {
	dp := znet.NewDataPack()
	headLen := dp.GetHeadLen()
	headData := make([]byte, headLen)
	if _, err := io.ReadFull(conn, headData); err != nil {
		return nil, err
	}

	msgHead, err := dp.Unpack(headData)
	if err != nil {
		return nil, err
	}

	if msgHead.GetDataLen() > 0 {
		message := msgHead.(*znet.Message)
		message.Data = make([]byte, msgHead.GetDataLen())
		if _, err := io.ReadFull(conn, message.Data); err != nil {
			return nil, err
		}
		msg := &pb.Msg{}
		if err := proto.Unmarshal(message.Data, msg); err != nil {
			return nil, err
		}
		return msg, nil
	}
	return nil, errors.New("len of head <= 0")
}
