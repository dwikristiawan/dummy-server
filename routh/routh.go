package routh

import (
	"mocking-server/internal/auth"

	"github.com/labstack/echo/v4"
)

// func RouthDummyServer(e *echo.Echo, h dummyServer.Handler) {
// 	r := e.Group("/v1/mock-server")
// 	r.POST("", h.SetDummyHandler)
// 	r.DELETE("/:id", h.RemoveDummyHandler)
// 	r.GET("", h.GetAllDummyHandler)
// 	r.GET("/:id", h.GetDummyByIdHandler)
// 	e.Any("*", h.MatchingDummyHandler)

// }

func RouthProfile(e *echo.Echo, h auth.AuthHandler) {
	r := e.Group("mockers/v1/auth")
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
}
