package model

import "time"

type DummyServer struct {
		Id                 string     `json:"id"`
		Method             string     `json:"method"`
		Path               string     `json:"path"`
		ResponseContenType string     `json:"response_conten_type"`
		ResponseBody       string     `json:"response_body"`
		ResponseCode       int        `json:"response_code"`
		ReferenceId        string     `json:"reference_id"`
		CreateAt           *time.Time `json:"create_at"`
		UpdateAt           *time.Time `json:"update_at"`
}
