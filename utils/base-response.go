package utils

import "github.com/labstack/echo/v4"

type BaseResponse struct {
	ResponseCode int
	Massage      string
	Data         interface{}
}

func BaseReturn(e echo.Context, data BaseResponse) error {
	return e.JSON(data.ResponseCode, data)
}
