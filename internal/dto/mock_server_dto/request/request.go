package request

import (
	"encoding/json"
	"mocking-server/internal/model"
)

type CreateColectionRequest struct {
	Name string `json:"name"`
}
type AddWorkSapcerequest struct {
	Name string `json:"name"`
}

type SetMockDataRequest struct {
	ChildrenId   string       `json:"children_id"`
	CollectionId string       `json:"collection_id"`
	Reqeust      ReqeustMock  `json:"request"`
	Response     ResponseMock `json:"response"`
}

type ReqeustMock struct {
	RequestMethod model.RequestMethod `json:"request_method"`
	Path          string              `json:"path"`
	RequestHeader json.RawMessage     `json:"request_header"`
	RequestBody   string              `json:"request_body"`
}
type ResponseMock struct {
	ResponseHeader json.RawMessage `json:"response_header"`
	ResponseBody   string          `json:"response_body"`
	ResponseCode   int             `json:"response_code"`
}
type MatchMockRequest struct {
	Body json.RawMessage
}
type AddMemberRequest struct {
	WorkspaceId string                `json:"work_space_id"`
	UserId      string                `json:"user_id"`
	Access      model.AccessWorkSpace `json:"access"`
}
type AddCollectionRequest struct {
	WorkspaceId string `json:"workspace_id"`
	Name        string `json:"name"`
}
type AddChildrenRequest struct {
	CollectionId string `json:"collection_id"`
	Perent       string `json:"perent"`
	Name         string `json:"name"`
}
