package mockserver

// import (
// 	"context"
// 	mockserverdto "mocking-server/internal/dto/mock_server_dto"
// 	"mocking-server/utils"

// 	"github.com/labstack/echo/v4"
// )

// type mockServerHandler struct {
// 	Controller MockServerController
// }

// func NewMockServerHandler(Controller MockServerController) *mockServerHandler {
// 	return &mockServerHandler{Controller: Controller}
// }

// type MockServerHandler interface {
// 	CreateCollectionHandler(echo.Context) error
// 	InquiryListCollectionHandler(echo.Context) error
// }

// func (hdr mockServerHandler) CreateCollectionHandler(e echo.Context) error {
// 	req := new(mockserverdto.CreateColectionRequest)
// 	err := e.Bind(&req)
// 	if err != nil {
// 		return utils.BaseReturn(e, utils.BadRequest(err))
// 	}
// 	response := hdr.Controller.CreateColectionController(context.Background())
// 	return utils.BaseReturn(e, response)
// }

// func (hdr mockServerHandler)InquiryListCollectionHandler(e echo.Context)error{
// 	e.Request().
// }
