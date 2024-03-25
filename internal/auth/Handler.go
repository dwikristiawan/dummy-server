package auth

import (
	"context"
	"mocking-server/internal/dto/auth_dto"
	"mocking-server/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type authHandler struct {
	controller AuthController
}

func NewAuthHandler(controller AuthController) *authHandler {
	return &authHandler{controller: controller}
}

type AuthHandler interface {
	Register(echo.Context) error
	Login(echo.Context) error
}

func (h authHandler) Register(e echo.Context) error {
	req := new(auth_dto.RegisterRequest)

	if err := e.Bind(&req); err != nil {
		log.Error("Err Register.e.Bind Err > ", err)
		return utils.BaseReturn(e, utils.BadRequest(err))
	}
	response := h.controller.Register(context.Background(), req)
	return utils.BaseReturn(e, response)

}

func (h authHandler) Login(e echo.Context) error {
	req := new(auth_dto.LoginRequest)

	if err := e.Bind(&req); err != nil {
		log.Error("Err Login.e.Bind Err > ", err)
		return utils.BaseReturn(e, utils.BadRequest(err))
	}
	response := h.controller.Login(context.Background(), req)
	return utils.BaseReturn(e, response)

}
