package config

import (
	"log/slog"

	"github.com/kelseyhightower/envconfig"
)

func New() (*Config, error) {
	c := new(Config)
	err := envconfig.Process("go-api", c)
	if err != nil {
		slog.Error("error reading env variables", slog.String("error", err.Error()))
		return nil, err
	}

	return c, nil
}
