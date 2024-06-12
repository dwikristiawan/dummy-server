package mockserversvc

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"mocking-server/internal"
	"mocking-server/internal/dto/mock_server_dto/response"
	"mocking-server/internal/model"
	"mocking-server/internal/repository/postgres"
	"mocking-server/internal/repository/postgres/mock_server_repository/children"
	"mocking-server/internal/repository/postgres/mock_server_repository/collection"
	"mocking-server/internal/repository/postgres/mock_server_repository/member"
	mockdata "mocking-server/internal/repository/postgres/mock_server_repository/mock_data"
	workspace "mocking-server/internal/repository/postgres/mock_server_repository/work_space"
	"mocking-server/utils"
	"net/http"
	"time"

	"github.com/labstack/gommon/log"
)

type service struct {
	postgresRepo         postgres.Repository
	workspaceRepository  workspace.Reppsitory
	memberRepository     member.Reppsitory
	mockDataRepository   mockdata.Reppsitory
	collectionRepository collection.Repository
	childrenRepository   children.Repository
}
type Service interface {
	AddWorkSapceService(context.Context, *model.WorkSpace) error
	AddMemberService(context.Context, *model.Member) error
	AddMockDataService(context.Context, *model.MockData) error
	MatchMockService(context.Context, *string, *string, *http.Header, *[]byte, *string) (int, *map[string]string, *[]byte, error)
	GetWorkSpaceService(context.Context) (*[]model.WorkSpace, error)
	AddCollectionService(context.Context, *model.Collection, *string) error
	AddChildrenService(context.Context, *model.Children) error
	GetCollectionByWorkspaceIdService(context.Context, *string) (*[]response.CollectionResponse, error)
	GetChildrenByCollectionIdService(context.Context, *string) (*[]model.Children, error)
	GetChildrenByChildrenIdService(context.Context, *string) (*[]model.Children, error)
	GetMockDataListByChildrenIdService(context.Context, *string) (*[]model.MockData, error)
}

