package routing

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Routing struct {
	ctx    context.Context
	Engine *gin.Engine
	Logger *slog.Logger
	DBConn *pgxpool.Pool
}
