package messaging

import (
	"context"
	"encoding/json"
	"iot-service/dto"
	"iot-service/pkg"
	"iot-service/service"
	"log"
)

type (
	FeederMessaging interface {
		FeederConsumer(topic, queue string)
	}

	feederMessagingInst struct {
		rabbitMQ        pkg.RabbitMQ
		feedLogsService service.FeedLogsService
	}
)

func NewFeederMessaging(rabbitMQ pkg.RabbitMQ, feedLogsService service.FeedLogsService) FeederMessaging {
	return &feederMessagingInst{rabbitMQ: rabbitMQ, feedLogsService: feedLogsService}
}

func (m feederMessagingInst) FeederConsumer(topic, queue string) {
	barcode, msg, err := m.rabbitMQ.Consume(m.rabbitMQ, topic, queue)
	if err != nil {
		log.Fatal(err, "Failed to publish message")
	}

	for message := range msg {
		log.Printf(" > Barcode %s Received message\n", *barcode)

		feederMessaging := dto.FeederMessaging{}
		json.Unmarshal([]byte(message.Body), &feederMessaging)
		dataToString, _ := json.Marshal(feederMessaging.Data)

		var outputGrCount = float64(0)
		for _, val := range feederMessaging.Data {
			outputGrCount += val.OutputGr
		}
		_, err = m.feedLogsService.Create(context.Background(), dto.FeederMessagingRequest{
			Barcode:       *barcode,
			ScheduleUUID:  feederMessaging.ScheduleUUID,
			Data:          string(dataToString),
			DataCount:     len(feederMessaging.Data),
			OutputGrCount: outputGrCount,
		})
		if err != nil {
			log.Println(err)
		}
	}
}
