package main

import (
	"feeder-service/database"
	"feeder-service/server/grpc"
	"feeder-service/server/rest"
	"log"
	"sync"

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
	// Init Server
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// Run REst API
	go func() {
		rest.HandlerRest()
		wg.Done()
	}()
	// Run gRPC
	go func() {
		grpc.HandlerGrpc(db)
		wg.Done()
	}()

	wg.Wait()
	log.Println("exit complate!")
}
