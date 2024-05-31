package dto

import "github.com/google/uuid"

type CreateBucketParams struct {
	Name      string
	StorageID uuid.UUID
}
