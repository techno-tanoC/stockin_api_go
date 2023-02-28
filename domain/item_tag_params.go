package domain

import (
	"stockin-api/queries"
	"time"
)

type ItemTagParams struct {
	ItemID UUID `json:"item_id"`
	TagID  UUID `json:"tag_id"`
}

func (params *ItemTagParams) BuildForInsert() *queries.InsertItemTagParams {
	id := NewUUID()
	now := time.Now()

	return &queries.InsertItemTagParams{
		ID:        id,
		ItemID:    params.ItemID,
		TagID:     params.TagID,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
