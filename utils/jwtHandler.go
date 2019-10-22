package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kiwsan/go-jwt-auth/constants"
	"github.com/kiwsan/go-jwt-auth/entities"
	"strconv"
	"time"
)

func CreateToken(username string) (map[string]string, error) {

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	exp := time.Now().Add(time.Minute * constants.TokenExpire).Unix()
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = exp

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	t, err := token.SignedString([]byte(constants.SecretToken)) //secret
	if err != nil {
		return nil, err
	}

	// Add refresh token to database
	rt, err := entities.NewRefreshToken(username)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  t,
		"refresh_token": rt.Token,
		"expire_date":   strconv.FormatUint(uint64(exp), 10),
	}, nil
}
