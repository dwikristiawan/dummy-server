package sample

import (
	"context"
	"dummy-server/dto"
)

type controller struct {
	service Service
}

func NewController(service Service) *controller {
	return &controller{service: service}
}

type Controller interface {
	SampleController(c context.Context) dto.BaseResponse
}

func (ctr controller) SampleController(c context.Context) dto.BaseResponse {
	serviceData, err := ctr.service.SampleService(c)
	if err != nil {
		return dto.BaseResponse{
			ResponseCode: 500,
			Massage:      "error on server",
			Data:         err}
	}
	return dto.BaseResponse{ResponseCode: 200, Massage: "success", Data: serviceData}
}
