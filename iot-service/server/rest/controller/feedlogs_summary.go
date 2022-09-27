package controller

import (
	"errors"
	"iot-service/internal_const"
	"iot-service/service"
	"iot-service/utils"
	"time"

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
	var (
		pondUuid = ctx.Param("pondUuid")
		date     = ctx.Param("date")
		msctx    = utils.NewMsContext(ctx)
	)

	// Validation Date
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return msctx.Fail(internal_const.ErrBadRequest(errors.New("invalid format date")))
	}

	data, err := c.services.Fatch(msctx.Request().Context(), pondUuid, date)
	if err != nil {
		return msctx.Fail(err)
	}

	return msctx.Success(data)
}
