package handlers

import (
	"context"
	"stockin-api/domain"
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
		return serverError(c, "internal server error")
	}

	return ok(c, items)
}

func (h *ItemHandler) FindByRange(c echo.Context) error {
	ctx := context.Background()

	params := domain.DefaultItemRangePrams()
	err := c.Bind(params)
	if err != nil {
		return clientError(c, "invalid params error")
	}
	err = params.Validate()
	if err != nil {
		return clientError(c, "invalid params error")
	}

	items, err := h.usecase.FindByRange(ctx, params)
	if err != nil {
		return serverError(c, "internal server error")
	}

	return ok(c, items)
}

func (h *ItemHandler) Create(c echo.Context) error {
	ctx := context.Background()

	params := new(domain.ItemParams)
	err := c.Bind(params)
	if err != nil {
		return clientError(c, "invalid params error")
	}

	item, err := h.usecase.Create(ctx, params)
	if err != nil {
		return serverError(c, "internal server error")
	}

	return ok(c, item)
}

func (h *ItemHandler) Update(c echo.Context) error {
	ctx := context.Background()

	id := c.Param("id")
	uuid, err := uuid.FromString(id)
	if err != nil {
		return clientError(c, "invalid id error")
	}

	params := new(domain.ItemParams)
	err = c.Bind(params)
	if err != nil {
		return clientError(c, "invalid params error")
	}

	item, err := h.usecase.Update(ctx, uuid, *params)
	if err != nil {
		return serverError(c, "internal server error")
	}

	return ok(c, item)
}

func (h *ItemHandler) Delete(c echo.Context) error {
	ctx := context.Background()

	id := c.Param("id")
	uuid, err := uuid.FromString(id)
	if err != nil {
		return clientError(c, "invalid id error")
	}

	err = h.usecase.Delete(ctx, uuid)
	if err != nil {
		return serverError(c, "internal server error")
	}

	return noContent(c)
}
