package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	ResponseCode int
	Massage      string
	Data         interface{}
}

func BaseReturn(e echo.Context, data *BaseResponse) error {
	return e.JSON(data.ResponseCode, data)
}
func BadRequest(err error) *BaseResponse {
	return &BaseResponse{
		ResponseCode: http.StatusBadRequest,
		Massage:      err.Error(),
		Data:         nil,
	}
}
func SuccessRequest(data interface{}) *BaseResponse {
	return &BaseResponse{
		ResponseCode: http.StatusOK,
		Massage:      "success",
		Data:         data,
	}
}
func ErrorServerRequest(err error) *BaseResponse {
	return &BaseResponse{
		ResponseCode: http.StatusInternalServerError,
		Massage:      err.Error(),
		Data:         nil,
	}
}
