package rest

import (
	"iot-service/internal_grpc"
	"iot-service/pkg"
	"iot-service/repository"
	"iot-service/server/rest/controller"
	"iot-service/service"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HandlerRest(db *gorm.DB, gRPC pkg.GRPC) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Controller
	healthCheck := controller.NewHealthCheckController()

	// Feed Logs Summary
	pondFeedersInternalGRPC := internal_grpc.NewPondFeedersInternalGRPC(gRPC)
	feedLogsRepository := repository.NewFeedLogsRepository(db)
	feedLogsSummaryService := service.NewPondFeedersService(pondFeedersInternalGRPC, feedLogsRepository)
	feedLogsSummaryController := controller.NewFeedLogsSummaryController(feedLogsSummaryService)

	// Routing
	e := echo.New()
	e.GET("/", healthCheck.HealthCheck)
	e.GET("/feedlogs-summary/:pondUuid/:date", feedLogsSummaryController.FeedLogsSummary)

	e.Logger.Fatal(e.Start(":" + port))
}
