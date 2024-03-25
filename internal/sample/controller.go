package sample

import (
	"context"
	"mocking-server/utils"
)

type controller struct {
	service Service
}

func NewController(service Service) *controller {
	return &controller{service: service}
}

type Controller interface {
	SampleController(c context.Context) utils.BaseResponse
}

func (ctr controller) SampleController(c context.Context) utils.BaseResponse {
	serviceData, err := ctr.service.SampleService(c)
	if err != nil {
		return utils.BaseResponse{
			ResponseCode: 500,
			Massage:      "error on server",
			Data:         err}
	}
	return utils.BaseResponse{ResponseCode: 200, Massage: "success", Data: serviceData}
}
