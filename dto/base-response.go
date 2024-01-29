package dto

type BaseResponse struct {
	ResponseCode int
	Massage      string
	Data         interface{}
}
