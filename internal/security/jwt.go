package security

import (
	"context"
	"encoding/json"
	"fmt"
	"mocking-server/config"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/gommon/log"
)

type jwtService struct {
	RootConfig *config.Root
}

func NewJwtService(RootConfig *config.Root) JwtService {
	return &jwtService{RootConfig: RootConfig}
}

type JwtService interface {
	generateToken(context.Context, *JwtCustomClaims, *[]byte) (string, error)
	CreateTokens(context.Context, *JwtCustomClaims) (*Tokens, error)
	ParseJwt(context.Context, *string) (*jwt.Token, error)
	JwtClaim(context.Context, *jwt.Token) (*JwtCustomClaims, error)
}

type JwtCustomClaims struct {
	Uuid     string          `json:"uuid"`
	Id       string          `json:"id"`
	Username string          `json:"username"`
	Name     string          `json:"name"`
	Roles    json.RawMessage `json:"role"`
	jwt.StandardClaims
}
type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (svc *jwtService) generateToken(c context.Context, user *JwtCustomClaims, key *[]byte) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uuid"] = user.Uuid
	claims["id"] = user.Id
	claims["username"] = user.Username
	claims["name"] = user.Name
	claims["role"] = user.Roles
	claims["exp"] = svc.RootConfig.Jwt.Expiration
	tokenString, err := token.SignedString(*key)
	if err != nil {
		log.Errorf("generateToken.token.SignedString Err: %v", err)
		return "", err
	}
	return tokenString, nil
}

func (svc jwtService) CreateTokens(c context.Context, user *JwtCustomClaims) (*Tokens, error) {
	byteSecretKey := []byte(svc.RootConfig.Jwt.RefreshKey)
	byteRefreshKey := []byte(svc.RootConfig.Jwt.RefreshKey)
	accessToken, err := svc.generateToken(c, user, &byteSecretKey)
	if err != nil {
		log.Errorf("CreateTokens.svc.generateToken Err: %v", err)
		return nil, err
	}
	refreshToken, err := svc.generateToken(c, user, &byteRefreshKey)
	if err != nil {
		log.Errorf("CreateTokens.svc.generateToken Err: %v", err.Error())
		return nil, err
	}
	return &Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (svc jwtService) ParseJwt(c context.Context, strJwt *string) (*jwt.Token, error) {
	token, err := jwt.Parse(*strJwt, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(svc.RootConfig.Jwt.SecretKey), nil
	})
	if err != nil {
		log.Errorf("ParseJwt Err: %v", err)
		return nil, err
	}
	return token, nil
}

func (svc jwtService) JwtClaim(c context.Context, token *jwt.Token) (*JwtCustomClaims, error) {
	if claim, ok := token.Claims.(jwt.MapClaims); ok {
		claimData := JwtCustomClaims{
			Uuid:           claim["uuid"].(string),
			Id:             claim["id"].(string),
			Username:       claim["username"].(string),
			Name:           claim["name"].(string),
			Roles:          claim["roles"].(json.RawMessage),
			StandardClaims: jwt.StandardClaims{},
		}
		return &claimData, nil
	}
	return nil, fmt.Errorf("failed claims token")

}
