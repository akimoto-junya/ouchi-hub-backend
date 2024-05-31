package repository

import (
	"context"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/model"
	"github.com/google/uuid"
)

type Content interface {
	Create(ctx context.Context, content *model.Content) error
	BatchCreate(ctx context.Context, contents []*model.Content) error
	ListByItemID(ctx context.Context, itemID uuid.UUID) ([]*model.Content, error)
}
