package chat

import (
	"google.golang.org/grpc"
	pb "lim/internal/proto"
	"log"
	"net"
	"sync"
)

type Server struct {
	pb.UnimplementedChatServiceServer
}

var clients sync.Map
var offline sync.Map

func (s *Server) Chat(stream pb.ChatService_ChatServer) error {
	ctx := stream.Context()

	for {
		select {
		case <-ctx.Done():
			log.Println("client close stream")
			return ctx.Err()
		default:
			r, err := stream.Recv()
			if err != nil {
				return err
			}

			switch r.MsgType {
			case pb.MessageType_LOGIN:
				// read offline
				receiveChan := make(chan *pb.Message, 10)
				clients.Store(r.Sender, receiveChan)
				go func() {
					for {
						select {
						case <-ctx.Done():
							return
						case msg := <-receiveChan:
							stream.Send(msg)
						}
					}
				}()
			case pb.MessageType_HEARTBEAT:

			case pb.MessageType_TEXT:
				if ch, ok := clients.Load(r.Receiver); ok {
					ch.(chan *pb.Message) <- r
				} else {
					// add to offline
				}
			}
		}
	}
}

func StartChatServer() {
	l, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return
	}
	server := grpc.NewServer()
	pb.RegisterChatServiceServer(server, &Server{})
	if err := server.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}

	log.Printf("chat server start\n")
}
