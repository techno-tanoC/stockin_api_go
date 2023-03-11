package handlers

import (
	"database/sql"
	"stockin-api/usecases"

	"github.com/labstack/echo/v4"
	"github.com/techno-tanoC/sins"
)

func BuildApp(db *sql.DB) *echo.Echo {
	e := echo.New()

	items := e.Group("/items")
	{
		usecase := usecases.NewItemUsecase(db)
		handler := NewItemHandler(usecase)

		items.GET("/:id", handler.Find)
		items.GET("/", handler.FindByRange)
		items.POST("/", handler.Create)
		items.PUT("/:id", handler.Update)
		items.DELETE("/:id", handler.Delete)
	}

	query := e.Group("/query")
	{
		fetcher := sins.NewFetcher()
		handler := NewQueryHandler(fetcher)
		query.POST("/title", handler.Title)
		query.POST("/thumbnail", handler.Thumbnail)
	}

	return e
}
