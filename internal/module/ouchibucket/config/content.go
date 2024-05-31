package config

import (
	"log/slog"
	"os"
)

func MustGetContentURLRoot() string {
	p := os.Getenv("OUCHI_HUB_CONTENT_ROOT")
	if p == "" {
		slog.Error("failed to get content root")
		os.Exit(1)
	}
	return p
}
