package repository

import (
	"context"
	"errors"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/model"
	"github.com/google/uuid"
)

var (
	ErrBucketNotFound = errors.New("bucket not found")
)

type Bucket interface {
	Create(ctx context.Context, bucket *model.Bucket) error
	GetByID(ctx context.Context, bucketID uuid.UUID) (*model.Bucket, error)
	ListByStorageID(ctx context.Context, storageID uuid.UUID) ([]*model.Bucket, error)
	Delete(ctx context.Context, bucketID uuid.UUID) error
}
