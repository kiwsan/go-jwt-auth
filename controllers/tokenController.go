package controllers

import (
	"net/http"

	"github.com/kiwsan/go-jwt-auth/services"
	"github.com/labstack/echo/v4"
)

type (
	refreshTokenReq struct {
		RefreshToken string `json:"refresh_token"`
	}
)

func RefreshAccessTokenPostHandler(c echo.Context) error {
	req := new(refreshTokenReq)
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
