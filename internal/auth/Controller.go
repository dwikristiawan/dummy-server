package auth

import (
	"errors"
	"mocking-server/internal/dto/auth_dto"
	"mocking-server/internal/model"
	"mocking-server/internal/service/users_svc"
	"mocking-server/utils"

	"github.com/labstack/gommon/log"
	"golang.org/x/net/context"
)

type controller struct {
	service users_svc.Service
}
type Controller interface {
	RegisterController(context.Context, *auth_dto.RegisterRequest) *utils.BaseResponse
	LoginController(context.Context, *auth_dto.LoginRequest) *utils.BaseResponse
	RefreshTokenController(context.Context, *string) *utils.BaseResponse
}

func NewController(svc users_svc.Service) Controller {
	return &controller{service: svc}
}

func (ctr *controller) RegisterController(c context.Context, req *auth_dto.RegisterRequest) *utils.BaseResponse {
	var errstr = ""
	if req.Username == "" {
		errstr = " username "
	}
	if req.Name == "" {
		errstr = errstr + " name "
	}
	if req.Password == "" {
		errstr = errstr + " password "
	}
	if errstr != "" {
		err := errors.New(errstr + " is required ")
		log.Errorf("Err Register Err > %v", err)
		return utils.BadRequest(err)
	}
	err := ctr.service.AddUserService(c, &model.Users{
		Id:        "",
		Username:  req.Username,
		Name:      req.Name,
		Password:  req.Password,
		Roles:     nil,
		Status:    "",
		CreatedAt: nil,
		UpdatedAt: nil,
	})
	if err != nil {
		return utils.BadRequest(err)
	}
	return utils.SuccessRequest(nil)
}

func (ctr *controller) LoginController(c context.Context, req *auth_dto.LoginRequest) *utils.BaseResponse {
	var errstr = ""
	if req.Username == "" {
		errstr = " username "
	}
	if req.Password == "" {
		errstr = errstr + " password "
	}
	if errstr != "" {
		err := errors.New(errstr + " is required ")
		log.Errorf("Err Login Err > %v", err)
		return utils.BadRequest(err)
	}

	token, err := ctr.service.LoginService(c, &model.Users{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return utils.BadRequest(err)
	}
	return utils.SuccessRequest(token)
}
func (ctr *controller) RefreshTokenController(c context.Context, token *string) *utils.BaseResponse {
	newToken, err := ctr.service.RefreshTokenService(c, token)
	if err != nil {
		return utils.BadRequest(err)
	}
	return utils.SuccessRequest(newToken)
}
