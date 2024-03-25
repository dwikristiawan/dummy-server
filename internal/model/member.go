package model

import "time"

type Member struct {
	Id           string     `json:"id"`
	UserId       string     `json:"user_id"`
	CollectionId string     `json:"collection_id"`
	Access       string     `json:"access"`
	IsActive     bool       `json:"is_active"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"update_at"`
}
