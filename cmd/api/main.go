package main

import (
	"database/sql"
	"log"
	"os"
	"stockin-api/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	database := os.Getenv("DATABASE")
	db, err := sql.Open("postgres", database)
	if err != nil {
		log.Fatal(err)
	}

	token := os.Getenv("TOKEN")
	auth := &Auth{token}

	e := handlers.BuildApp(db)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.KeyAuth(auth.auth))

	e.Logger.Fatal(e.Start(":3000"))
}

type Auth struct {
	token string
}

func (a *Auth) auth(key string, c echo.Context) (bool, error) {
	return key == a.token, nil
}
