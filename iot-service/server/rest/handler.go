package rest

import (
	"iot-service/pkg"
	"iot-service/server/rest/controller"
	"iot-service/service"
	"log"
	"net/http"
	"os"
)

func HandlerRest() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Common
	gRPC := pkg.NewGrpcDial()

	// Controller
	healthCheck := controller.NewHealthCheckController()

	// Chat
	chatService := service.NewChatService(gRPC)
	chatController := controller.NewChatCheckController(chatService)

	http.HandleFunc("/", healthCheck.HealthCheck)
	http.HandleFunc("/chat", chatController.Chat)

	log.Println("Listing for Rest API " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
