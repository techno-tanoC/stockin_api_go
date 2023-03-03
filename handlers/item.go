package handlers

import (
	"context"
	"net/http"
	"stockin-api/usecases"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
)

type ItemHandler struct {
	usecase *usecases.ItemUsecaseImpl
}

func NewItemHandler(usecase *usecases.ItemUsecaseImpl) *ItemHandler {
	return &ItemHandler{usecase}
}

func (h *ItemHandler) Find(c echo.Context) error {
	ctx := context.Background()

	id := c.Param("id")
	uuid, err := uuid.FromString(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id error")
	}

	items, err := h.usecase.Find(ctx, uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "item not found error")
	}

	return c.JSON(http.StatusOK, items)
}
