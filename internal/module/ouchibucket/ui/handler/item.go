package handler

import (
	"context"

	"connectrpc.com/connect"
	itemv1 "github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchibucket/item/v1"
	"github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchibucket/item/v1/itemv1connect"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/usecase"
	"github.com/google/uuid"
)

type itemHandler struct {
	item usecase.Item
}

func NewItemHandler(i usecase.Item) itemv1connect.ItemServiceHandler {
	return &itemHandler{
		item: i,
	}
}

func (h *itemHandler) InsertItem(ctx context.Context, req *connect.Request[itemv1.InsertItemRequest]) (*connect.Response[itemv1.InsertItemResponse], error) {
	// TODO:
	return nil, nil
}

func (h *itemHandler) GetDirectory(ctx context.Context, req *connect.Request[itemv1.GetDirectoryRequest]) (*connect.Response[itemv1.GetDirectoryResponse], error) {
	itemID, err := uuid.Parse(req.Msg.GetItemId())
	if err != nil {
		return nil, err
	}
	dir, err := h.item.GetDirectory(ctx, itemID, int(req.Msg.GetDepth()))
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&itemv1.GetDirectoryResponse{
		Directory: itemv1.ToDirectoryMessage(dir),
	}), nil
}

func (h *itemHandler) MoveItem(ctx context.Context, req *connect.Request[itemv1.MoveItemRequest]) (*connect.Response[itemv1.MoveItemResponse], error) {
	// TODO:
	return nil, nil
}
