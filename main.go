package main

import (
	"fmt"
	"github.com/kiwsan/go-jwt-auth/bootstrapper"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	err := bootstrapper.OnStart(e)
	if err != nil {
		fmt.Println(err)
	}

	e.Logger.Fatal(e.Start(":8000"))

}
