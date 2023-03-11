package handlers

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/techno-tanoC/sins"
)

type (
	URL struct {
		URL string `json:"url"`
	}

	Title struct {
		Title string `json:"title"`
	}

	Thumbnail struct {
		Thumbnail string `json:"thumbnail"`
	}
)

type QueryHandler struct {
	title sins.TitleFinder
	thumb sins.ThumbnailFinder
}

func NewQueryHandler(fetcher sins.Fetcher) *QueryHandler {
	title := sins.NewTitleFinder(fetcher)
	thumb := sins.NewThumbnailFinder(fetcher)
	return &QueryHandler{
		title: title,
		thumb: thumb,
	}
}

func (h *QueryHandler) Title(c echo.Context) error {
	params := new(URL)
	err := c.Bind(&params)
	if err != nil {
		return clientError(c, "invalid params error")
	}

	title, err := h.title.Find(params.URL)
	if err != nil {
		return serverError(c, "internal server error")
	}
	trimmed := strings.TrimSpace(title)

	return ok(c, Title{Title: trimmed})
}

func (h *QueryHandler) Thumbnail(c echo.Context) error {
	params := new(URL)
	err := c.Bind(&params)
	if err != nil {
		return clientError(c, "invalid params error")
	}

	thumb, err := h.thumb.Find(params.URL)
	if err != nil {
		return serverError(c, "internal server error")
	}

	return ok(c, Thumbnail{Thumbnail: thumb})
}
