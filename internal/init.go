package internal

import (
	"math/rand"
	"stockin-api/domain"

	"github.com/brianvoe/gofakeit/v6"
)

func init() {
	gofakeit.AddFuncLookup("uuidv7", gofakeit.Info{
		Category:    "custom",
		Description: "UUID v7",
		Example:     "0186a750-5bff-7024-bb74-1c1b5b58ec66",
		Output:      "UUID",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			return domain.NewUUID(), nil
		},
	})
}
