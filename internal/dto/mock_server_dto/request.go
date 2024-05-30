package mockserverdto

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
	ChildrenId string       `json:"children_id"`
	Reqeust    ReqeustMock  `json:"request"`
	Response   ResponseMock `json:"response"`
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
