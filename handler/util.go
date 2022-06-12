package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Data struct {
	Data interface{} `json:"data"`
}

func ok(c echo.Context, d interface{}) error {
	return c.JSON(http.StatusOK, Data{d})
}
