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
	db *gorm.DB
)

func init() {
	var err error
	db, err = database.Database()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	// Connection RabbitMQ
	rabbitMQ := pkg.NewRabbitMQ()
	defer rabbitMQ.Connection.Close()
	defer rabbitMQ.Chanel.Close()

	// Init module messaging
	messaging.MessagingConsumer(rabbitMQ, db)

	rest.HandlerRest(db)
	log.Println("exit complate!")
}
