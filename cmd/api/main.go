package main

import (
	"context"
	"fmt"
	"log"
	"stockin/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Port  int    `env:"PORT,default=3000"`
	Token string `env:"TOKEN,required"`
}

func main() {
	ctx := context.Background()
	conf := new(Config)
	err := envconfig.Process(ctx, conf)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.KeyAuth(auth(conf.Token)))

	title := e.Group("/title")
	title.POST("/query", handler.TitleQuery)

	thumbnail := e.Group("/thumbnail")
	thumbnail.POST("/query", handler.ThumbnailQuery)

	err = e.Start(fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		log.Fatal(err)
	}
}

func auth(token string) middleware.KeyAuthValidator {
	return func(key string, c echo.Context) (bool, error) {
		return key == token, nil
	}
}
