package controller

import (
	"io"
	"net/http"
)

type (
	HealthCheckController struct {
	}
)

func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}
func (c *HealthCheckController) HealthCheck(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "ok!")
}
