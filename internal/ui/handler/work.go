package handler

import (
	"context"

	"connectrpc.com/connect"
	workv1 "github.com/akimoto-junya/ouchi-hub-backend/internal/ui/grpc/ouchihub/work/v1"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/ui/grpc/ouchihub/work/v1/workv1connect"
)

type workHandler struct{}

func NewWorkHandler() workv1connect.WorkServiceHandler {
	return &workHandler{}
}

func (h *workHandler) ListWorks(ctx context.Context, req *connect.Request[workv1.ListWorksRequest]) (*connect.Response[workv1.ListWorksResponse], error) {
	return nil, nil
}

func (h *workHandler) CreateWork(ctx context.Context, req *connect.Request[workv1.CreateWorkRequest]) (*connect.Response[workv1.CreateWorkResponse], error) {
	return nil, nil
}
