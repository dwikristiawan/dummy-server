package mockserver

import (
	"context"
	"fmt"
	mockserverdto "mocking-server/internal/dto/mock_server_dto/request"
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
	MatchMockController(context.Context, *string, *string, *http.Header, *[]byte, *string) (int, *map[string]string, *[]byte, error)
	AddMemberController(context.Context, *mockserverdto.AddMemberRequest) *utils.BaseResponse
	GetWorkspaceController(context.Context) *utils.BaseResponse
	AddCollectionController(context.Context, *mockserverdto.AddCollectionRequest) *utils.BaseResponse
	AddChildrenController(context.Context, *mockserverdto.AddChildrenRequest) *utils.BaseResponse
	GetCollectionByWorkspaceIdController(context.Context, *string) *utils.BaseResponse
	GetChildrenByCollectionIdController(context.Context, *string) *utils.BaseResponse
	GetChildrenByChildrenIdController(context.Context, *string) *utils.BaseResponse
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
	reqHeader := req.Reqeust.RequestHeader
	resHeader := req.Response.ResponseHeader
	err := ctr.service.AddMockDataService(c, &model.MockData{
		Id:             "",
		ChildrenId:     req.ChildrenId,
		CollectionId:   req.CollectionId,
		RequestMethod:  req.Reqeust.RequestMethod,
		Path:           req.Reqeust.Path,
		RequestHeader:  &reqHeader,
		ResponseHeader: &resHeader,
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
	return utils.SuccessRequest(nil)
}

func (ctr controller) MatchMockController(c context.Context, method *string, path *string, header *http.Header, body *[]byte, workspaceId *string) (int, *map[string]string, *[]byte, error) {
	if *method == "" || *path == "" || header == nil {
		return http.StatusInternalServerError, nil, nil, fmt.Errorf("internal server error")
	}
	return ctr.service.MatchMockService(c, method, path, header, body, workspaceId)
}

func (ctr controller) AddMemberController(c context.Context, req *mockserverdto.AddMemberRequest) *utils.BaseResponse {
	err := ctr.service.AddMemberService(c, &model.Member{
		Id:          "",
		WorkspaceId: req.WorkspaceId,
		UserId:      req.UserId,
		Access:      req.Access,
		IsActive:    false,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	})
	if err != nil {
		return utils.BadRequest(err)
	}
	return utils.SuccessRequest(nil)
}

func (ctr controller) GetWorkspaceController(c context.Context) *utils.BaseResponse {
	data, err := ctr.service.GetWorkSpaceService(c)
	if err != nil {
		return utils.ErrorServerRequest(err)
	}
	return utils.SuccessRequest(data)
}

func (ctr controller) AddCollectionController(c context.Context, req *mockserverdto.AddCollectionRequest) *utils.BaseResponse {
	err := ctr.service.AddCollectionService(c, &model.Collection{
		Id:          "",
		WorkspaceId: req.WorkspaceId,
		ReferenceId: "",
		CreatedAt:   nil,
		UpdatedAt:   nil,
	}, &req.Name)
	if err != nil {
		return utils.ErrorServerRequest(err)
	}
	return utils.SuccessRequest(nil)
}
func (ctr controller) AddChildrenController(c context.Context, req *mockserverdto.AddChildrenRequest) *utils.BaseResponse {
	err := ctr.service.AddChildrenService(c, &model.Children{
		Id:           "",
		CollectionId: req.CollectionId,
		Name:         req.Name,
		Perent:       req.Perent,
		ReferenceId:  "",
		CreatedAt:    nil,
		UpdatedAt:    nil,
	})
	if err != nil {
		return utils.ErrorServerRequest(err)
	}
	return utils.SuccessRequest(nil)
}

func (ctr controller) GetCollectionByWorkspaceIdController(c context.Context, workspaceId *string) *utils.BaseResponse {
	if *workspaceId == "" {
		return utils.BadRequest(fmt.Errorf("workspace_id is required"))
	}
	res, err := ctr.service.GetCollectionByWorkspaceIdService(c, workspaceId)
	if err != nil {
		return utils.ErrorServerRequest(err)
	}
	return utils.SuccessRequest(res)
}
func (ctr controller) GetChildrenByCollectionIdController(c context.Context, collectionId *string) *utils.BaseResponse {
	if *collectionId == "" {
		return utils.BadRequest(fmt.Errorf("workspace_id is required"))
	}
	res, err := ctr.service.GetChildrenByCollectionIdService(c, collectionId)
	if err != nil {
		return utils.ErrorServerRequest(err)
	}
	return utils.SuccessRequest(res)
}
func (ctr controller) GetChildrenByChildrenIdController(c context.Context, childrenId *string) *utils.BaseResponse {
	if *childrenId == "" {
		return utils.BadRequest(fmt.Errorf("workspace_id is required"))
	}
	res, err := ctr.service.GetChildrenByChildrenIdService(c, childrenId)
	if err != nil {
		return utils.ErrorServerRequest(err)
	}
	return utils.SuccessRequest(res)
}
func (ctr controller) GetMockDataListByChildrenIdController(c context.Context, childrenId *string) *utils.BaseResponse {
	if *childrenId == "" {
		return utils.BadRequest(fmt.Errorf("children_id is required"))
	}
	res, err := ctr.service.GetMockDataListByChildrenIdService(c, childrenId)
	if err != nil {
		return utils.ErrorServerRequest(err)
	}
	return utils.SuccessRequest(res)
}
