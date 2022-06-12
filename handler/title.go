package handler

import (
	"fmt"
	"stockin/domain"

	"github.com/labstack/echo/v4"
)

func TitleQuery(c echo.Context) error {
	u := new(domain.URL)
	err := c.Bind(u)
	if err != nil {
		return fmt.Errorf("bind error: %w", err)
	}

	title, err := domain.TitleQuery(u)
	if err != nil {
		return fmt.Errorf("title query error: %w", err)
	}

	err = json(c, domain.Title{
		Title: title.Title,
	})
	if err != nil {
		return fmt.Errorf("json error: %w", err)
	}

	return nil
}
