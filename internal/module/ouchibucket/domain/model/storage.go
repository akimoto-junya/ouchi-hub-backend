package model

import "github.com/google/uuid"

type Storage struct {
	ID      uuid.UUID
	OwnerID uuid.UUID
}

func NewStorage(ownerID string) (*Storage, error) {
	ownerUUID, err := uuid.Parse(ownerID)
	if err != nil {
		return nil, err
	}
	return &Storage{
		ID:      uuid.New(),
		OwnerID: ownerUUID,
	}, nil
}
