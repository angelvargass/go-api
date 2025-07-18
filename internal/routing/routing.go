package routing

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/angelvargass/go-api/internal/health"
	"github.com/angelvargass/go-api/internal/middleware"
	"github.com/angelvargass/go-api/internal/ping"
	"github.com/angelvargass/go-api/internal/sample"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(ctx context.Context, logger *slog.Logger, logFile *os.File, pool *pgxpool.Pool) *Routing {
	return &Routing{
		ctx:    ctx,
		Engine: initGinInstance(logFile),
		Logger: logger,
		DBConn: pool,
	}
}

func initGinInstance(logFile *os.File) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestIDMiddleware())
	r.Use(middleware.JSONLogger(logFile))
	return r
}

func (r *Routing) InitRoutes() {
	r.Logger.Info("initializing routes")
	sampleInstance := sample.New(r.ctx, r.DBConn, r.Logger)

	v1 := r.Engine.Group("/v1")

	healthRoute := v1.Group("/healthz")
	{
		healthRoute.GET("", health.Health)
	}

	pingRoutes := v1.Group("/ping")
	{
		pingRoutes.GET("", ping.Pong)
	}

	sampleRoutes := v1.Group("/sample")
	{
		sampleRoutes.GET("", sampleInstance.GetSamples)
	}

	routes := r.Engine.Routes()
	for _, route := range routes {
		r.Logger.Info(fmt.Sprintf("route: %s %s", route.Method, route.Path))
	}
}
