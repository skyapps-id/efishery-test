package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	HealthCheckController struct {
	}
)

func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}
func (c *HealthCheckController) HealthCheck(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "ok!")
}
