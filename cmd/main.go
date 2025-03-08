package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/angelvargass/go-api/internal/config"
	"github.com/angelvargass/go-api/internal/logger"
	"github.com/angelvargass/go-api/internal/middleware"
	"github.com/angelvargass/go-api/internal/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.New()
	utils.HandleError(slog.Default(), "error instanciating config instance", err)

	logFile, err := os.OpenFile(config.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	utils.HandleError(slog.Default(), "error opening log file", err)

	slog.Info("creating logger instance")
	logger := logger.New(config.LogLevel, logFile)
	slog.SetDefault(logger)

	slog.Info("creating gin instance")
	r := InitGinInstance(logFile)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err = r.Run()
	if err != nil {
		slog.Error(err.Error())
	}
}

func InitGinInstance(logFile *os.File) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.JSONLoggerMiddleware())
	r.Use(middleware.JSONLoggerWriter(logFile))
	return r
}
