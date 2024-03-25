package auth

import (
	"errors"
	"mocking-server/internal/dto/auth_dto"
	"mocking-server/internal/model"
	"mocking-server/internal/security"
	"mocking-server/internal/service/users_svc"
	"mocking-server/utils"

	"github.com/labstack/gommon/log"
	"golang.org/x/net/context"
)

type authController struct {
	service users_svc.Service
}

func NewAuthController(service users_svc.Service) *authController {
	return &authController{service: service}
}

type AuthController interface {
	Register(context.Context, *auth_dto.RegisterRequest) *utils.BaseResponse
	Login(context.Context, *auth_dto.LoginRequest) *utils.BaseResponse
}

func (ctr authController) Register(c context.Context, req *auth_dto.RegisterRequest) *utils.BaseResponse {
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
	err := ctr.service.InsertService(c, &model.Users{
		Username: req.Username,
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return utils.BadRequest(err)
	}
	return utils.SuccessRequest(nil)
}

func (ctr authController) Login(c context.Context, req *auth_dto.LoginRequest) *utils.BaseResponse {
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

	users, err := ctr.service.UserInquiryService(c, &model.Users{Username: req.Username})
	if (*users)[0].Username == "" && err == nil {
		err = errors.New(" not found ")
		log.Errorf("Err Login.ctr.service.GetUserService Err > %v", err)
	}
	if err != nil {
		return utils.BadRequest(err)
	}
	user := (*users)[0]

	if err = security.CompareHashingData(user.Password, req.Password); err != nil {
		return utils.BadRequest(err)
	}
	token, err := security.CreateTokens(user)
	if err != nil {
		log.Errorf("Err Login.security.CreateTokens Err > ", err)
		return utils.BadRequest(err)
	}
	return utils.SuccessRequest(token)

}
