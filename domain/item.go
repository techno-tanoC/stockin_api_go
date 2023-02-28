package domain

import (
	"stockin-api/queries"
	"time"
)

type Item struct {
	ID        UUID
	Title     string
	URL       string
	Thumbnail string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ItemFromModel(model *queries.Item) *Item {
	return &Item{
		ID:        model.ID,
		Title:     model.Title,
		URL:       model.Url,
		Thumbnail: model.Thumbnail,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
