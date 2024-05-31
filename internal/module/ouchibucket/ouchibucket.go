package ouchibucket

import (
	"net/http"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/di"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/ui"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterOuchiBucket(s *http.ServeMux, db *pgxpool.Pool) {
	d := di.NewDependencies(db)
	ui.RegisterHandlers(s, d)
}
