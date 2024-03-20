package handler

import (
	"context"

	"connectrpc.com/connect"
	makerv1 "github.com/akimoto-junya/ouchi-hub-backend/internal/ui/grpc/ouchihub/maker/v1"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/ui/grpc/ouchihub/maker/v1/makerv1connect"
)

type makerHandler struct{}

func NewMakerHandler() makerv1connect.MakerServiceHandler {
	return &makerHandler{}
}

func (h *makerHandler) ListMakers(ctx context.Context, req *connect.Request[makerv1.ListMakersRequest]) (*connect.Response[makerv1.ListMakersResponse], error) {
	return nil, nil
}
