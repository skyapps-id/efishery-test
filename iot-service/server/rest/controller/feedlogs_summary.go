package controller

import (
	"context"
	"io"
	"iot-service/service"
	"net/http"
)

type (
	FeedLogsSummaryController struct {
		services service.PondFeedersService
	}
)

func NewFeedLogsSummaryController(services service.PondFeedersService) *FeedLogsSummaryController {
	return &FeedLogsSummaryController{services: services}
}
func (c *FeedLogsSummaryController) FeedLogsSummary(w http.ResponseWriter, req *http.Request) {
	data, err := c.services.Fatch(context.Background(), "26f1b9ee-65d9-4c7d-afb1-f7137fefa784")
	if err != nil {
		io.WriteString(w, "error!")

	}
	io.WriteString(w, *data)
}
