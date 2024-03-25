package users_svc

import (
	"context"
	"mocking-server/internal/model"
	"mocking-server/internal/repository/postgres/users"
	"mocking-server/internal/security"

	"github.com/labstack/gommon/log"
)

type service struct {
	repository users.Repository
}

func NewService(repository users.Repository) *service {
	return &service{repository: repository}
}

type Service interface {
	UserInquiryService(context.Context, *model.Users) (*[]model.Users, error)
	InsertService(context.Context, *model.Users) error
	UpdateUsersService(context.Context, *model.Users) error
	RemoveUsersService(context.Context, *model.Users) error
}

func (svc service) UserInquiryService(c context.Context, req *model.Users) (*[]model.Users, error) {
	response, err := svc.repository.SelectUser(c, req)
	return response, err
}

func (svc service) InsertService(c context.Context, req *model.Users) error {
	if req.Password != "" {
		hashedPassword, err := security.StrHashing(req.Password)
		log.Infof(hashedPassword)
		if err != nil {
			log.Errorf("Err InsertService.security.StrHashing Err > %v", err)
			return err
		}
		req.Password = hashedPassword
	}
	err := svc.repository.InsertUser(c, req)
	return err
}

func (svc service) UpdateUsersService(c context.Context, req *model.Users) error {
	return svc.repository.UpdateUser(c, req)
}
func (svc service) RemoveUsersService(c context.Context, req *model.Users) error {
	return svc.repository.DeleteUser(c, req)
}
