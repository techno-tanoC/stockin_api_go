package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)
	err := e.Start(":3000")
	if err != nil {
		e.Logger.Fatal(err)
	}
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
