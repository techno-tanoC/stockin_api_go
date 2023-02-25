package domain

import "github.com/gofrs/uuid"

type UUID = uuid.UUID

func NewUUID() UUID {
	return uuid.Must(uuid.NewV7())
}
