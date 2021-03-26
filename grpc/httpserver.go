package grpc

import (
	"github.com/aceld/zinx/zlog"
	"google.golang.org/grpc"
	"io"
	"lim/conf"
	pb "lim/proto"
	"net"
)

var (
	AddingFriendChan = make(chan *pb.Msg)
)

type Server struct {
	pb.UnimplementedInternalServer
}

func (s *Server) Send(server pb.Internal_SendServer) error {
	ctx := server.Context()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-AddingFriendChan:
			server.Send(msg)
		}

		msg, err := server.Recv()
		if err == io.EOF {
			zlog.Info("grpc server closed")
			return nil
		}
		if err != nil {
			zlog.Fatalf("grpc server received error: %s", err)
			continue
		}

		// offline chat msg
		switch msg.Type {
		case pb.Msg_Chat:
		// TODO

		default:
			panic("unexpected msg")
		}
	}
}

func StartGrpcServer() {
	l, err := net.Listen("tcp", conf.RpcPort)
	if err != nil {
		zlog.Fatalf("failed to start grpc client: %s", err)
		return
	}

	server := grpc.NewServer()
	pb.RegisterInternalServer(server, &Server{})
	err = server.Serve(l)
	if err != nil {
		zlog.Fatalf("failed to serve: %s", err)
	}
}
