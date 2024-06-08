package mockserver

import (
	"bytes"
	"fmt"
	"io/ioutil"
	mockserverdto "mocking-server/internal/dto/mock_server_dto"
	"mocking-server/utils"

	"github.com/labstack/echo/v4"
)

// import (
// 	"context"
// 	mockserverdto "mocking-server/internal/dto/mock_server_dto"
// 	"mocking-server/utils"

// 	"github.com/labstack/echo/v4"
// )

type handler struct {
	Controller Controller
}

func NewHandler(Controller Controller) Handler {
	return &handler{Controller: Controller}
}

type Handler interface {
	AddWorkSapceHandler(echo.Context) error
	SetMockDataHandler(echo.Context) error
	MatchMockHandler(echo.Context) error
}

func (h handler) AddWorkSapceHandler(e echo.Context) error {
	req := new(mockserverdto.AddWorkSapcerequest)
	err := e.Bind(&req)
	if err != nil {
		return utils.BaseReturn(e, utils.BadRequest(err))
	}
	res := h.Controller.AddWorkSapceController(e.Request().Context(), req)
	return utils.BaseReturn(e, res)
}

func (h handler) SetMockDataHandler(e echo.Context) error {
	req := new(mockserverdto.SetMockDataRequest)

	err := e.Bind(&req)
	if err != nil {
		return e.JSON(echo.ErrBadRequest.Code, utils.BadRequest(err))
	}
	res := h.Controller.SetMockDataController(e.Request().Context(), req)
	return utils.BaseReturn(e, res)
}
func (h handler) MatchMockHandler(e echo.Context) error {
	header := e.Request().Header
	body, err := ioutil.ReadAll(e.Request().Body)
	if err != nil {
		return e.JSON(echo.ErrBadRequest.Code, utils.BadRequest(err))
	}
	method := e.Request().Method
	path := e.Request().URL.Path

	resCode, resheader, resbody, err := h.Controller.MatchMockController(e.Request().Context(), &method, &path, &header, &body)
	if err != nil {
		return e.JSON(resCode, err.Error())
	}
	for key, value := range *resheader {
		e.Response().Header().Add(key, value)
	}
	e.Response().Status = resCode
	headerRes := *resheader
	byteReaderBody := bytes.NewReader(*resbody)
	fmt.Println(headerRes[echo.HeaderContentType])
	return e.Stream(resCode, headerRes[echo.HeaderContentType], byteReaderBody)
}
