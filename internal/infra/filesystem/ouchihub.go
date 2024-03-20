package filesystem

import (
	"errors"
	"fmt"
	"os"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/config"
)

var ErrStorageNotFound = errors.New("storage dir does not exists")

type OuchiHubStorage interface {
}

type storage struct{}

func NewOuchiHubStorage() (OuchiHubStorage, error) {
	path := config.MustGetStoragePath()
	if d, err := os.Stat(path); err != nil || !d.IsDir() {
		if err == nil {
			err = fmt.Errorf("%s is not directory, mode (%s)", d.Name(), d.Mode().String())
		}
		return nil, fmt.Errorf("%w: %w", ErrStorageNotFound, err)
	}
	return &storage{}, nil
}
