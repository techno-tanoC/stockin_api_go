package handlers

import (
	"database/sql"
	"stockin-api/usecases"

	"github.com/labstack/echo/v4"
)

func BuildApp(db *sql.DB) *echo.Echo {
	e := echo.New()

	items := e.Group("/items")
	{
		usecase := usecases.NewItemUsecase(db)
		handler := NewItemHandler(usecase)

		items.GET("/:id", handler.Find)
		items.POST("/", handler.Create)
	}

	return e
}
