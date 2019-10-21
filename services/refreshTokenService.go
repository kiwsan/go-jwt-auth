package services

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/kiwsan/go-jwt-auth/utils"
)

//https://godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac
func RefreshAccessToken(username string, refreshToken string) (map[string]string, error) {
	// Parse takes the token string and a function for looking up the key.
	// The latter is especially useful if you use multiple keys for your application.
	// The standard is to use 'kid' in the head of the token to identify
	// which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("3a52cf25791d406da5e35c4bb446f476"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get the user record from database or
		// run through your business logic to verify if the user can log in
		fmt.Println(claims["username"])

		t, err := utils.CreateToken(username)
		if err != nil {
			return nil, err
		}

		return t, err
	}

	return nil, err
}
