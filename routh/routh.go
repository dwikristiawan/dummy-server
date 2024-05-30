package routh

import (
	"mocking-server/internal/auth"
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
