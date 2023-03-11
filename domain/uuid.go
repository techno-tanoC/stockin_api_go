package domain

import "github.com/gofrs/uuid"

type UUID = uuid.UUID

var maxUUID = uuid.Must(uuid.FromString("ffffffff-ffff-ffff-ffff-ffffffffffff"))

func NewUUID() UUID {
	return uuid.Must(uuid.NewV7())
}
