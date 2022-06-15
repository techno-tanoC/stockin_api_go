package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type URL struct {
	URL string `json:"url"`
}

type Data struct {
	Data interface{} `json:"data"`
}

func json(c echo.Context, d interface{}) error {
	return c.JSON(http.StatusOK, Data{d})
}
