package filesystem

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/config"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/model"
	"github.com/google/uuid"
)

var (
	ErrStorageNotFound = errors.New("storage dir does not exists")
	ErrRootDirNotFound = errors.New("root dir does not exists")
	ErrDirAccess       = errors.New("directory access denied")
)

type OuchiHubStorage interface {
	FindWorkDir(bucketID uuid.UUID, path string) (*model.Directory, error)
}

type storage struct {
	rootDir string
}

func NewOuchiHubStorage() (OuchiHubStorage, error) {
	path := config.MustGetStoragePath()
	if d, err := os.Stat(path); err != nil || !d.IsDir() {
		if err == nil {
			err = fmt.Errorf("%s is not directory, mode (%s)", d.Name(), d.Mode().String())
		}
		return nil, fmt.Errorf("%w: %w", ErrStorageNotFound, err)
	}
	return &storage{rootDir: path}, nil
}

func (s *storage) FindWorkDir(bucketID uuid.UUID, relPath string) (*model.Directory, error) {
	path, err := filepath.Abs(filepath.Join(s.rootDir, relPath))
	if err != nil {
		return nil, err
	} else if !strings.HasPrefix(path, s.rootDir) {
		return nil, ErrDirAccess
	} else if r, err := os.Stat(path); err != nil || !r.IsDir() {
		if err == nil {
			err = fmt.Errorf("%s is not directory, mode (%s)", r.Name(), r.Mode().String())
		}
		return nil, fmt.Errorf("%w: %w", ErrRootDirNotFound, err)
	}


	dirs := map[string]*model.Directory{}
	if err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// directory
		if d.IsDir() {
			if _, ok := dirs[path]; !ok {
				dirs[path] = &model.Directory{
					ID:           uuid.New(),
					Name:         filepath.Base(path),
					Files:        []*model.File{},
					Directorires: []*model.Directory{},
				}
			}
			return nil
		}
		// file
		if _, ok := dirs[filepath.Dir(path)]; !ok {
			dirs[path] = &model.Directory{
				ID:           uuid.New(),
				Name:         filepath.Base(path),
				Files:        []*model.File{},
				Directorires: []*model.Directory{},
			}
		}

		itemID := uuid.New()
		file := &model.File{
			ID:   itemID,
			Name: filepath.Base(path),
			Entity: &model.Content{ // IDはセットしない
				ID:       uuid.New(),
				ItemID:   itemID,
				Path:     filepath.Join(relPath, filepath.Base(path)),
				FileType: DetectFileType(path),
			},
		}
		dirs[filepath.Dir(path)].Files = append(dirs[filepath.Dir(path)].Files, file)
		return nil
	}); err != nil {
		return nil, err
	}

	root := dirs[path]
	for k, v := range dirs {
		if k == path {
			continue
		}
		dirs[filepath.Dir(k)].Directorires = append(dirs[filepath.Dir(k)].Directorires, v)
	}
	return root, nil
}
