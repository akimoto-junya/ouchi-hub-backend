package model

import "github.com/google/uuid"

type Work struct {
	ID           uuid.UUID
	Title        string
	Author       Author
	Category     Category
	Group        Group
	IsRestricted bool
	Root         uuid.UUID
}
