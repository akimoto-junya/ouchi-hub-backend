package ui

import (
	"net/http"

	"connectrpc.com/grpcreflect"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/ui/grpc/ouchihub/category/v1/categoryv1connect"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/ui/grpc/ouchihub/greet/v1/greetv1connect"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/ui/grpc/ouchihub/item/v1/itemv1connect"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/ui/grpc/ouchihub/maker/v1/makerv1connect"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/ui/grpc/ouchihub/work/v1/workv1connect"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/ui/handler"
)

func NewServer() *http.ServeMux {
	s := http.NewServeMux()

	// server reflection
	reflecter := grpcreflect.NewStaticReflector(
		greetv1connect.GreetServiceName,
		categoryv1connect.CategoryServiceName,
		makerv1connect.MakerServiceName,
		itemv1connect.ItemServiceName,
		workv1connect.WorkServiceName,
	)
	s.Handle(grpcreflect.NewHandlerV1(reflecter))
	s.Handle(grpcreflect.NewHandlerV1Alpha(reflecter))

	// register handler
	var p string
	var h http.Handler
	// greet
	p, h = greetv1connect.NewGreetServiceHandler(&handler.GreetHandler{})
	s.Handle(p, h)
	// category
	p, h = categoryv1connect.NewCategoryServiceHandler(&handler.CategoryHandler{})
	s.Handle(p, h)
	return s
}
