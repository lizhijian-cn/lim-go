package grpc

import (
	"context"
	"github.com/aceld/zinx/zlog"
	"google.golang.org/grpc"
	"lim/conf"
	pb "lim/proto"
)

func StartGrpcClient() {
	conn, err := grpc.Dial(conf.RpcPort)
	if err != nil {
		zlog.Fatalf("failed to start grpc client: %s", err)
	}

	client := pb.NewInternalClient(conn)
	stream, err := client.Send(context.Background())
	if err != nil {
		zlog.Fatalf("failed to open stream: %v", err)
	}

	//ctx := stream.Context()
}
