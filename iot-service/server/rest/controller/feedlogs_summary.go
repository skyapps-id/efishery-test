package controller

import (
	"context"
	"iot-service/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	FeedLogsSummaryController struct {
		services service.PondFeedersService
	}
)

func NewFeedLogsSummaryController(services service.PondFeedersService) *FeedLogsSummaryController {
	return &FeedLogsSummaryController{services: services}
}
func (c *FeedLogsSummaryController) FeedLogsSummary(ctx echo.Context) error {
	data, err := c.services.Fatch(context.Background(), "26f1b9ee-65d9-4c7d-afb1-f7137fefa784")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, data)
}
