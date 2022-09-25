package messaging

import (
	"feeder-generator/pkg"
	"fmt"
	"log"
)

type (
	FeederMessaging interface {
		FeederPublish(barcode, payload string) error
	}

	feederMessagingInst struct {
		rabbitMQ pkg.RabbitMQ
	}
)

func NewFeederMessaging(rabbitMQ pkg.RabbitMQ) FeederMessaging {
	return &feederMessagingInst{rabbitMQ: rabbitMQ}
}

func (m feederMessagingInst) FeederPublish(barcode, payload string) error {
	err := m.rabbitMQ.Publish(m.rabbitMQ, barcode, payload)
	if err != nil {
		return err
	}
	log.Printf("Sending message success: %s", payload)
	fmt.Println("Helo")
	return nil
}
