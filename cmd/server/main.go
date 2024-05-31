package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/di"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/ui"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	s := http.NewServeMux()
	ui.RegisterServerReflection(s)

	db, err := pgxpool.New(context.Background(), "postgres://user:password@db:5432/ouchihub")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	dependencies := di.NewDependencies()
	ui.RegisterHandlers(s, dependencies)

	// module
	ouchibucket.RegisterOuchiBucket(s, db)

	slog.Info("starting server")
	if err := http.ListenAndServe(
		"0.0.0.0:9000",
		h2c.NewHandler(s, &http2.Server{}),
	); err != nil {
		slog.Error("failed to start server")
	}
}
