package domain

import (
	"fmt"
	"stockin-api/queries"
)

type ItemRangeParams struct {
	Before UUID  `query:"before"`
	Limit  int32 `query:"limit"`
}

func DefaultItemRangePrams() *ItemRangeParams {
	return &ItemRangeParams{
		Before: maxUUID,
		Limit:  50,
	}
}

func (params *ItemRangeParams) Validate() error {
	if params.Limit < 1 || 50 < params.Limit {
		return fmt.Errorf("invalid limit error: %d", params.Limit)
	}
	return nil
}

func (params *ItemRangeParams) Build() *queries.FindItemsByRangeParams {
	return &queries.FindItemsByRangeParams{
		ID:    params.Before,
		Limit: params.Limit,
	}
}
