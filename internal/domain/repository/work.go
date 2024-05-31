package repository

import (
	"context"

	"github.com/google/uuid"
)

type Work interface {
	// DEPRECATED: replace with upload feature
	Load(ctx context.Context, workID uuid.UUID, relPath string) error
}
