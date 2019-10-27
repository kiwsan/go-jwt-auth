package routes

import (
	"github.com/kiwsan/go-jwt-auth/controllers"
	"github.com/labstack/echo/v4"
)

func NewRoutes(e *echo.Echo) {

	//home controller
	e.GET("/", controllers.HomeGetHandler)

	//identity controller
	e.POST("/sign-in", controllers.LoginPostHandler)
	e.POST("/sign-up", controllers.RegisterPostHandler)
	e.GET("/me", controllers.MeGetHandler, controllers.IsLoggedIn)

	//token controller
	e.POST("/access-tokens/:refreshToken/refresh", controllers.RefreshAccessTokenPostHandler)
	e.POST("/refresh-tokens/:refreshToken/revoke", controllers.RevokeRefreshTokenPostHandler, controllers.IsLoggedIn)

}
