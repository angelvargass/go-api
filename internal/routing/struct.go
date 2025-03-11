package routing

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Routing struct {
	Engine *gin.Engine
	Logger *slog.Logger
}
