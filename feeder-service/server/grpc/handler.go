package grpc

import (
	"feeder-service/repository"
	"feeder-service/server/grpc/pb"
	"feeder-service/server/grpc/pb/feeder"
	"feeder-service/server/grpc/pb/pond"
	"feeder-service/server/grpc/pb/pond_feeders"
	"feeder-service/service"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func HandlerGrpc(db *gorm.DB) {
	port := os.Getenv("GRPC_PORT")
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

	// Feeder
	feederRepository := repository.NewFeederRepository(db)
	feederService := service.NewFeederService(feederRepository)
	feeder.RegisterFeederServiceServer(grpcServer, &pb.FeederGRPC{
		Services: feederService,
	})

	// PondFeeders
	pondfeedersRepository := repository.NewPondFeederRepository(db)
	pondfeedersService := service.NewPondFeederService(pondfeedersRepository)
	pond_feeders.RegisterPondFeedersServiceServer(grpcServer, &pb.PondFeedersGRPC{
		Services: pondfeedersService,
	})

	log.Println("Listing for gRPC " + port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
