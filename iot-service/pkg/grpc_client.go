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

func NewGrpcDial(Host string) GRPC {
	conn, err := grpc.Dial(Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return GRPC{
		ClientConn: conn,
	}
}
