package sample

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/angelvargass/go-api/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	SampleTableName = "sample"
)

func New(ctx context.Context, dbConn *pgxpool.Pool, logger *slog.Logger) *Instance {
	return &Instance{
		ctx:    ctx,
		dbConn: dbConn,
		Logger: logger,
	}
}

func (i *Instance) GetSamples(c *gin.Context) {
	query := "SELECT * FROM " + SampleTableName
	rows, err := i.dbConn.Query(i.ctx, query)
	utils.HandleError(i.Logger, "error querying database", err)
	defer rows.Close()

	samples, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Sample])
	utils.HandleError(i.Logger, "error parsing rows to Sample struct", err)

	c.JSON(http.StatusOK, gin.H{"data": samples})
}
