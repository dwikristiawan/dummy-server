package routh

import (
	"mocking-server/internal/auth"
	mockserver "mocking-server/internal/rest/mock_server"
	"mocking-server/internal/security"

	"github.com/labstack/echo/v4"
)

func RouthAuth(e *echo.Echo, h auth.Handler, middleware security.MiddlewareService) {
	r := e.Group("mockers/v1/auth")
	r.POST("/register", h.RegisterHandler)
	r.POST("/login", h.LoginHandler)
	r.GET("/refresh-access", h.RefreshTokenHandler)
	r.GET("/hello", func(c echo.Context) error { return c.JSON(200, "hello") }, middleware.MiddlewareSecurity(nil))
}

func MockServerRouth(e *echo.Echo, h mockserver.Handler, midleware security.MiddlewareService) {
	r := e.Group("mockers/v1/mock")
	r.POST("", h.SetMockDataHandler, midleware.MiddlewareSecurity(nil))
	x := e.Group("")
	x.Any("/*", h.MatchMockHandler)
}
