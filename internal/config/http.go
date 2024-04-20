package config

import (
	"errors"
	"net"
	"os"
)

// keys for env
const (
	httpHost = "HTTP_HOST"
	httpPort = "HTTP_PORT"
)

type httpConfig struct {
	Addr string
}

func newHTTPConfig() (*httpConfig, error) {
	var (
		host = os.Getenv(httpHost)
		port = os.Getenv(httpPort)
	)
	if len(port) == 0 {
		return nil, errors.New("empty HTTP_PORT env")
	}
	addr := net.JoinHostPort(host, port)
	return &httpConfig{
		Addr: addr,
	}, nil
}
