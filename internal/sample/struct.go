package sample

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type Instance struct {
	ctx    context.Context
	dbConn *pgx.Conn
	Logger *slog.Logger
}

type Sample struct {
	ID     int    `json:"id"`
	Sample string `json:"sample"`
}
