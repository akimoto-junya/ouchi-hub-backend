package usecase

import (
	"context"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/model"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/repository"
	"github.com/google/uuid"
)

type Storage interface {
	GetStorage(ctx context.Context, ownerID uuid.UUID) (*model.Storage, error)
	CreateStorage(ctx context.Context, storage *model.Storage) error
	DeleteStorage(ctx context.Context, ownerID uuid.UUID) error
}

type storage struct {
	repo repository.Storage
}

func NewStorage(repo repository.Storage) Storage {
	return &storage{repo: repo}
}

func (s *storage) GetStorage(ctx context.Context, ownerID uuid.UUID) (*model.Storage, error) {
	m, err := s.repo.GetByOwnerID(ctx, ownerID)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (s *storage) CreateStorage(ctx context.Context, storage *model.Storage) error {
	if err := s.repo.Create(ctx, storage); err != nil {
		return err
	}
	return nil
}

func (s *storage) DeleteStorage(ctx context.Context, ownerID uuid.UUID) error {
	if err := s.repo.Delete(ctx, ownerID); err != nil {
		return err
	}
	return nil
}
