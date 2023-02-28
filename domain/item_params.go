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
