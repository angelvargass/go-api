package routing

import (
	"log/slog"
	"os"

	"github.com/angelvargass/go-api/internal/middleware"
	"github.com/angelvargass/go-api/internal/ping"
	"github.com/gin-gonic/gin"
)

func New(logger *slog.Logger, logFile *os.File) *Routing {
	return &Routing{
		Engine: initGinInstance(logFile),
		Logger: logger,
	}
}

func initGinInstance(logFile *os.File) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.JSONLoggerMiddleware())
	r.Use(middleware.JSONLoggerWriter(logFile))
	return r
}

func (r *Routing) InitRoutes() {
	v1 := r.Engine.Group("/v1")

	pingRoutes := v1.Group("/ping")
	{
		pingRoutes.GET("", ping.Pong)
	}
}
