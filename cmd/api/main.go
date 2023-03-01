package main

import (
	"database/sql"
	"log"
	"os"
	"stockin-api/handlers"

	_ "github.com/lib/pq"
)

func main() {
	database := os.Getenv("DATABASE")
	db, err := sql.Open("postgres", database)
	if err != nil {
		log.Fatal(err)
	}

	e := handlers.BuildApp(db)
	e.Logger.Fatal(e.Start(":3000"))
}
