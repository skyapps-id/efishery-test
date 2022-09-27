package utils

import (
	"iot-service/utils/errors"

	"github.com/labstack/echo/v4"
)

type (
	MsContext struct {
		echo.Context
	}
	Success struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
		Meta    interface{} `json:"meta,omitempty"`
	}

	Failed struct {
		Success bool   `json:"success"`
		Code    int    `json:"code"`
		Message string `json:"message"`
		Error   string `json:"error"`
	}
)

func NewMsContext(parent echo.Context) *MsContext {
	return &MsContext{parent}
}

func (c *MsContext) Success(data interface{}, httpCode ...int) error {
	var code = 200
	if httpCode != nil {
		code = httpCode[0]
	}
	return c.JSON(code, Success{
		Success: true,
		Message: "success",
		Data:    data,
	})
}

func (c *MsContext) Fail(err error) error {
	var (
		ed = errors.ExtractError(err)
	)

	return c.JSON(ed.HttpCode, Failed{
		Success: false,
		Code:    ed.Code,
		Message: "failed",
		Error:   ed.Message,
	})
}
