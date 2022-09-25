package main

import (
	"feeder-generator/messaging"
	"feeder-generator/pkg"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Messaging Define
	rabbitMQ := pkg.NewRabbitMQ()
	defer rabbitMQ.Connection.Close()
	defer rabbitMQ.Chanel.Close()

	// Init module messaging
	messaging.MessagingConsumer(rabbitMQ)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	healthCheck := func(w http.ResponseWriter, req *http.Request) {
		// Publish RabbitMQ
		barcode := "00002-AL03005090R-F3ot"
		body := "Hi halovina, keep in touch"
		publish := messaging.NewFeederMessaging(rabbitMQ)
		err := publish.FeederPublish(barcode, body)
		if err != nil {
			log.Fatal(err, "Failed to publish message")
		}

		io.WriteString(w, "ok!")
	}
	http.HandleFunc("/", healthCheck)

	log.Println("Listing for Rest API " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
