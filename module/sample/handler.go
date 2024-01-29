package sample

import (
	"context"

	"github.com/labstack/echo/v4"
)

type handler struct {
	controller Controller
}

func NewHandler(controller Controller) Handler {
	return &handler{controller: controller}
}

type Handler interface {
	SampleHandler(echo.Context) error
}

func (h handler) SampleHandler(e echo.Context) error {
	response := h.controller.SampleController(context.Background())
	return e.JSON(response.ResponseCode, response)
}
