package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kiwsan/go-jwt-auth/constants"
	"time"
)

// https://travix.io/encapsulating-dependencies-in-go-b0fd74021f5a
type RefreshToken struct {
	Token      string
	UserName   string
	ExpireDate int64
}

// encapsulation model
func NewRefreshToken(username string) (*RefreshToken, error) {

	exp := time.Now().Add(time.Hour * constants.RefreshTokenExpire).Unix()

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["exp"] = exp

	rt, err := refreshToken.SignedString([]byte(constants.SecretToken)) //secret
	if err != nil {
		return nil, err
	}

	return &RefreshToken{
		Token:      rt,
		UserName:   username,
		ExpireDate: exp,
	}, nil
}
