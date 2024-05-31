package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/model"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/repository"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/usecase/dto"
	utilContext "github.com/akimoto-junya/ouchi-hub-backend/internal/util/context"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/util/transaction"
)

type Bucket interface {
	ListBuckets(ctx context.Context, storageID uuid.UUID) ([]*model.Bucket, error)
	CreateBucket(ctx context.Context, params *dto.CreateBucketParams) (*uuid.UUID, error)
	SyncBucket(ctx context.Context, bucketID uuid.UUID, relPath string) error
}

type bucket struct {
	db          *pgxpool.Pool
	bucketRepo  repository.Bucket
	itemRepo    repository.Item
	contentRepo repository.Content
}

var (
	ErrBucketNotFound = errors.New("bucket not found")
)

func NewBucket(bucketRepo repository.Bucket, itemRepo repository.Item, contentRepo repository.Content, db *pgxpool.Pool) Bucket {
	return &bucket{
		bucketRepo:  bucketRepo,
		itemRepo:    itemRepo,
		contentRepo: contentRepo,
		db:          db,
	}
}

func (b *bucket) ListBuckets(ctx context.Context, storageID uuid.UUID) ([]*model.Bucket, error) {
	return b.bucketRepo.ListByStorageID(ctx, storageID)
}

func (b *bucket) CreateBucket(ctx context.Context, params *dto.CreateBucketParams) (*uuid.UUID, error) {
	tx, finishTx, err := transaction.Transact(ctx, b.db)
	if err != nil {
		return nil, err
	}
	ctx = utilContext.SetTx(ctx, tx)
	modelBucket := model.NewBucket(params.Name, params.StorageID)
	if err := b.bucketRepo.Create(ctx, modelBucket); err != nil {
		return nil, finishTx(err)
	}
	if err := b.itemRepo.CreateRootDirectory(ctx, model.NewDirectory(modelBucket.Name), modelBucket.ID); err != nil {
		return nil, finishTx(err)
	}
	if err := finishTx(nil); err != nil {
		return nil, err
	}
	return &modelBucket.ID, nil
}

func (b *bucket) SyncBucket(ctx context.Context, bucketID uuid.UUID, relPath string) error {
	m, err := b.bucketRepo.GetByID(ctx, bucketID)
	if err != nil {
		if errors.Is(err, repository.ErrBucketNotFound) {
			return ErrBucketNotFound
		}
		return err
	} else if m == nil {
		return ErrBucketNotFound
	}

	tx, finishTx, err := transaction.Transact(ctx, b.db)
	if err != nil {
		return err
	}
	ctx = utilContext.SetTx(ctx, tx)

	if err := b.itemRepo.DeleteByBucketID(ctx, bucketID); err != nil {
		return finishTx(err)
	}
	rootDir, err := b.itemRepo.LoadRootDirectory(ctx, bucketID, relPath)
	if err != nil {
		return finishTx(err)
	}
	if err := b.itemRepo.CreateRootDirectory(ctx, rootDir, bucketID); err != nil {
		return finishTx(err)
	}
	contents := rootDir.ListContentsRecursive()
	if err := b.contentRepo.BatchCreate(ctx, contents); err != nil {
		return finishTx(err)
	}

	return finishTx(nil)
}
