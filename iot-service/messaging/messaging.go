package messaging

import (
	"iot-service/pkg"
	"iot-service/service"
)

func MessagingConsumer(rabbitMQ pkg.RabbitMQ, feedLogsService service.FeedLogsService) {
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
