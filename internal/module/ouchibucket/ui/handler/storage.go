package handler

import (
	"context"

	"connectrpc.com/connect"
	storagev1 "github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchibucket/storage/v1"
	"github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchibucket/storage/v1/storagev1connect"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/model"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/usecase"
	"github.com/google/uuid"
)

type storageHandler struct {
	storage usecase.Storage
}

func NewStorageHandler(s usecase.Storage) storagev1connect.StorageServiceHandler {
	return &storageHandler{
		storage: s,
	}
}

func (h *storageHandler) GetStorage(ctx context.Context, req *connect.Request[storagev1.GetStorageRequest]) (*connect.Response[storagev1.GetStorageResponse], error) {
	ownerID, err := uuid.Parse(req.Msg.GetOwnerId())
	if err != nil {
		return nil, err
	}
	storage, err := h.storage.GetStorage(ctx, ownerID)
	if err != nil {
		return nil, err
	}
	message := &storagev1.Storage{
		Id:      storage.ID.String(),
		OwnerId: storage.OwnerID.String(),
	}
	return connect.NewResponse(&storagev1.GetStorageResponse{Storage: message}), nil
}

func (h *storageHandler) CreateStorage(ctx context.Context, req *connect.Request[storagev1.CreateStorageRequest]) (*connect.Response[storagev1.CreateStorageResponse], error) {
	storage, err := model.NewStorage(req.Msg.GetOwnerId())
	if err != nil {
		return nil, err
	}
	if err := h.storage.CreateStorage(ctx, storage); err != nil {
		return nil, err // TODO:
	}
	return connect.NewResponse(&storagev1.CreateStorageResponse{Id: storage.ID.String()}), nil
}

func (h *storageHandler) DeleteStorage(ctx context.Context, req *connect.Request[storagev1.DeleteStorageRequest]) (*connect.Response[storagev1.DeleteStorageResponse], error) {
	ownerID, err := uuid.Parse(req.Msg.GetOwnerId())
	if err != nil {
		return nil, err
	}
	if err := h.storage.DeleteStorage(ctx, ownerID); err != nil {
		return nil, err
	}
	return connect.NewResponse(&storagev1.DeleteStorageResponse{}), nil
}
