package messaging

import (
	"iot-service/pkg"
	"iot-service/repository"
	"iot-service/service"

	"gorm.io/gorm"
)

func MessagingConsumer(rabbitMQ pkg.RabbitMQ, db *gorm.DB) {
	// Load Dependency
	feedLogsRepository := repository.NewFeedLogsRepository(db)
	feedLogsService := service.NewFeedLogsService(feedLogsRepository)

	// Messaging Consumer Module
	barcodes := []string{
		"00001-AL03005090R-SMIT",
		"00002-AL03005090R-F3ot",
		"00005-AL15005090R-dxbm",
	}

	messaging := NewFeederMessaging(rabbitMQ, feedLogsService)
	for _, barcode := range barcodes {
		go func(topic string) {
			messaging.FeederConsumer(topic, "feedlogs")
		}(barcode)
	}
}
