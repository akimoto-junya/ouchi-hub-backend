package repository

import (
	"context"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/model"
	"github.com/google/uuid"
)

type Storage interface {
	Create(ctx context.Context, storage *model.Storage) error
	GetByOwnerID(ctx context.Context, ownerID uuid.UUID) (*model.Storage, error)
	Delete(ctx context.Context, ownerID uuid.UUID) error
}
