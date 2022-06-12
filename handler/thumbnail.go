package handler

import (
	"fmt"
	"stockin/domain"

	"github.com/labstack/echo/v4"
)

func ThumbnailQuery(c echo.Context) error {
	u := new(domain.URL)
	err := c.Bind(u)
	if err != nil {
		return fmt.Errorf("bind error: %w", err)
	}

	url, err := domain.ThumbnailQuery(u)
	if err != nil {
		return fmt.Errorf("thumbnail query error: %w", err)
	}

	err = json(c, domain.Thumbnail{
		URL: url.URL,
	})
	if err != nil {
		return fmt.Errorf("json error: %w", err)
	}

	return nil
}
