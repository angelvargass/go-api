package middleware

import (
	"encoding/json"
	"io"
	"log/slog"
	"os"

	"github.com/angelvargass/go-api/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func formatLogEntry(params gin.LogFormatterParams) string {
	log := make(map[string]interface{})

	log["client_ip"] = params.ClientIP
	log["status_code"] = params.StatusCode
	log["path"] = params.Path
	log["method"] = params.Method
	log["start_time"] = params.TimeStamp
	log["request_id"] = params.Keys["request_id"]

	s, err := json.Marshal(log)
	utils.HandleError(slog.Default(), "error marshalling log entry", err)
	return string(s) + "\n"
}

func generateUUID() string {
	return uuid.New().String()
}

func JSONLogger(logFile *os.File) gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Output: io.MultiWriter(logFile, os.Stdout),
		Formatter: func(params gin.LogFormatterParams) string {
			return formatLogEntry(params)
		},
	})
}

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("request_id", generateUUID())
		c.Next()
	}
}
