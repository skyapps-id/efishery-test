package rest

import (
	"feeder-service/server/rest/controller"
	"log"
	"net/http"
	"os"
)

func HandlerRest() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	// Controller
	healthCheck := controller.NewHealthCheckController()

	http.HandleFunc("/health-check", healthCheck.HealthCheck)

	log.Println("Listing for Rest API " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
