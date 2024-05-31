package repository

import (
	"context"
	"errors"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/infra/db"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/model"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/repository"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/util/pointer"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type bucket struct {
	pool    *pgxpool.Pool
	queries *db.Queries
}

func NewBucket(pool *pgxpool.Pool) repository.Bucket {
	return &bucket{
		pool:    pool,
		queries: db.New(pool),
	}
}

func (b *bucket) Create(ctx context.Context, bucket *model.Bucket) error {
	if err := b.queries.CreateBucket(ctx, db.CreateBucketParams{
		ID:        bucket.ID.String(),
		Name:      bucket.Name,
		StorageID: bucket.StorageID.String(),
	}); err != nil {
		return err
	}
	return nil
}

func (b *bucket) GetByID(ctx context.Context, bucketID uuid.UUID) (*model.Bucket, error) {
	dbBucket, err := b.queries.FindBucketByID(ctx, bucketID.String())
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.ErrBucketNotFound
		}
		return nil, err
	}

	return &model.Bucket{
		ID:              uuid.MustParse(dbBucket.ID),
		Name:            dbBucket.Name,
		StorageID:       uuid.MustParse(dbBucket.StorageID),
		RootDirectoryID: pointer.ToPtr(uuid.MustParse(dbBucket.RootDirectoryID)),
	}, nil
}

func (b *bucket) ListByStorageID(ctx context.Context, storageID uuid.UUID) ([]*model.Bucket, error) {
	dbBuckets, err := b.queries.FindBuckets(ctx, storageID.String())
	if err != nil {
		return nil, err
	}

	ret := make([]*model.Bucket, len(dbBuckets))
	for i, dbBucket := range dbBuckets {
		ret[i] = &model.Bucket{
			ID:              uuid.MustParse(dbBucket.ID),
			Name:            dbBucket.Name,
			StorageID:       uuid.MustParse(dbBucket.StorageID),
			RootDirectoryID: pointer.ToPtr(uuid.MustParse(dbBucket.RootDirectoryID)),
		}
	}
	return ret, nil
}

func (b *bucket) Delete(ctx context.Context, bucketID uuid.UUID) error {
	if err := b.queries.DeleteBucket(ctx, bucketID.String()); err != nil {
		return err
	}
	return nil
}
