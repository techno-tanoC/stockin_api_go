package domain

import (
	"stockin-api/queries"
	"time"
)

type TagParams struct {
	Name string
}

func (params *TagParams) BuildForInsert() *queries.InsertTagParams {
	id := NewUUID()
	now := time.Now()

	return &queries.InsertTagParams{
		ID:        id,
		Name:      params.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (params *TagParams) BuildForUpdate(id UUID) *queries.UpdateTagParams {
	now := time.Now()

	return &queries.UpdateTagParams{
		ID:        id,
		Name:      params.Name,
		UpdatedAt: now,
	}
}

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
