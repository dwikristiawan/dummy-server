package model

import "time"

type Collection struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	ReferenceId string     `json:"reference_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
