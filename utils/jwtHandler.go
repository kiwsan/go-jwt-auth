package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kiwsan/go-jwt-auth/constants"
	"github.com/kiwsan/go-jwt-auth/models"
	"time"
)

func CreateToken(email string) (*models.TokenResponse, error) {

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	exp := time.Now().Add(time.Minute * constants.TokenExpire).Unix()
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = exp

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	t, err := token.SignedString([]byte(constants.SecretToken)) //secret
	if err != nil {
		return nil, err
	}

	// Add refresh token to database
	rt, err := CreateRefreshToken(email)
	if err != nil {
		return nil, err
	}

	return &models.TokenResponse{
		AccessToken:  t,
		RefreshToken: rt.Token,
		ExpireDate:   exp,
	}, nil
}
