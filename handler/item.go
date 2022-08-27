package handler

import (
	"context"
	"fmt"
	"stockin/domain"
	"stockin/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ItemIndex(db domain.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		before := c.QueryParam("before")
		if before == "" {
			before = "ffffffff-ffff-ffff-ffff-ffffffffffff"
		}

		limitStr := c.QueryParam("limit")
		if limitStr == "" {
			limitStr = "50"
		}
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return fmt.Errorf("parse from error: %w", err)
		}

		ctx := context.Background()
		items, err := domain.ItemIndex(ctx, db, before, limit)
		if err != nil {
			return fmt.Errorf("insert error: %w", err)
		}

		err = json(c, items)
		if err != nil {
			return fmt.Errorf("json error: %w", err)
		}

		return nil
	}
}

func ItemCreate(db domain.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		params := new(domain.ItemParams)
		err := c.Bind(params)
		if err != nil {
			return fmt.Errorf("bind error: %w", err)
		}

		ctx := context.Background()
		item, err := domain.ItemCreate(ctx, db, params)
		if err != nil {
			return fmt.Errorf("insert error: %w", err)
		}

		err = json(c, item)
		if err != nil {
			return fmt.Errorf("json error: %w", err)
		}

		return nil
	}
}

func ItemUpdate(db domain.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		params := new(domain.ItemParams)
		err := c.Bind(params)
		if err != nil {
			return fmt.Errorf("bind error: %w", err)
		}

		ctx := context.Background()
		item, err := domain.ItemUpdate(ctx, db, id, params)
		if err != nil {
			return fmt.Errorf("item update error: %w", err)
		}

		err = json(c, item)
		if err != nil {
			return fmt.Errorf("json error: %w", err)
		}

		return nil
	}
}

func ItemDelete(db domain.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		ctx := context.Background()
		err := domain.ItemDelete(ctx, db, id)
		if err != nil {
			return fmt.Errorf("item delete error: %w", err)
		}

		err = noContent(c)
		if err != nil {
			return fmt.Errorf("no content error: %w", err)
		}

		return nil
	}
}

func ItemExport(db domain.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		items, err := domain.ItemExport(ctx, db)
		if err != nil {
			return fmt.Errorf("item export error: %w", err)
		}

		err = rawJson(c, items)
		if err != nil {
			return fmt.Errorf("json error: %w", err)
		}

		return nil
	}
}

func ItemImport(db domain.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		params := new([]*models.Item)
		err := c.Bind(params)
		if err != nil {
			return fmt.Errorf("bind error: %w", err)
		}

		ctx := context.Background()
		err = domain.ItemImport(ctx, db, *params)
		if err != nil {
			return fmt.Errorf("item import error: %w", err)
		}

		err = noContent(c)
		if err != nil {
			return fmt.Errorf("no content error: %w", err)
		}

		return nil
	}
}
