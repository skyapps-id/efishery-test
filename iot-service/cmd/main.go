package main

import (
	"iot-service/database"
	"iot-service/messaging"
	"iot-service/pkg"
	"iot-service/server/rest"
	"log"

	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	gRPC pkg.GRPC
)

func init() {
	var err error
	db, err = database.Database()
	if err != nil {
		log.Println(err)
	}

	gRPC = pkg.NewGrpcDial()

}

func main() {
	// Connection RabbitMQ
	rabbitMQ := pkg.NewRabbitMQ()
	defer rabbitMQ.Connection.Close()
	defer rabbitMQ.Chanel.Close()

	// Init module messaging
	messaging.MessagingConsumer(rabbitMQ, db, gRPC)

	rest.HandlerRest(db, gRPC)
	log.Println("exit complate!")
}
