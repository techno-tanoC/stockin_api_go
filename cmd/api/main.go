package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"stockin-api/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Database string `env:"DATABASE,required"`
	Port     int    `env:"PORT,default=3000"`
	Token    string `env:"TOKEN,required"`
}

func main() {
	ctx := context.Background()
	conf := new(Config)
	err := envconfig.Process(ctx, conf)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", conf.Database)
	if err != nil {
		log.Fatal(err)
	}

	auth := &Auth{token: conf.Token}

	e := handlers.BuildApp(db)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.KeyAuth(auth.auth))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}

type Auth struct {
	token string
}

func (a *Auth) auth(key string, c echo.Context) (bool, error) {
	return key == a.token, nil
}
