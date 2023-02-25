package domain

import (
	"stockin-api/queries"
	"time"
)

type ItemParams struct {
	Title     string
	URL       string
	Thumbnail string
}

func (params *ItemParams) BuildForInsert() *queries.InsertItemParams {
	id := NewUUID()
	now := time.Now()

	return &queries.InsertItemParams{
		ID:        id,
		Title:     params.Title,
		Url:       params.URL,
		Thumbnail: params.Thumbnail,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (params *ItemParams) BuildForUpdate(id UUID) *queries.UpdateItemParams {
	now := time.Now()

	return &queries.UpdateItemParams{
		ID:        id,
		Title:     params.Title,
		Url:       params.URL,
		Thumbnail: params.Thumbnail,
		UpdatedAt: now,
	}
}

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
