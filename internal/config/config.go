package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTPConfig  *httpConfig // contains a config for the http server
	PGConfig    *pgConfig   // contains a config for postgres
	JWT_SECRET  string
	PASS_SECRET string
}

// New loads env for db, secrets and http server from .env
// and returns the configuration structure for the service.
func New(path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}
	p, err := newPGConfig()
	if err != nil {
		return nil, err
	}
	h, err := newHTTPConfig()
	if err != nil {
		return nil, err
	}
	return &Config{
		HTTPConfig:  h,
		PGConfig:    p,
		JWT_SECRET:  os.Getenv("JWT_SECRET_KEY"),
		PASS_SECRET: os.Getenv("PASS_HASH_SECRET"),
	}, nil
}
