package config

import (
	"errors"
	"os"
)

// keys for env
const (
	dsnEnvProd  = "PG_DSN_PROD"
	dsnEnvLocal = "PG_DSN_LOCAL"
)

type pgConfig struct {
	DSN string
}

func newPGConfig(env string) (*pgConfig, error) {
	var dsn string
	if env == "prod" {
		dsn = os.Getenv(dsnEnvProd)
	} else if env == "local" {
		dsn = os.Getenv(dsnEnvLocal)
	}
	if len(dsn) == 0 {
		return &pgConfig{}, errors.New("empty PG_DSN env")
	}
	return &pgConfig{DSN: dsn}, nil
}
