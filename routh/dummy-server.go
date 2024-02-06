package routh

import (
	"dummy-server/module/dummyServer"

	"github.com/labstack/echo/v4"
)

func RouthDummyServer(e *echo.Echo, h dummyServer.Handler) {
	r := e.Group("/v1/mock-server")
	r.POST("", h.SetDummyHandler)
	r.DELETE("/:id", h.RemoveDummyHandler)
	r.GET("", h.GetAllDummyHandler)
	r.GET("/:id", h.GetDummyByIdHandler)
	e.Any("*", h.MatchingDummyHandler)

}
