package security

import (
	"mocking-server/internal/model"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("SECRET_KEY")
var refreshSecretKey = []byte("REFRESH_SECRET_KEY")

func generateToken(user model.Users, key []byte, expiration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["username"] = user.Username
	claims["exp"] = expiration
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

type tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func CreateTokens(user model.Users) (tokens, error) {
	accessToken, err := generateToken(user, secretKey, time.Hour*24)
	if err != nil {
		return tokens{}, err
	}
	refreshToken, err := generateToken(user, refreshSecretKey, time.Hour*24*30)
	if err != nil {
		return tokens{}, err
	}
	return tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
