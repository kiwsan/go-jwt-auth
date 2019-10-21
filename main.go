package main

import (
	"github.com/kiwsan/go-jwt-auth/bootstrapper"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	bootstrapper.OnStart(e)

	e.Logger.Fatal(e.Start(":8000"))

}
