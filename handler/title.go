package handler

import (
	"fmt"
	"stockin/domain"

	"github.com/labstack/echo/v4"
)

type Title struct {
	Title string `json:"title"`
}

func TitleQuery(c echo.Context) error {
	u := new(URL)
	err := c.Bind(u)
	if err != nil {
		return fmt.Errorf("bind error: %w", err)
	}

	title, err := domain.TitleQuery(u.URL)
	if err != nil {
		return fmt.Errorf("title query error: %w", err)
	}

	err = json(c, Title{
		Title: title,
	})
	if err != nil {
		return fmt.Errorf("json error: %w", err)
	}

	return nil
}
