package auth

import (
	"mocking-server/internal/dto/auth_dto"
	"mocking-server/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type handler struct {
	controller Controller
}
type Handler interface {
	RegisterHandler(echo.Context) error
	LoginHandler(echo.Context) error
	RefreshTokenHandler(echo.Context) error
}

func NewHandler(ctr Controller) Handler {
	return &handler{controller: ctr}
}

func (h *handler) RegisterHandler(e echo.Context) error {
	req := new(auth_dto.RegisterRequest)

	if err := e.Bind(&req); err != nil {
		log.Errorf("Register.e.Bind Err:  %v", err)
		return utils.BaseReturn(e, utils.BadRequest(err))
	}
	response := h.controller.RegisterController(e.Request().Context(), req)
	return utils.BaseReturn(e, response)

}

func (h *handler) LoginHandler(e echo.Context) error {
	req := new(auth_dto.LoginRequest)

	if err := e.Bind(&req); err != nil {
		log.Errorf("Err Login.e.Bind Err: %v", err)
		return utils.BaseReturn(e, utils.BadRequest(err))
	}
	response := h.controller.LoginController(e.Request().Context(), req)
	return utils.BaseReturn(e, response)
}
func (h *handler) RefreshTokenHandler(e echo.Context) error {
	token := e.Request().Header.Get("Authorization")
	if token == "" {
		return utils.BaseReturn(e, utils.ErrorForbiden(nil))
	}
	response := h.controller.RefreshTokenController(e.Request().Context(), &token)
	return utils.BaseReturn(e, response)
}
