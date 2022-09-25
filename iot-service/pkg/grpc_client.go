package pkg

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	GRPC struct {
		ClientConn *grpc.ClientConn
	}
)

func NewGrpcDial() GRPC {
	conn, err := grpc.Dial("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return GRPC{
		ClientConn: conn,
	}
}
