package main

import (
	"feeder-service/server/grpc"
	"feeder-service/server/rest"
	"log"
	"sync"
)

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
		grpc.HandlerGrpc()
		wg.Done()
	}()

	wg.Wait()
	log.Println("exit complate!")
}
