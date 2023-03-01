package handlers

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func BuildApp(db *sql.DB) *echo.Echo {
	e := echo.New()

	return e
}
