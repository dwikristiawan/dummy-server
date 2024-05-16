package model

import (
	"encoding/json"
	"time"
)

type Users struct {
	Id        string          `json:"id"`
	Username  string          `json:"username"`
	Name      string          `json:"name"`
	Password  string          `json:"password"`
	Roles     json.RawMessage `json:"roles"`
	Status    string          `json:"status"`
	CreatedAt *time.Time      `json:"created_at"`
	UpdatedAt *time.Time      `json:"updated_at"`
}
type Roles struct {
	Roles map[string]interface{}
}
