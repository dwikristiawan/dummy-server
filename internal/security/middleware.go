package security

import (
	"context"
	"mocking-server/config"
	"mocking-server/internal"
	"mocking-server/utils"

	"github.com/labstack/echo/v4"
)

type middlewareService struct {
	Jwt        JwtService
	rootConfig *config.Root
}

func NewMiddlewareService(Jwt JwtService, rootConfig *config.Root) MiddlewareService {
	return &middlewareService{
		Jwt:        Jwt,
		rootConfig: rootConfig,
	}
}

type MiddlewareService interface {
	MiddlewareSecurity(role *map[string]interface{}) echo.MiddlewareFunc
}

func (svc middlewareService) MiddlewareSecurity(role *map[string]interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			jwtRequest := c.Request().Header.Get("Authorization")
			if jwtRequest == "" {
				return c.JSON(echo.ErrUnauthorized.Code, utils.BaseResponse{
					ResponseCode: echo.ErrUnauthorized.Code,
					Massage:      echo.ErrUnauthorized.Error(),
					Data:         nil,
				})
			}
			token, err := svc.Jwt.ParseJwt(c.Request().Context(), &jwtRequest, []byte(svc.rootConfig.Jwt.SecretKey))
			if err != nil || !token.Valid {
				return c.JSON(echo.ErrUnauthorized.Code, utils.BaseResponse{
					ResponseCode: echo.ErrUnauthorized.Code,
					Massage:      err.Error(),
					Data:         nil,
				})
			}
			claimData, err := svc.Jwt.JwtClaim(c.Request().Context(), token)
			if err != nil {
				return c.JSON(echo.ErrUnauthorized.Code, utils.BaseResponse{
					ResponseCode: echo.ErrUnauthorized.Code,
					Massage:      err.Error(),
					Data:         nil,
				})
			}
			ctx := context.WithValue(c.Request().Context(), internal.USER, claimData)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
