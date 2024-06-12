package response

import "time"

type CollectionResponse struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	WorkspaceId string     `json:"workspace_id"`
	ReferenceId string     `json:"reference_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
