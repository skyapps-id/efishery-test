package controller

import (
	"io"
	"iot-service/service"
	"net/http"
)

type (
	ChatCheckController struct {
		services service.ChatService
	}
)

func NewChatCheckController(services service.ChatService) *ChatCheckController {
	return &ChatCheckController{services: services}
}
func (c *ChatCheckController) Chat(w http.ResponseWriter, req *http.Request) {
	data, err := c.services.Fatch()
	if err != nil {
		io.WriteString(w, "error!")

	}
	io.WriteString(w, *data)
}
