package config

import (
	"log/slog"
	"os"
)

func MustGetStoragePath() string {
	p := os.Getenv("OUCHI_HUB_STORAGE")
	if p == "" {
		slog.Error("failed to get storage path")
		os.Exit(1)
	}
	return p
}
