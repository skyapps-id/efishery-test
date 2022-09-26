package main

import (
	"feeder-generator/messaging"
	"feeder-generator/pkg"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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
		io.WriteString(w, "ok!")
	}

	sendfeeder := func(w http.ResponseWriter, req *http.Request) {
		barcode := strings.Split(req.URL.Path, "/")[2]
		byte, _ := os.ReadFile(fmt.Sprintf("./json/%s.json", barcode))

		// Publish
		publish := messaging.NewFeederMessaging(rabbitMQ)
		err := publish.FeederPublish(barcode, string(byte))
		if err != nil {
			log.Fatal(err, "Failed to publish message")
		}

		io.WriteString(w, "success!")
	}

	http.HandleFunc("/health-check", healthCheck)
	http.HandleFunc("/send-feeder/00001-AL03005090R-SMIT", sendfeeder)
	http.HandleFunc("/send-feeder/00002-AL03005090R-F3ot", sendfeeder)
	http.HandleFunc("/send-feeder/00005-AL15005090R-dxbm", sendfeeder)

	log.Println("Listing for Rest API " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
