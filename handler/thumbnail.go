package handler

import (
	"fmt"
	"stockin/domain"

	"github.com/labstack/echo/v4"
)

type Thumbnail struct {
	URL string `json:"url"`
}

func ThumbnailQuery(c echo.Context) error {
	u := new(URL)
	err := c.Bind(u)
	if err != nil {
		return fmt.Errorf("bind error: %w", err)
	}

	url, err := domain.ThumbnailQuery(u.URL)
	if err != nil {
		return fmt.Errorf("thumbnail query error: %w", err)
	}

	err = json(c, Thumbnail{
		URL: url,
	})
	if err != nil {
		return fmt.Errorf("json error: %w", err)
	}

	return nil
}
