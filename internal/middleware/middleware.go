package middleware

import (
	"encoding/json"
	"io"
	"log/slog"
	"os"

	"github.com/angelvargass/go-api/internal/utils"
	"github.com/gin-gonic/gin"
)

func formatLogEntry(params gin.LogFormatterParams) string {
	log := make(map[string]interface{})

	log["client_ip"] = params.ClientIP
	log["status_code"] = params.StatusCode
	log["path"] = params.Path
	log["method"] = params.Method
	log["start_time"] = params.TimeStamp
	log["remote_addr"] = params.ClientIP
	log["response_time"] = params.Latency.String()

	s, err := json.Marshal(log)
	utils.HandleError(slog.Default(), "error marshalling log entry", err)
	return string(s) + "\n"
}

func JSONLoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(
		func(params gin.LogFormatterParams) string {
			return formatLogEntry(params)
		},
	)
}

func JSONLoggerWriter(logFile *os.File) gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Output: io.MultiWriter(logFile, os.Stdout),
		Formatter: func(params gin.LogFormatterParams) string {
			return formatLogEntry(params)
		},
	})
}
