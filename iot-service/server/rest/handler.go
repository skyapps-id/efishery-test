package rest

import (
	"iot-service/internal_grpc"
	"iot-service/pkg"
	"iot-service/repository"
	"iot-service/server/rest/controller"
	"iot-service/service"
	"log"
	"net/http"
	"os"

	"gorm.io/gorm"
)

func HandlerRest(db *gorm.DB) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Common
	gRPC := pkg.NewGrpcDial()

	// Controller
	healthCheck := controller.NewHealthCheckController()

	// Feed Logs Summary
	pondFeedersInternalGRPC := internal_grpc.NewPondFeedersInternalGRPC(gRPC)
	feedLogsRepository := repository.NewFeedLogsRepository(db)
	feedLogsSummaryService := service.NewPondFeedersService(pondFeedersInternalGRPC, feedLogsRepository)
	feedLogsSummaryController := controller.NewFeedLogsSummaryController(feedLogsSummaryService)

	// Routing
	http.HandleFunc("/", healthCheck.HealthCheck)
	http.HandleFunc("/feedlogs-summary", feedLogsSummaryController.FeedLogsSummary)

	log.Println("Listing for Rest API " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
