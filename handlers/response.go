package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Data struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type Empty struct{}

func ok(c echo.Context, data interface{}) error {
	d := Data{data, ""}
	return c.JSON(http.StatusOK, d)
}

func noContent(c echo.Context) error {
	return c.JSON(http.StatusNoContent, "")
}

func raw(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

func clientError(c echo.Context, message string) error {
	d := Data{Empty{}, message}
	return c.JSON(http.StatusBadRequest, d)
}

func serverError(c echo.Context, message string) error {
	d := Data{Empty{}, message}
	return c.JSON(http.StatusInternalServerError, d)
}
