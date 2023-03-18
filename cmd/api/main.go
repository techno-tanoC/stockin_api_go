package main

import (
	"database/sql"
	"log"
	"os"
	"stockin-api/handlers"

	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	database := os.Getenv("DATABASE")
	db, err := sql.Open("postgres", database)
	if err != nil {
		log.Fatal(err)
	}

	e := handlers.BuildApp(db)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":3000"))
}
