package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kiwsan/go-jwt-auth/constants"
	"time"
)

type JsonWebToken struct {
	Token      string
	UserName   string
	ExpireDate int64
	CreatedAt  time.Time
}

// encapsulation model
func CreateRefreshToken(username string) (*JsonWebToken, error) {

	exp := time.Now().Add(time.Hour * constants.RefreshTokenExpire).Unix()

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["exp"] = exp

	rt, err := refreshToken.SignedString([]byte(constants.SecretToken)) //secret
	if err != nil {
		return nil, err
	}

	return &JsonWebToken{
		Token:      rt,
		UserName:   username,
		ExpireDate: exp,
		CreatedAt:  time.Now().UTC(),
	}, nil
}
