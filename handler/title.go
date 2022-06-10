package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func TitleQuery(c echo.Context) error {
	u := new(URL)
	err := c.Bind(u)
	if err != nil {
		return fmt.Errorf("bind error: %w", err)
	}

	doc, err := fetchDocument(u.URL)
	if err != nil {
		return fmt.Errorf("fetch document error: %w", err)
	}

	title := doc.Find("html > head > title").First().Text()
	c.JSON(http.StatusOK, Title{
		Title: title,
	})
	return nil
}
