package messaging

import (
	"iot-service/pkg"
	"log"
)

type (
	FeederMessaging interface {
		FeederConsumer(topic, queue string)
	}

	feederMessagingInst struct {
		rabbitMQ pkg.RabbitMQ
	}
)

func NewFeederMessaging(rabbitMQ pkg.RabbitMQ) FeederMessaging {
	return &feederMessagingInst{rabbitMQ: rabbitMQ}
}

func (m feederMessagingInst) FeederConsumer(topic, queue string) {
	barcode, msg, err := m.rabbitMQ.Consume(m.rabbitMQ, topic, queue)
	if err != nil {
		log.Fatal(err, "Failed to publish message")
	}

	for message := range msg {
		log.Printf(" >Barcode %s Received message: %s\n", *barcode, message.Body)
	}
}
