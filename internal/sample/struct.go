package sample

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Instance struct {
	ctx    context.Context
	dbConn *pgxpool.Pool
	Logger *slog.Logger
}

type Sample struct {
	ID     int    `json:"id"`
	Sample string `json:"sample"`
}
