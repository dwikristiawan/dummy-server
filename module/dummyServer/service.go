package dummyServer

import (
	"context"
	"errors"
	"time"

	"dummy-server/module/dummyServer/model"

	"github.com/google/uuid"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) service {
	return service{repository: repository}
}

type Service interface {
	SetDummyService(context.Context, *model.DummyServer) (string, error)
	RemoveDummyService(context.Context, string) (string, error)
	GetAllDummyService(context.Context) (*[]model.DummyServer, error)
	GetDummyByIdService(context.Context, string) (*model.DummyServer, error)
}

func (svc service) SetDummyService(c context.Context, newDummy *model.DummyServer) (string, error) {
	var (
		err error
	)

	curentTime := time.Now()
	id, _ := uuid.NewRandom()
	switch newDummy.Id {
	case "":
		newDummy.Id = id.String()
		newDummy.CreateAt = &curentTime
		err = svc.repository.SetDummyRepository(c, newDummy)
		if err != nil {
			return "", err
		}
	default:
		var oldDammy bool
		oldDammy, err = svc.repository.IsExist(c, newDummy.Id)
		if err != nil {
			return "", err
		}
		if !oldDammy {
			err = errors.New("error : data with id not found")
			return "", err
		}
		newDummy.UpdateAt = &curentTime
		err = svc.repository.SetDummyRepository(c, newDummy)
		if err != nil {
			return "", err
		}

	}

	return "set dummy is success", nil
}
func (s service) RemoveDummy(c context.Context, id string) (string, error) {
	err := s.repository.RemoveDummy(c, id)
	if err != nil {
		return "removing fialed", err
	}
	return "success", nil
}
