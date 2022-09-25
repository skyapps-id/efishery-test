package messaging

import "iot-service/pkg"

func MessagingConsumer(rabbitMQ pkg.RabbitMQ) {
	// Messaging Consumer Module
	barcodes := []string{
		"00001-AL03005090R-SMIT",
		"00002-AL03005090R-F3ot",
		"00005-AL15005090R-dxbm",
	}

	messaging := NewFeederMessaging(rabbitMQ)
	for _, barcode := range barcodes {
		go func(topic string) {
			messaging.FeederConsumer(topic, "feedlogs")
		}(barcode)
	}
}
