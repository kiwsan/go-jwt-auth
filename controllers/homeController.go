package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func HomeGetHandler(c echo.Context) error  {
	return c.String(http.StatusOK, "Identity Service!")
}