package logger

import (
	"io"
	"log/slog"
	"os"
	"strings"
)

func New(logLevel string, logFile *os.File) *slog.Logger {
	slog.Info("creating new logger with log level", slog.String("level", logLevel))
	level := new(slog.LevelVar)

	switch strings.ToLower(logLevel) {
	case "debug":
		level.Set(slog.LevelDebug)
	case "info":
		level.Set(slog.LevelInfo)
	case "warn":
		level.Set(slog.LevelWarn)
	case "error":
		level.Set(slog.LevelError)
	default:
		level.Set(slog.LevelInfo)
	}

	cfg := &slog.HandlerOptions{
		Level: level,
	}

	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logFile, os.Stdout), cfg))
	return logger
}