func NewService(
	postgresRepo postgres.Repository,
	workspaceRepository workspace.Reppsitory,
	memberRepository member.Reppsitory,
	mockDataRepository mockdata.Reppsitory,
	collectionRepository collection.Repository,
	childrenRepository children.Repository) Service {
	return &service{
		postgresRepo:         postgresRepo,
		workspaceRepository:  workspaceRepository,
		memberRepository:     memberRepository,
		mockDataRepository:   mockDataRepository,
		collectionRepository: collectionRepository,
		childrenRepository:   childrenRepository}
}
func (svc service) AddWorkSapceService(c context.Context, req *model.WorkSpace) error {
	tx, err := svc.postgresRepo.DBBegin()
	if err != nil {
		return err
	}
	//add work space
	userId := fmt.Sprint(c.Value(internal.USER_ID))
	curentTime := time.Now()
	req.Id = utils.IdUuid()
	req.ReferenceId = userId
	req.CreatedAt = &curentTime
	req.UpdatedAt = nil

	err = svc.workspaceRepository.InsertWorkSpace(c, tx, req)
	if err != nil {
		tx.Rollback()
		return err
	}
	newMember := model.Member{
		Id:          utils.IdUuid(),
		WorkspaceId: req.Id,
		UserId:      userId,
		Access:      model.CREATOR,
		IsActive:    true,
		CreatedAt:   &curentTime,
		UpdatedAt:   nil,
	}
	err = svc.memberRepository.InsertMember(c, tx, &newMember)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (svc service) AddMemberService(c context.Context, req *model.Member) error {
	curentTime := time.Now()
	tx, err := svc.postgresRepo.DBBegin()
	if err != nil {
		return err
	}
	req.Id = utils.IdUuid()
	req.CreatedAt = &curentTime
	err = svc.memberRepository.InsertMember(c, tx, req)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (svc service) AddMockDataService(c context.Context, req *model.MockData) error {
	tx, err := svc.mockDataRepository.DBBegin()
	if err != nil {
		return err
	}
	curentTime := time.Now()

	req.CreatedAt = &curentTime
	req.UpdatedAt = nil
	req.Id = utils.IdUuid()
	req.ReferenceId = fmt.Sprint(c.Value(internal.USER_ID))
	err = svc.mockDataRepository.InsertMockData(c, tx, req)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (svc service) MatchMockService(c context.Context, method *string, path *string, header *http.Header, body *[]byte, workspaceId *string) (int, *map[string]string, *[]byte, error) {

	var base64Body string
	if body != nil {
		base64Body = base64.RawStdEncoding.EncodeToString(*body)
	}
	mockDatas, err := svc.mockDataRepository.SelectMockDataByworkspaceId(c, &model.MockData{
		RequestMethod: model.RequestMethod(*method),
		Path:          fmt.Sprintf("/%s", *path),
		RequestBody:   base64Body,
	}, workspaceId)
	if err != nil {
		return http.StatusInternalServerError, nil, nil, err
	}
	if len(*mockDatas) == 0 {
		err = fmt.Errorf("no data is match")
		return http.StatusNotFound, nil, nil, err
	}

	var resMockData model.MockData
	var count = 0
	var ismatch = true
	for _, mockData := range *mockDatas {
		if mockData.RequestHeader == nil {
			resMockData = mockData
		} else {
			var tmpCount int
			var headerMock map[string]string
			err := json.Unmarshal(*mockData.RequestHeader, &headerMock)
			if err != nil {
				return http.StatusInternalServerError, nil, nil, err
			}

			ismatch = true
			for key, value := range headerMock {
				reqValue := header.Get(key)
				if reqValue != value {
					ismatch = false
				}
				tmpCount = tmpCount + 1
			}

			if count < tmpCount {
				count = tmpCount
			}

			if ismatch {
				resMockData = mockData
			}
		}

	}

	if !ismatch && count != 0 {
		err = fmt.Errorf("found response, but dosn't match request header")
		return http.StatusBadRequest, nil, nil, err
	}

	var responseHeader map[string]string

	err = json.Unmarshal(*resMockData.ResponseHeader, &responseHeader)
	if err != nil {
		log.Errorf("setResponseMatcher.json.Unmarshal Err: %v", err)
		return http.StatusInternalServerError, nil, nil, err
	}
	if len(responseHeader) < 1 {
		err = fmt.Errorf("header not found")
		log.Errorf("setResponseMatcher.len(responseHeader) Err: %v", err)
		return http.StatusBadRequest, nil, nil, err
	}
	decodedLen := base64.RawStdEncoding.DecodedLen(len(resMockData.ResponseBody))
	byteBody := make([]byte, decodedLen)
	_, err = base64.StdEncoding.Decode(byteBody, []byte(resMockData.ResponseBody))
	if err != nil {
		log.Errorf("setResponseMatcher.base64.RawStdEncoding.DecodeString Err: %v", err)
		return http.StatusInternalServerError, nil, nil, err
	}
	return resMockData.ResponseCode, &responseHeader, &byteBody, nil

}
func (svc service) GetWorkSpaceService(c context.Context) (*[]model.WorkSpace, error) {
	userId := fmt.Sprint(c.Value(internal.USER_ID))
	data, err := svc.workspaceRepository.SelectWorkSpaceByMemberId(c, &userId)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (svc service) AddCollectionService(c context.Context, req *model.Collection, name *string) error {
	tx, err := svc.postgresRepo.DBBegin()
	if err != nil {
		return err
	}
	req.Id = utils.IdUuid()
	req.ReferenceId = fmt.Sprint(c.Value(internal.USER_ID))
	curentTime := time.Now()
	req.CreatedAt = &curentTime
	err = svc.collectionRepository.InsertCollection(c, tx, req)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = svc.childrenRepository.InsertChildren(c, tx, &model.Children{
		Id:           utils.IdUuid(),
		CollectionId: req.Id,
		Name:         *name,
		Perent:       req.Id,
		ReferenceId:  fmt.Sprint(c.Value(internal.USER_ID)),
		CreatedAt:    &curentTime,
		UpdatedAt:    nil,
	})
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func (svc service) AddChildrenService(c context.Context, children *model.Children) error {
	tx, err := svc.postgresRepo.DBBegin()
	if err != nil {
		return err
	}
	children.Id = utils.IdUuid()
	curentTime := time.Now()
	children.CreatedAt = &curentTime
	children.ReferenceId = fmt.Sprint(c.Value(internal.USER_ID))
	err = svc.childrenRepository.InsertChildren(c, tx, children)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func (svc service) GetCollectionByWorkspaceIdService(c context.Context, workspaceId *string) (*[]response.CollectionResponse, error) {
	collections, err := svc.collectionRepository.SelectByWorkspaceId(c, workspaceId)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

func (svc service) GetChildrenByCollectionIdService(c context.Context, collectionId *string) (*[]model.Children, error) {
	childrens, err := svc.childrenRepository.SelectByCollectionId(c, collectionId)
	if err != nil {
		return nil, err
	}
	return childrens, nil
}
func (svc service) GetChildrenByChildrenIdService(c context.Context, collectionId *string) (*[]model.Children, error) {
	childrens, err := svc.childrenRepository.SelectByChildrenId(c, collectionId)
	if err != nil {
		return nil, err
	}
	return childrens, nil
}
func (svc service) GetMockDataListByChildrenIdService(c context.Context, childrenId *string) (*[]model.MockData, error) {
	res, err := svc.mockDataRepository.SelectMockData(c, &model.MockData{ChildrenId: *childrenId})
	if err != nil {
		return nil, err
	}
	return res, nil
}
