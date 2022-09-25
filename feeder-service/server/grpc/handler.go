package grpc

import (
	"feeder-service/server/grpc/pb"
	"feeder-service/server/grpc/pb/chat"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func HandlerGrpc() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9001"
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &pb.Server{})

	log.Println("Listing for gRPC " + port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
