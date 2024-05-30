package mockserver

import (
	"context"
	"fmt"
	mockserverdto "mocking-server/internal/dto/mock_server_dto"
	"mocking-server/internal/model"
	mockserversvc "mocking-server/internal/service/mockserver_svc"
	"mocking-server/utils"
	"net/http"
	"time"
)

type controller struct {
	service mockserversvc.Service
}

type Controller interface {
	AddWorkSapceController(context.Context, *mockserverdto.AddWorkSapcerequest) *utils.BaseResponse
	SetMockDataController(context.Context, *mockserverdto.SetMockDataRequest) *utils.BaseResponse
	MatchMockController(context.Context, *string, *string, *http.Header, *[]byte) (int, *http.ResponseWriter, *[]byte, error)
}

func NewController(service mockserversvc.Service) Controller {
	return &controller{
		service: service,
	}
}

func (ctr controller) AddWorkSapceController(c context.Context, req *mockserverdto.AddWorkSapcerequest) *utils.BaseResponse {
	if req.Name == "" {
		return utils.BadRequest(fmt.Errorf("name is require"))
	}
	newWorkSpace := &model.WorkSpace{
		Id:          "",
		Name:        req.Name,
		ReferenceId: "",
		CreatedAt:   &time.Time{},
		UpdatedAt:   &time.Time{},
	}

	err := ctr.service.AddWorkSapceService(c, newWorkSpace)

	if err != nil {
		return utils.BadRequest(err)
	}
	return utils.SuccessRequest(nil)
}

func (ctr controller) SetMockDataController(c context.Context, req *mockserverdto.SetMockDataRequest) *utils.BaseResponse {
	var errStr []string
	if req.Reqeust.Path == "" {
		errStr = append(errStr, "path is require")
	}
	if req.ChildrenId == "" {
		errStr = append(errStr, "children_id is require")
	}
	if req.Reqeust.RequestMethod == "" {
		errStr = append(errStr, "request_method is require")

	}
	if req.Response.ResponseBody == "" {
		errStr = append(errStr, "response_body is require")
	}
	if req.Response.ResponseCode == 0 {
		errStr = append(errStr, "response_code is require")
	}
	if len(errStr) > 0 {
		return utils.BadRequest(fmt.Errorf("%v", errStr))
	}
	err := ctr.service.AddMockDataService(c, &model.MockData{
		Id:             "",
		ChildrenId:     req.ChildrenId,
		RequestMethod:  req.Reqeust.RequestMethod,
		Path:           req.Reqeust.Path,
		RequestHeader:  req.Reqeust.RequestHeader,
		ResponseHeader: req.Response.ResponseHeader,
		RequestBody:    req.Reqeust.RequestBody,
		ResponseBody:   req.Response.ResponseBody,
		ResponseCode:   req.Response.ResponseCode,
		ReferenceId:    "",
		CreatedAt:      &time.Time{},
		UpdatedAt:      &time.Time{},
	})
	if err != nil {
		return utils.ErrorServerRequest(err)
	}
	return nil
}

func (ctr controller) MatchMockController(c context.Context, method *string, path *string, header *http.Header, body *[]byte) (int, *http.ResponseWriter, *[]byte, error) {
	if *method == "" || *path == "" || header == nil {
		return http.StatusInternalServerError, nil, nil, fmt.Errorf("internal server error")
	}
	return ctr.service.MatchMockService(c, method, path, header, body)
}
