package security

import (
	"context"
	"fmt"
	"mocking-server/config"
	"time"

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
	generateToken(context.Context, *JwtCustomClaims, *[]byte, time.Duration) (string, error)
	CreateTokens(context.Context, *JwtCustomClaims) (*Tokens, error)
	ParseJwt(context.Context, *string, []byte) (*jwt.Token, error)
	JwtClaim(context.Context, *jwt.Token) (*JwtCustomClaims, error)
}

type JwtCustomClaims struct {
	Uuid     string                 `json:"uuid"`
	Id       string                 `json:"id"`
	Username string                 `json:"username"`
	Name     string                 `json:"name"`
	Roles    map[string]interface{} `json:"role"`
	jwt.StandardClaims
}
type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (svc *jwtService) generateToken(c context.Context, user *JwtCustomClaims, key *[]byte, expire time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uuid"] = user.Uuid
	claims["id"] = user.Id
	claims["username"] = user.Username
	claims["name"] = user.Name
	claims["role"] = user.Roles
	claims["exp"] = expire
	tokenString, err := token.SignedString(*key)
	if err != nil {
		log.Errorf("generateToken.token.SignedString Err: %v", err)
		return "", err
	}
	return tokenString, nil
}

func (svc jwtService) CreateTokens(c context.Context, user *JwtCustomClaims) (*Tokens, error) {
	byteSecretKey := []byte(svc.RootConfig.Jwt.SecretKey)
	byteRefreshKey := []byte(svc.RootConfig.Jwt.RefreshKey)
	duration, err := time.ParseDuration(svc.RootConfig.Jwt.Expiration)
	if err != nil {
		return nil, err
	}
	reDuration, err := time.ParseDuration(svc.RootConfig.Jwt.ReExpiration)
	if err != nil {
		return nil, err
	}
	accessToken, err := svc.generateToken(c, user, &byteSecretKey, duration)
	if err != nil {
		log.Errorf("CreateTokens.svc.generateToken Err: %v", err)
		return nil, err
	}
	refreshToken, err := svc.generateToken(c, user, &byteRefreshKey, reDuration)
	if err != nil {
		log.Errorf("CreateTokens.svc.generateToken Err: %v", err.Error())
		return nil, err
	}
	return &Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (svc jwtService) ParseJwt(c context.Context, strJwt *string, key []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(*strJwt, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		log.Errorf("ParseJwt Err: %v", err)
		return nil, err
	}
	return token, nil
}

func (svc jwtService) JwtClaim(c context.Context, token *jwt.Token) (*JwtCustomClaims, error) {
	if claim, ok := token.Claims.(jwt.MapClaims); ok {
		claimData := JwtCustomClaims{}
		if uuid, ok := claim["uuid"].(string); ok {
			claimData.Uuid = uuid
		} else {
			return nil, fmt.Errorf("uuid not found or is not a string")
		}
		if id, ok := claim["id"].(string); ok {
			claimData.Id = id
		} else {
			return nil, fmt.Errorf("id not found or is not a string")
		}
		if username, ok := claim["username"].(string); ok {
			claimData.Username = username
		} else {
			return nil, fmt.Errorf("username not found or is not a string")
		}
		if name, ok := claim["name"].(string); ok {
			claimData.Name = name
		} else {
			return nil, fmt.Errorf("name not found or is not a string")
		}
		if roles, ok := claim["role"].(map[string]interface{}); ok {
			claimData.Roles = roles
		} else {
			return nil, fmt.Errorf("roles not found or is not a valid json.RawMessage")
		}
		claimData.StandardClaims = jwt.StandardClaims{}
		return &claimData, nil
	}
	return nil, fmt.Errorf("failed to parse claims from token")

}
