package controllers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/kiwsan/go-jwt-auth/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("3a52cf25791d406da5e35c4bb446f476"),
})

func MeGetHandler(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	return c.String(http.StatusOK, "Welcome "+username+"!")
}

func LoginPostHandler(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Check in your db if the user exists or not
	if username != "admin" && password != "password" {
		return echo.ErrUnauthorized
	}

	// Get claims from database

	// Create token
	jwtToken, err := utils.CreateToken(username) // Add claims
	if err != nil {
		return err
	}

	// return jwtToken
	return c.JSON(http.StatusOK, jwtToken)
}
