package model

import "time"

type Users struct {
	Id        string     `json:"id"`
	Username  string     `json:"username"`
	Name      string     `json:"name"`
	Password  string     `json:"password"`
	Status    string     `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
