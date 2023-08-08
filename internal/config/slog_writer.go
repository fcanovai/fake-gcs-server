package config

import (
	"context"

	"golang.org/x/exp/slog"
)

type slogWriter struct {
	logger *slog.Logger
	level  slog.Level
}

func (w *slogWriter) Write(p []byte) (n int, err error) {
	w.logger.Log(context.Background(), w.level, string(p))
	return len(p), nil
}
