package main

import (
	"os"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
	"github.com/gogapopp/Skoof/internal/config"
	httpserver "github.com/gogapopp/Skoof/internal/http_server"
	"github.com/gogapopp/Skoof/internal/libs/logger"
	"github.com/gogapopp/Skoof/internal/service"
	"github.com/gogapopp/Skoof/internal/storage/postgres"
	"github.com/gogapopp/Skoof/internal/storage/postgres/auth"
	"github.com/gogapopp/Skoof/internal/storage/postgres/community"
)

func main() {
	var (
		logger = must(logger.New())
		config = must(config.New(".env"))

		conn        = must(postgres.New(config.PGConfig.DSN))
		authDB      = auth.NewAuthStorage(conn)
		communityDB = community.NewCommunityStorage(conn)

		authService = service.New(config.PASS_SECRET, config.JWT_SECRET, authDB)
		_           = communityDB

		r = chi.NewRouter()
	)
	defer conn.Close()

	// Initializes middlewares for all server requests,
	// other middlewares can be initialized in the New function, see httpserver.New().
	r.Use(
		middleware.RequestID,
		middleware.Logger,
		httprate.Limit(5, time.Second),
	)

	// Initializes server routes and returns a completed http server.
	server := httpserver.New(r, logger, config, authService)

	logger.Infof("runnig server at: %s", config.HTTPConfig.Addr)
	if err := server.ListenAndServe(); err != nil {
		logger.Errorf("server shutdown with error: %w", err)
		os.Exit(1)
	}
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}
