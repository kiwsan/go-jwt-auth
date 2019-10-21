package main

import (
	"github.com/labstack/echo/v4"
	"github.com/kiwsan/go-jwt-auth/controllers"
)

func  main()  {
	e := echo.New()
	
	//home controller
	e.GET("/", controllers.HomeGetHandler)

	//identity controller
	e.POST("/login", controllers.LoginPostHandler)
	e.GET("/me", controllers.MeGetHandler, controllers.IsLoggedIn)

	//token controller
	e.POST("/refresh-tokens", controllers.RefreshAccessTokenPostHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
