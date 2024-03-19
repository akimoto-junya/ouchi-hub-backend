package repository

import (
	"context"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/domain/model"
	"github.com/google/uuid"
)

type Item interface {
	FindDirectChildren(ctx context.Context, targetDirID uuid.UUID) ([]model.Item, error)
	//FindRecurcive(ctx context.Context, targetDirID uuid.UUID) (*model.Directory, error)
	CreateRootDirectory(ctx context.Context, root model.Directory, workID uuid.UUID) error
	DeleteItemTree(ctx context.Context, workID, id uuid.UUID) error
	InsertItemTree(ctx context.Context, childRoot model.Item, workID, targetDirID uuid.UUID) error
}
