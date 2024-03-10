package model

import "github.com/google/uuid"

type Work struct {
	ID           uuid.UUID
	Title        string
	Category     Category
	Group        Group
	IsRestricted bool
}
