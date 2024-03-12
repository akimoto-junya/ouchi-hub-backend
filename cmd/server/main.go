package main

import (
	"fmt"
	"net/http"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/ui"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	s := ui.NewServer()
	fmt.Println("starting server")
	http.ListenAndServe(
		"0.0.0.0:9000",
		h2c.NewHandler(s, &http2.Server{}),
	)
}
