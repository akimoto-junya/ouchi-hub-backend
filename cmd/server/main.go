package main

import (
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/ui"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	if err := filepath.WalkDir("/data", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	}); err != nil {
		slog.Error(err.Error())
	}

	s := ui.NewServer()
	slog.Info("starting server")
	if err := http.ListenAndServe(
		"0.0.0.0:9000",
		h2c.NewHandler(s, &http2.Server{}),
	); err != nil {
		slog.Error("failed to start server")
	}
}
