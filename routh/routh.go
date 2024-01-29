package sample

import "github.com/labstack/echo/v4"

func Routh(e *echo.Echo, h Handler) {
	sample := e.Group("/v1/sample")
	sample.GET("", h.SampleHandler)
}
