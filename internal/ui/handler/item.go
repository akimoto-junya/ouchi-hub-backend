package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/infra/filesystem"
	itemv1 "github.com/akimoto-junya/ouchi-hub-backend/internal/ui/grpc/ouchihub/item/v1"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/ui/grpc/ouchihub/item/v1/itemv1connect"
	"google.golang.org/protobuf/types/known/emptypb"
)

type itemHandler struct {
}

func NewItemHandler(storage filesystem.OuchiHubStorage) itemv1connect.ItemServiceHandler {
	return &itemHandler{}
}

func (h *itemHandler) GetDirectory(ctx context.Context, req *connect.Request[itemv1.GetDirectoryRequest]) (*connect.Response[itemv1.GetDirectoryResponse], error) {
	return nil, nil
}

func (h *itemHandler) MoveItem(ctx context.Context, req *connect.Request[itemv1.MoveItemRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, nil
}
