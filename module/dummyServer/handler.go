package dummyServer

import (
	"context"
	"dummy-server/module/dummyServer/dto"
	"dummy-server/utils"

	"github.com/labstack/echo/v4"
)

type handler struct {
	controller Controller
}

func NewHandler(controller Controller) *handler {
	return &handler{controller: controller}
}

type Handler interface {
	SetDummyHandler(echo.Context) error
	RemoveDummyHandler(echo.Context)error
	GetAllDummyHandler(echo.Context)error
	GetDummyByIdHandler(echo.Context)error
	MatchingDummyHandler(echo.Context)error
}

func (h handler) SetDummyHandler(e echo.Context) error {
	var req *dto.SetDummyRequest
	err := e.Bind(&req)
	if err != nil {
		response:=utils.BaseResponse{ResponseCode: 400, Massage: "invalid request", Data: err}
		return utils.BaseReturn(e,response)
	}

	response := h.controller.SetDummyController(context.Background(), req)
	return utils.BaseReturn(e,response)
}

func(h handler)RemoveDummyHandler(e echo.Context)error{
	id:=e.QueryParam("id")
	response:=h.controller.RemoveDummyController(context.Background(),id)
	return utils.BaseReturn(e,response)
}
func(h handler)GetAllDummyHandler(e echo.Context)error{
	response:=h.controller.GetAllDummyController(context.Background())
	return utils.BaseReturn(e,response)
}
func(h handler)GetDummyByIdHandler(e echo.Context)error{
	id:=e.QueryParam("id")
	response:=h.controller.GetDummyByIdController(context.Background(),id)
	return utils.BaseReturn(e,response)
}
