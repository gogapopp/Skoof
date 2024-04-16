package config

import (
	"errors"
	"os"
)

const dsnEnv = "PG_DSN"

type pgConfig struct {
	DSN string
}

func newPGConfig() (*pgConfig, error) {
	dsn := os.Getenv(dsnEnv)
	if len(dsn) == 0 {
		return &pgConfig{}, errors.New("empty PG_DSN env")
	}
	return &pgConfig{DSN: dsn}, nil
}
