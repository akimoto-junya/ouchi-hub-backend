package handler

import (
	"context"

	"connectrpc.com/connect"
	categoryv1 "github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchihub/category/v1"
	"github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchihub/category/v1/categoryv1connect"
)

type categoryHandler struct{}

func NewCategoryHandler() categoryv1connect.CategoryServiceHandler {
	return &categoryHandler{}
}

func (h *categoryHandler) ListCategories(context.Context, *connect.Request[categoryv1.ListCategoriesRequest]) (*connect.Response[categoryv1.ListCategoriesResponse], error) {
	res := connect.NewResponse(&categoryv1.ListCategoriesResponse{})
	return res, nil
}
