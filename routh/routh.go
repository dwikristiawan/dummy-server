package routh

import (
	"mocking-server/internal/auth"

	"github.com/labstack/echo/v4"
)

func RouthAuth(e *echo.Echo, h auth.Handler) {
	r := e.Group("mockers/v1/auth")
	r.POST("/register", h.RegisterHandler)
	r.POST("/login", h.LoginHandler)
}
