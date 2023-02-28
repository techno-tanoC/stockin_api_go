package domain

import (
	"stockin-api/queries"
	"time"
)

type Item struct {
	ID        UUID      `json:"id"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Thumbnail string    `json:"thumbnail"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
