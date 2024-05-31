package repository

import (
	"context"
	"errors"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/model"
	"github.com/google/uuid"
)

var (
	ErrEmptyRootDir = errors.New("empty root dir")
)

type Item interface {
	GetDirectory(ctx context.Context, targetDirID uuid.UUID, depth int) (*model.Directory, error)
	CreateRootDirectory(ctx context.Context, root *model.Directory, bucketID uuid.UUID) error
	InsertItemTree(ctx context.Context, childRoot model.Item, bucketID, targetDirID uuid.UUID) error
	DeleteItemTree(ctx context.Context, bucketID, id uuid.UUID) error
	DeleteByBucketID(ctx context.Context, bucketID uuid.UUID) error
	// deprecated
	LoadRootDirectory(ctx context.Context, bucketID uuid.UUID, relPath string) (*model.Directory, error)
}
