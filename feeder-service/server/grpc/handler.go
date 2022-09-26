package grpc

import (
	"feeder-service/repository"
	"feeder-service/server/grpc/pb"
	"feeder-service/server/grpc/pb/pond"
	"feeder-service/service"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func HandlerGrpc(db *gorm.DB) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9001"
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Pond
	pondRepository := repository.NewPondRepository(db)
	pondService := service.NewPondService(pondRepository)

	pond.RegisterPondServiceServer(grpcServer, &pb.PondGRPC{
		Services: pondService,
	})

	log.Println("Listing for gRPC " + port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
