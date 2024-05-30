package mockserversvc

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"mocking-server/internal"
	"mocking-server/internal/model"
	"mocking-server/internal/repository/postgres"
	"mocking-server/internal/repository/postgres/member"
	mockdata "mocking-server/internal/repository/postgres/mock_server_repository/mock_data"
	workspace "mocking-server/internal/repository/postgres/mock_server_repository/work_space"
	"mocking-server/utils"
	"net/http"
	"time"

	"github.com/labstack/gommon/log"
)

type service struct {
	repository          postgres.Reppsitory
	workspaceRepository workspace.Reppsitory
	memberRepository    member.Reppsitory
	mockDataRepository  mockdata.Reppsitory
}
type Service interface {
	AddWorkSapceService(context.Context, *model.WorkSpace) error
	AddMemberService(context.Context, *model.Member) error
	AddMockDataService(context.Context, *model.MockData) error
	MatchMockService(context.Context, *string, *string, *http.Header, *[]byte) (int, *http.ResponseWriter, *[]byte, error)
}

func NewService(
	repository postgres.Reppsitory,
	workspaceRepository workspace.Reppsitory,
	memberRepository member.Reppsitory,
	mockDataRepository mockdata.Reppsitory) Service {
	return &service{
		workspaceRepository: workspaceRepository,
		memberRepository:    memberRepository,
		mockDataRepository:  mockDataRepository}
}
func (svc service) AddWorkSapceService(c context.Context, req *model.WorkSpace) error {
	tx, err := svc.repository.DBBegin()
	if err != nil {
		log.Errorf("AddWorkSapceService.svc.repository.DBBegin() Err: %v", err)
		return err
	}
	//add work space
	curentTime := time.Now()
	req.Id = utils.IdUuid()
	req.ReferenceId = fmt.Sprint(c.Value(internal.USER_ID))
	req.CreatedAt = &curentTime

	err = svc.workspaceRepository.InsertWorkSpace(c, tx, req)
	if err != nil {
		log.Errorf("AddWorkSapceService.svc.workspaceRepository.InsertWorkSpace Err: %v", err)
		tx.Rollback()
		return err
	}
	// newType := model.Type{
	// 	Id:        utils.IdUuid(),
	// 	Name:      string(model.WORK_SPACE),
	// 	CreatedAt: &time.Time{},
	// 	UpdatedAt: &time.Time{},
	// }

	newMember := model.Member{
		Id: utils.IdUuid(),
		//TypeId:    newType,
		UserId:    fmt.Sprint(c.Value(internal.USER_ID)),
		Access:    model.CREATOR,
		IsActive:  true,
		CreatedAt: &curentTime,
		UpdatedAt: &time.Time{},
	}
	err = svc.memberRepository.InsertMember(c, tx, &newMember)
	if err != nil {
		log.Errorf("AddWorkSapceService.svc.memberRepository.InsertMember Err: %v", err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (svc service) AddMemberService(c context.Context, req *model.Member) error {
	curentTime := time.Now()
	req.Id = utils.IdUuid()
	req.Access = model.READ
	req.CreatedAt = &curentTime
	return nil
}

func (svc service) AddMockDataService(c context.Context, req *model.MockData) error {

	tx, err := svc.repository.DBBegin()
	if err != nil {
		return err
	}
	curentTime := time.Now()

	req.CreatedAt = &curentTime
	req.Id = utils.IdUuid()
	req.ReferenceId = fmt.Sprint(c.Value(internal.USER_ID))
	err = svc.mockDataRepository.InsertMockData(c, tx, req)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (svc service) MatchMockService(c context.Context, method *string, path *string, header *http.Header, body *[]byte) (int, *http.ResponseWriter, *[]byte, error) {
	var base64Body string
	if body != nil {
		base64Body = base64.RawStdEncoding.EncodeToString(*body)
	}
	mockDatas, err := svc.mockDataRepository.SelectMockData(c, &model.MockData{
		RequestMethod: model.RequestMethod(*method),
		Path:          *path,
		RequestBody:   base64Body,
	})
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
		}
		var tmpCount int
		var headerMock map[string]string
		err := json.Unmarshal(mockData.RequestHeader, &headerMock)
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

	if !ismatch && count != 0 {
		err = fmt.Errorf("found response, but dosn't match request header")
		return http.StatusBadRequest, nil, nil, err
	}

	var responseHeader map[string]string
	var responseWriter http.ResponseWriter

	err = json.Unmarshal(resMockData.ResponseHeader, &responseHeader)
	if err != nil {
		log.Errorf("setResponseMatcher.json.Unmarshal Err: %v", err)
		return http.StatusInternalServerError, nil, nil, err
	}
	if len(responseHeader) < 1 {
		err = fmt.Errorf("header not found")
		log.Errorf("setResponseMatcher.len(responseHeader) < 1   Err: %v", err)
		return http.StatusBadRequest, nil, nil, err
	}
	for key, value := range responseHeader {
		responseWriter.Header().Set(key, value)
	}
	byteBody, err := base64.RawStdEncoding.DecodeString(resMockData.ResponseBody)
	if err != nil {
		return http.StatusInternalServerError, nil, nil, err
	}
	return resMockData.ResponseCode, &responseWriter, &byteBody, nil

}
