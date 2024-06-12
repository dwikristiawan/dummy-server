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

func MockServerRouth(e *echo.Echo, h mockserver.Handler, middleware security.MiddlewareService) {
	r := e.Group("mockers/v1/mock")
	r.POST("", h.SetMockDataHandler, middleware.MiddlewareSecurity(nil))
	r.POST("/work-space", h.AddWorkSapceHandler, middleware.MiddlewareSecurity(nil))
	r.POST("/collection", h.AddCollectionHandler, middleware.MiddlewareSecurity(nil))
	r.POST("/children", h.AddChildrenHandler, middleware.MiddlewareSecurity(nil))
	r.GET("/collection/by-workspace/:workspace_id", h.GetCollectionByWorkspaceIdHandler, middleware.MiddlewareSecurity(nil))
	r.GET("/children/by-collection/:collection_id", h.GetChildrenByCollectionIdHandler, middleware.MiddlewareSecurity(nil))
	r.GET("/children/by-children/:children_id", h.GetChildrenByChildrenIdHandler, middleware.MiddlewareSecurity(nil))
	x := e.Group("")
	x.Any("/:workspace_id/*", h.MatchMockHandler)
}
