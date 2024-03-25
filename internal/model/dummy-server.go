package model

import "time"

type DummyServer struct {
	Id                 string     `json:"id"`
	Method             string     `json:"method"`
	Path               string     `json:"path"`
	RequestHeader string     `json:"response_conten_type"`
	ResponseBody       string     `json:"response_body"`
	ResponseCode       int        `json:"response_code"`
	ReferenceId        string     `json:"reference_id"`
	CreatedAt          *time.Time `json:"created_at"`
	CollectionId       string     `json:"collection_id"`
	UpdatedAt          *time.Time `json:"updated_at"`
}
