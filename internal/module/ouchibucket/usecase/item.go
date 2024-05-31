package usecase

import (
	"context"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/model"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/repository"
	"github.com/google/uuid"
)

type Item interface {
	GetDirectory(ctx context.Context, itemID uuid.UUID, depth int) (*model.Directory, error)
}

type item struct {
	repo repository.Item
}

func NewItem(repo repository.Item) Item {
	return &item{repo: repo}
}

func (i *item) GetDirectory(ctx context.Context, itemID uuid.UUID, depth int) (*model.Directory, error) {
	return i.repo.GetDirectory(ctx, itemID, depth)
}
