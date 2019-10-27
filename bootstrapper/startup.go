package bootstrapper

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/kiwsan/go-jwt-auth/routes"
	"github.com/kiwsan/go-jwt-auth/utils"
	"github.com/labstack/echo/v4"
	"os"
)

func OnStart(e *echo.Echo) error {

	//routes
	routes.NewRoutes(e)

	//environment
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		fmt.Print("default: ")
		environment = "development"
	}

	println(environment)

	err := configor.Load(&utils.Config, "appsettings."+environment+".yml")
	if err != nil {
		return err
	}

	return nil
}
