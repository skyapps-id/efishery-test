package messaging

import (
	"context"
	"iot-service/internal_grpc"
	"iot-service/pkg"
	"iot-service/repository"
	"iot-service/service"
	"log"

	"gorm.io/gorm"
)

func MessagingConsumer(rabbitMQ pkg.RabbitMQ, db *gorm.DB) {
	// Load Dependency
	feederGRPC := internal_grpc.NewFeederInternalGRPC()
	feederService := service.NewFeederService(feederGRPC)
	feedLogsRepository := repository.NewFeedLogsRepository(db)
	feedLogsService := service.NewFeedLogsService(feedLogsRepository)

	// Messaging Consumer Module
	feeders, err := feederService.Fetch(context.Background())
	if err != nil {
		log.Fatal(" -> Can't Fetch feeder devices, please run feeder service. or check connection")
	}

	messaging := NewFeederMessaging(rabbitMQ, feedLogsService)
	for _, record := range feeders {
		go func(topic string) {
			messaging.FeederConsumer(topic, "feedlogs")
		}(record.Barcode)
	}
}
