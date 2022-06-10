package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ThumbnailQuery(c echo.Context) error {
	u := new(URL)
	err := c.Bind(u)
	if err != nil {
		return fmt.Errorf("bind error: %w", err)
	}

	doc, err := fetchDocument(u.URL)
	if err != nil {
		return fmt.Errorf("fetch document error: %w", err)
	}

	// TODO find favicon if not found
	url := doc.Find(`html > head > meta[property="og:image"]`).First().AttrOr("content", "")
	c.JSON(http.StatusOK, Thumbnail{
		URL: url,
	})

	return nil
}
