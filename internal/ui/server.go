package ui

import (
	"net/http"

	"connectrpc.com/grpcreflect"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/ui/grpc/ouchihub/greet/v1/greetv1connect"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/ui/handler"
)

func NewServer() *http.ServeMux {
	s := http.NewServeMux()

	// server reflection
	reflecter := grpcreflect.NewStaticReflector(
		greetv1connect.GreetServiceName,
	)
	s.Handle(grpcreflect.NewHandlerV1(reflecter))
	s.Handle(grpcreflect.NewHandlerV1Alpha(reflecter))

	// register handler
	var p string
	var h http.Handler
	p, h = greetv1connect.NewGreetServiceHandler(&handler.GreetHandler{})
	s.Handle(p, h)
	return s
}
