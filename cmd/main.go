package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/angelvargass/go-api/internal/config"
	"github.com/angelvargass/go-api/internal/logger"
	"github.com/angelvargass/go-api/internal/routing"
	"github.com/angelvargass/go-api/internal/utils"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	fmt.Println(ctx)
	defer stop()

	config, err := config.New()
	utils.HandleError(slog.Default(), "error instanciating config instance", err)

	logFile, err := os.OpenFile(config.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	utils.HandleError(slog.Default(), "error opening log file", err)

	slog.Info("creating logger instance")
	logger := logger.New(config.LogLevel, logFile)
	slog.SetDefault(logger)

	slog.Info("creating routing instance")
	routing := routing.New(logger, logFile)

	slog.Info("initializing Gin routes")
	routing.InitRoutes()

	slog.Info("initializing Gin server")
	err = routing.Engine.Run()
	utils.HandleError(logger, "error initializing gin engine", err)
}
