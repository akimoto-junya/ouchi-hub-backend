package handler

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	greetv1 "github.com/akimoto-junya/ouchi-hub-backend/internal/ui/grpc/ouchihub/greet/v1"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/ui/grpc/ouchihub/greet/v1/greetv1connect"
)

type greetHandler struct{}

func NewGreetHandler() greetv1connect.GreetServiceHandler {
	return &greetHandler{}
}

func (s *greetHandler) Greet(ctx context.Context, req *connect.Request[greetv1.GreetRequest]) (*connect.Response[greetv1.GreetResponse], error) {
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
