package dummyServer

import (
	"context"
	"dummy-server/module/dummyServer/dto"
	"dummy-server/module/dummyServer/model"
	"dummy-server/utils"
)

type controller struct {
	service Service
}

func NewController(service Service) *controller {
	return &controller{service: service}
}

type Controller interface {
	SetDummyController(context.Context, *dto.SetDummyRequest) utils.BaseResponse
	RemoveDummyController(context.Context, string) utils.BaseResponse
	GetAllDummyController(context.Context) utils.BaseResponse
	GetDummyByIdController(context.Context, string) utils.BaseResponse
}

func (ctr controller) SetDummyController(c context.Context, req *dto.SetDummyRequest) utils.BaseResponse {
	if req.Method == "" || req.Path == "" || req.ResponseBody == "" || req.ResponseCode == 0 {
		return utils.BaseResponse{
			ResponseCode: 400,
			Massage:      "method, path, response_body, and response_code is required",
			Data:         nil,
		}
	}
	dummyData := &model.DummyServer{
		Id:                 req.Id,
		Method:             req.Method,
		Path:               req.Path,
		ResponseContenType: req.ResponseContentType,
		ResponseBody:       req.ResponseBody,
		ResponseCode:       req.ResponseCode,
	}
	response, err := ctr.service.SetDummyService(c, dummyData)
	if err != nil {
		return utils.BaseResponse{
			ResponseCode: 400,
			Massage:      "failed",
			Data:         err,
		}
	}

	return utils.BaseResponse{
		ResponseCode: 200,
		Massage:      "succes",
		Data:         response,
	}
}
func (ctr controller) RemoveDummyController(c context.Context, id string) utils.BaseResponse {
	if id == "" {
		return utils.BaseResponse{ResponseCode: 400, Massage: "id is require,request id failed"}
	}
	msg, err := ctr.service.RemoveDummyService(c, id)
	if err != nil {
		return utils.BaseResponse{ResponseCode: 400, Massage: msg, Data: err}
	}
	return utils.BaseResponse{ResponseCode: 200, Massage: msg}
}
func (ctr controller) GetAllDummyController(c context.Context) utils.BaseResponse {
	data, err := ctr.service.GetAllDummyService(c)
	if err != nil {
		return utils.BaseResponse{ResponseCode: 400, Massage: "error: " + err.Error(), Data: err}
	}
	return utils.BaseResponse{ResponseCode: 200, Massage: "success", Data: data}
}
func (ctr controller) GetDummyByIdController(c context.Context, id string) utils.BaseResponse {
	if id == "" {
		return utils.BaseResponse{ResponseCode: 400, Massage: "id is require,request id failed"}
	}
	data, err := ctr.service.GetDummyByIdService(c, id)
	if err != nil {
		return utils.BaseResponse{ResponseCode: 400, Massage: "request failed", Data: err}
	}
	return utils.BaseResponse{ResponseCode: 200, Massage: "success", Data: data}
}
