package mockserver

import (
	"bytes"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	mockserverdto "mocking-server/internal/dto/mock_server_dto/request"
	"mocking-server/utils"
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
	AddMemberHandler(echo.Context) error
	GetWorkSapceHandler(echo.Context) error
	MatchMockHandler(echo.Context) error
	AddCollectionHandler(echo.Context) error
	AddChildrenHandler(echo.Context) error
	GetCollectionByWorkspaceIdHandler(echo.Context) error
	GetChildrenByCollectionIdHandler(echo.Context) error
	GetChildrenByChildrenIdHandler(echo.Context) error
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
	workspaceId := e.Param("workspace_id")
	path := e.Param("*")
	header := e.Request().Header
	body, err := ioutil.ReadAll(e.Request().Body)
	if err != nil {
		return e.JSON(echo.ErrBadRequest.Code, utils.BadRequest(err))
	}
	method := e.Request().Method
	resCode, resheader, resbody, err := h.Controller.MatchMockController(e.Request().Context(), &method, &path, &header, &body, &workspaceId)
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

func (h handler) AddMemberHandler(e echo.Context) error {
	req := new(mockserverdto.AddMemberRequest)
	err := e.Bind(&req)
	if err != nil {
		return err
	}
	res := h.Controller.AddMemberController(e.Request().Context(), req)
	return utils.BaseReturn(e, res)
}

func (h handler) GetWorkSapceHandler(e echo.Context) error {
	res := h.Controller.GetWorkspaceController(e.Request().Context())
	return utils.BaseReturn(e, res)
}

func (h handler) AddCollectionHandler(e echo.Context) error {
	req := new(mockserverdto.AddCollectionRequest)
	err := e.Bind(&req)
	if err != nil {
		return utils.BaseReturn(e, utils.ErrorServerRequest(err))
	}
	res := h.Controller.AddCollectionController(e.Request().Context(), req)
	return utils.BaseReturn(e, res)
}
func (h handler) AddChildrenHandler(e echo.Context) error {
	req := new(mockserverdto.AddChildrenRequest)
	err := e.Bind(&req)
	if err != nil {
		return utils.BaseReturn(e, utils.ErrorServerRequest(err))
	}
	res := h.Controller.AddChildrenController(e.Request().Context(), req)
	return utils.BaseReturn(e, res)
}
func (h handler) GetCollectionByWorkspaceIdHandler(e echo.Context) error {
	workspaceId := e.Param("workspace_id")
	res := h.Controller.GetCollectionByWorkspaceIdController(e.Request().Context(), &workspaceId)
	return utils.BaseReturn(e, res)
}
func (h handler) GetChildrenByCollectionIdHandler(e echo.Context) error {
	collectionId := e.Param("collection_id")
	res := h.Controller.GetChildrenByCollectionIdController(e.Request().Context(), &collectionId)
	return utils.BaseReturn(e, res)
}
func (h handler) GetChildrenByChildrenIdHandler(e echo.Context) error {
	childrenId := e.Param("children_id")
	res := h.Controller.GetChildrenByChildrenIdController(e.Request().Context(), &childrenId)
	return utils.BaseReturn(e, res)
}
