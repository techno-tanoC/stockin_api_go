package domain

import (
	"stockin-api/queries"
	"time"
)

type Tag struct {
	ID        UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func TagFromModel(model *queries.Tag) *Tag {
	return &Tag{
		ID:        model.ID,
		Name:      model.Name,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
