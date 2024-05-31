package model

import "github.com/google/uuid"

type Work struct {
	ID           uuid.UUID
	Title        string
	Category     Category
	Maker        Maker
	IsRestricted bool
	RootDirID    uuid.UUID
}
