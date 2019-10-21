package bootstrapper

import (
	"github.com/kiwsan/go-jwt-auth/routes"
	"github.com/labstack/echo/v4"
)

func OnStart(e *echo.Echo) {

	//routes
	routes.NewRoutes(e)

}
