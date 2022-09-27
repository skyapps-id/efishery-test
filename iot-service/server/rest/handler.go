package rest

import (
	"iot-service/internal_grpc"
	"iot-service/repository"
	"iot-service/server/rest/controller"
	"iot-service/service"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HandlerRest(db *gorm.DB) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Controller
	healthCheck := controller.NewHealthCheckController()

	// Feed Logs Summary
	pondFeedersInternalGRPC := internal_grpc.NewPondFeedersInternalGRPC()
	feedLogsRepository := repository.NewFeedLogsRepository(db)
	feedLogsSummaryService := service.NewPondFeedersService(pondFeedersInternalGRPC, feedLogsRepository)
	feedLogsSummaryController := controller.NewFeedLogsSummaryController(feedLogsSummaryService)

	// Routing
	e := echo.New()
	e.GET("/health-check", healthCheck.HealthCheck)
	e.GET("/feedlogs-summary/:pondUuid/:date", feedLogsSummaryController.FeedLogsSummary)

	e.Logger.Fatal(e.Start(":" + port))
}
