package main

import (
	"stockin/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.KeyAuth(auth("debug")))

	title := e.Group("/title")
	title.POST("/query", handler.TitleQuery)

	thumbnail := e.Group("/thumbnail")
	thumbnail.POST("/query", handler.ThumbnailQuery)

	err := e.Start(":3000")
	if err != nil {
		e.Logger.Fatal(err)
	}
}

func auth(token string) middleware.KeyAuthValidator {
	return func(key string, c echo.Context) (bool, error) {
		return key == token, nil
	}
}
