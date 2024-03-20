package filesystem

import (
	"errors"
	"os"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/config"
)

var ErrStorageNotFound = errors.New("storage dir does not exists")

type OuchiHubStorage interface {
}

type storage struct{}

func NewOuchiHubStorage() (OuchiHubStorage, error) {
	path := config.MustGetStoragePath()
	if d, err := os.Stat(path); err == nil || !d.IsDir() {
		return nil, ErrStorageNotFound
	}
	return &storage{}, nil
}
