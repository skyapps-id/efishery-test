package main

import (
	"iot-service/messaging"
	"iot-service/pkg"
	"iot-service/server/rest"
	"log"
)

func main() {
	// Connection RabbitMQ
	rabbitMQ := pkg.NewRabbitMQ()
	defer rabbitMQ.Connection.Close()
	defer rabbitMQ.Chanel.Close()

	// Init module messaging
	messaging.MessagingConsumer(rabbitMQ)

	rest.HandlerRest()
	log.Println("exit complate!")
}
