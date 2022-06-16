package handler

import (
	"context"
	"fmt"
	"stockin/domain"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ItemParams struct {
	Title     string `json:"title"`
	URL       string `json:"url"`
	Thumbnail string `json:"thumbnail"`
}

func ItemIndex(db domain.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		fromStr := c.QueryParam("from")
		from, err := strconv.ParseInt(fromStr, 10, 64)
		if err != nil {
			return fmt.Errorf("parse from error: %w", err)
		}

		limitStr := c.QueryParam("limit")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return fmt.Errorf("parse from error: %w", err)
		}

		ctx := context.Background()
		items, err := domain.ItemIndex(ctx, db, from, limit)
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
		params := new(ItemParams)
		err := c.Bind(params)
		if err != nil {
			return fmt.Errorf("bind error: %w", err)
		}

		ctx := context.Background()
		item, err := domain.ItemCreate(ctx, db, params.Title, params.URL, params.Thumbnail)
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
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return fmt.Errorf("invalid id error: %w", err)
		}

		params := new(ItemParams)
		err = c.Bind(params)
		if err != nil {
			return fmt.Errorf("bind error: %w", err)
		}

		ctx := context.Background()
		item, err := domain.ItemUpdate(ctx, db, id, params.Title, params.URL, params.Thumbnail)
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
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return fmt.Errorf("invalid id error: %w", err)
		}

		ctx := context.Background()
		err = domain.ItemDelete(ctx, db, id)
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