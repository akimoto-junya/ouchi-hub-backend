package di

import (
	"log/slog"
	"os"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/infra/filesystem"
	infraRepo "github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/infra/repository"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/usecase"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Dependencies struct {
	Storage usecase.Storage
	Bucket  usecase.Bucket
	Item    usecase.Item
	Content usecase.Content
}

func NewDependencies(db *pgxpool.Pool) *Dependencies {
	fileStorage, err := filesystem.NewOuchiHubStorage()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	contentRepo := infraRepo.NewContent(db)
	itemRepo := infraRepo.NewItem(fileStorage, db)
	bucketRepo := infraRepo.NewBucket(db)
	storageRepo := infraRepo.NewStorage(db)

	return &Dependencies{
		Storage: usecase.NewStorage(storageRepo),
		Bucket:  usecase.NewBucket(bucketRepo, itemRepo, contentRepo, db),
		Item:    usecase.NewItem(itemRepo),
		Content: usecase.NewContent(contentRepo),
	}
}
