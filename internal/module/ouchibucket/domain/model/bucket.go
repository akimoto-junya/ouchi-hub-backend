package model

import "github.com/google/uuid"

type Bucket struct {
	ID              uuid.UUID
	Name            string
	StorageID       uuid.UUID
	RootDirectoryID *uuid.UUID
}

func NewBucket(name string, storageID uuid.UUID) *Bucket {
	return &Bucket{
		ID:              uuid.New(),
		Name:            name,
		StorageID:       storageID,
		RootDirectoryID: nil,
	}
}
