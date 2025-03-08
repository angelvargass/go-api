package utils

import (
	"log/slog"
	"os"
)

func HandleError(logger *slog.Logger, errorMessage string, err error) {
	if err != nil {
		logger.Error(errorMessage, slog.String("error", err.Error()))
		os.Exit(1)
	}
}
