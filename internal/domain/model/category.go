package model

import (
	"github.com/google/uuid"
)

type Category struct {
	ID   uuid.UUID
	Name string
}
