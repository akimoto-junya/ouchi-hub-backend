package repository

import (
	"context"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/infra/db"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/model"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type storage struct {
	pool    *pgxpool.Pool
	queries *db.Queries
}

func NewStorage(pool *pgxpool.Pool) repository.Storage {
	return &storage{
		pool:    pool,
		queries: db.New(pool),
	}
}

func (s *storage) Create(ctx context.Context, storage *model.Storage) error {
	if err := s.queries.CreateStorage(ctx, db.CreateStorageParams{
		ID:      storage.ID.String(),
		OwnerID: storage.OwnerID.String(),
	}); err != nil {
		return err
	}
	return nil
}

func (s *storage) GetByOwnerID(ctx context.Context, ownerID uuid.UUID) (*model.Storage, error) {
	storage, err := s.queries.FindStorageByOwnerID(ctx, ownerID.String())
	if err != nil {
		return nil, err
	}
	return &model.Storage{
		ID:      uuid.MustParse(storage.ID),
		OwnerID: uuid.MustParse(storage.OwnerID),
	}, nil
}

func (s *storage) Delete(ctx context.Context, ownerID uuid.UUID) error {
	if err := s.queries.DeleteStorage(ctx, ownerID.String()); err != nil {
		return err
	}
	return nil
}
