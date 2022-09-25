package pkg

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type (
	RabbitMQ struct {
		Connection *amqp.Connection
		Chanel     *amqp.Channel
	}
)

func NewRabbitMQ() RabbitMQ {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err, "Failed to connect rabbitmq")
	}
	log.Println("Connecting to rabbitmq success")

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err, "Failed to open a channel")
	}

	return RabbitMQ{
		Connection: conn,
		Chanel:     ch,
	}
}

func (i RabbitMQ) Publish(rabbitMQ RabbitMQ, topic, queue, payload string) error {
	defineTopic := fmt.Sprintf("/cobox/%s/feedlogs", topic)

	// We create a topic
	if err := rabbitMQ.Chanel.ExchangeDeclare(defineTopic, "topic", true, false, false, false, nil); err != nil {
		panic("error declaring the topic: " + err.Error())
	}

	body := payload
	err := rabbitMQ.Chanel.Publish(
		defineTopic, // exchange
		topic,
		false, // mandatory
		false, // immadiate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return err
	}

	return nil
}

func (i RabbitMQ) Consume(rabbitMQ RabbitMQ, topic, queue string) (*string, <-chan amqp.Delivery, error) {
	defineTopic := fmt.Sprintf("/cobox/%s/feedlogs", topic)
	defineQueue := fmt.Sprintf("%s/%s", topic, queue)

	// We create a topic
	if err := rabbitMQ.Chanel.ExchangeDeclare(defineTopic, "topic", true, false, false, false, nil); err != nil {
		return nil, nil, err
	}

	// We create a queue
	if _, err := rabbitMQ.Chanel.QueueDeclare(defineQueue, true, false, false, false, nil); err != nil {
		return nil, nil, err
	}

	// We create a bind
	if err := rabbitMQ.Chanel.QueueBind(defineQueue, "#", defineTopic, false, nil); err != nil {
		return nil, nil, err
	}

	msg, err := rabbitMQ.Chanel.Consume(
		defineQueue, // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		return nil, nil, err
	}

	return &topic, msg, nil
}
