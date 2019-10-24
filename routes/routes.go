package routes

import (
	"github.com/kiwsan/go-jwt-auth/controllers"
	"github.com/labstack/echo/v4"
)

func NewRoutes(e *echo.Echo) {

	//home controller
	e.GET("/", controllers.HomeGetHandler)

	//identity controller
	e.POST("/login", controllers.LoginPostHandler)
	e.POST("/register", controllers.RegisterPostHandler)
	e.GET("/me", controllers.MeGetHandler, controllers.IsLoggedIn)

	//token controller
	e.POST("/refresh-tokens", controllers.RefreshAccessTokenPostHandler)

}
