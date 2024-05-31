package ui

import (
	"net/http"

	"github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchihub/category/v1/categoryv1connect"
	"github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchihub/maker/v1/makerv1connect"
	"github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchihub/work/v1/workv1connect"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/di"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/ui/handler"
)

func RegisterHandlers(s *http.ServeMux, d *di.Dependencies) {
	var p string
	var h http.Handler
	// category
	p, h = categoryv1connect.NewCategoryServiceHandler(handler.NewCategoryHandler())
	s.Handle(p, h)
	// maker
	p, h = makerv1connect.NewMakerServiceHandler(handler.NewMakerHandler())
	s.Handle(p, h)
	// work
	p, h = workv1connect.NewWorkServiceHandler(handler.NewWorkHandler())
	s.Handle(p, h)
}
