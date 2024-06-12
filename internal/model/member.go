package model

import "time"

type Member struct {
	Id          string          `json:"id"`
	WorkspaceId string          `json:"workspace_id"`
	UserId      string          `json:"user_id"`
	Access      AccessWorkSpace `json:"access"`
	IsActive    bool            `json:"is_active"`
	CreatedAt   *time.Time      `json:"created_at"`
	UpdatedAt   *time.Time      `json:"update_at"`
}

type Types struct {
	Name string `json:"name"`
}

type TypeName string
type AccessWorkSpace string

const (
	WORK_SPACE TypeName = "WORK_SPACE"

	CREATOR AccessWorkSpace = "CREATOR"
	ADMIN   AccessWorkSpace = "ADMIN"
	MODIF   AccessWorkSpace = "MODIF"
	READ    AccessWorkSpace = "READ"
)
