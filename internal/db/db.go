package db

import (
	"context"

	"github.com/angelvargass/go-api/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(ctx context.Context, config *config.DBConfig) (*pgxpool.Pool, error) {
	poolConfig := createConfigurationObject()
	connConfig, err := pgxpool.ParseConfig(config.DBUrl)
	if err != nil {
		return nil, err
	}

	connConfig.MaxConnLifetime = poolConfig.MaxConnLifeTime
	connConfig.MaxConnIdleTime = poolConfig.MaxConnIdleTime
	connConfig.HealthCheckPeriod = poolConfig.HealthCheckPeriod
	connConfig.MaxConns = poolConfig.MaxConns
	connConfig.MinConns = poolConfig.MinConns

	pool, err := pgxpool.NewWithConfig(ctx, connConfig)
	return pool, err
}

func createConfigurationObject() *Config {
	return &Config{
		// number of vCPU (2) on the neonDB cluster, and x4 for the number of connections (4 is an estimation that may perform best for this workload)
		// goal is to maintain between 2x and 4x connections open based on the number of CPUs of the database cluster
		MaxConns:          2 * 4,
		MinConns:          2 * 2,
		MaxConnLifeTime:   30,
		MaxConnIdleTime:   10,
		HealthCheckPeriod: 5,
	}
}
