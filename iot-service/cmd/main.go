package main

import (
	"iot-service/database"
	"iot-service/messaging"
	"iot-service/pkg"
	"iot-service/repository"
	"iot-service/server/rest"
	"iot-service/service"
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

	feedLogsRepository := repository.NewFeedLogsRepository(db)
	feedLogsService := service.NewFeedLogsService(feedLogsRepository)

	// Init module messaging
	messaging.MessagingConsumer(rabbitMQ, feedLogsService)

	rest.HandlerRest()
	log.Println("exit complate!")
}
