package controllers

import (
	"github.com/kiwsan/go-jwt-auth/constants"
	"github.com/labstack/echo/v4/middleware"
)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(constants.SecretToken),
})