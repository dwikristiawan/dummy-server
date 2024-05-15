package routh

import (
	"mocking-server/internal/sample"

	"github.com/labstack/echo/v4"
)

func RouthSample(e *echo.Echo, h sample.Handler) {
	r := e.Group("/v1/sample")
	r.GET("", h.SampleHandler)
}
