package utils

import "github.com/google/uuid"

func IdUuid() string {
	id, _ := uuid.NewRandom()
	return id.String()
}
