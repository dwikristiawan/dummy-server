package model

import (
	"encoding/json"
	"time"
)

type Collection struct {
	Id          string     `json:"id"`
	WorkspaceId string     `json:"workspace_id"`
	ReferenceId string     `json:"reference_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type Children struct {
	Id           string     `json:"id"`
	CollectionId string     `json:"collection"`
	Name         string     `json:"name"`
	Perent       string     `json:"perent"`
	ReferenceId  string     `json:"reference_id"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type MockData struct {
	Id             string           `json:"id"`
	ChildrenId     string           `json:"children_id"`
	CollectionId   string           `json:"collection_id"`
	RequestMethod  RequestMethod    `json:"request_method"`
	Path           string           `json:"path"`
	RequestHeader  *json.RawMessage `json:"request_header"`
	ResponseHeader *json.RawMessage `json:"response_header"`
	RequestBody    string           `json:"request_body"`
	ResponseBody   string           `json:"response_body"`
	ResponseCode   int              `json:"response_code"`
	ReferenceId    string           `jsin:"reference_id"`
	CreatedAt      *time.Time       `json:"created_at"`
	UpdatedAt      *time.Time       `json:"updated_at"`
}

type RequestMethod string
