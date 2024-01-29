package sample

import (
	"context"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

type Service interface {
	SampleService(c context.Context) (string, error)
}

func (s service) SampleService(c context.Context) (string, error) {
	return "this is sample", nil
}
