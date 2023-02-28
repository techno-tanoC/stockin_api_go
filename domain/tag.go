package domain

import (
	"stockin-api/queries"
	"time"
)

type Tag struct {
	ID        UUID      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func TagFromModel(model *queries.Tag) *Tag {
	return &Tag{
		ID:        model.ID,
		Name:      model.Name,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
