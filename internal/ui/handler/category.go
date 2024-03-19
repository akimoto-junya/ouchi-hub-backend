package handler

import (
	"context"

	"connectrpc.com/connect"
	categoryv1 "github.com/akimoto-junya/ouchi-hub-backend/internal/ui/grpc/ouchihub/category/v1"
)

type CategoryHandler struct{}

func (h *CategoryHandler) ListCategories(context.Context, *connect.Request[categoryv1.ListCategoriesRequest]) (*connect.Response[categoryv1.ListCategoriesResponse], error) {
	res := connect.NewResponse(&categoryv1.ListCategoriesResponse{})
	return res, nil
}
