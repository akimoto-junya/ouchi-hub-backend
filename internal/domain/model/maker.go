package model

import "github.com/google/uuid"

type Maker struct {
	ID           uuid.UUID
	Name         string
	IsRestricted bool
}
