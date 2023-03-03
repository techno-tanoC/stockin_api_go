package handlers

import (
	"context"
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
		return clientError(c, "invalid id error")
	}

	items, err := h.usecase.Find(ctx, uuid)
	if err != nil {
		return clientError(c, "item not found error")
	}

	return ok(c, items)
}
