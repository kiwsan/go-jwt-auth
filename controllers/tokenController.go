package controllers

import (
	"github.com/kiwsan/go-jwt-auth/models"
	"github.com/kiwsan/go-jwt-auth/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RefreshAccessTokenPostHandler(c echo.Context) error {

	req := new(models.RefreshTokenRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	jwtToken, err := services.RefreshAccessToken("admin", req.RefreshToken)
	if err != nil {
		return err
	}

	// return jwtToken
	return c.JSON(http.StatusOK, jwtToken)
}

func RevokeAccessTokenHandler() {
}

func RevokeRefreshTokenHandler() {
}
