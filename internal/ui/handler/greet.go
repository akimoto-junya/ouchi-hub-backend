package handler

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	greetv1 "github.com/akimoto-junya/ouchi-hub-backend/internal/ui/grpc/ouchihub/greet/v1"
)

type GreetHandler struct{}

func (s *GreetHandler) Greet(ctx context.Context, req *connect.Request[greetv1.GreetRequest]) (*connect.Response[greetv1.GreetResponse], error) {
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
