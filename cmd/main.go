package main

import (
	"os"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
	"github.com/gogapopp/Skoof/internal/config"
	"github.com/gogapopp/Skoof/internal/handler"
	"github.com/gogapopp/Skoof/internal/libs/logger"
	"github.com/gogapopp/Skoof/internal/service"
	"github.com/gogapopp/Skoof/internal/storage/postgres"
)

func main() {
	var (
		logger      = must(logger.New())
		config      = must(config.New(".env"))
		database    = must(postgres.New(config.PGConfig.DSN))
		authService = service.New(config.PASS_SECRET, config.JWT_SECRET, database)
		r           = chi.NewRouter()
	)
	defer database.Close()

	// Initializes middlewares for all server requests,
	// other middlewares can be initialized in the Routes function, see handler.Routes.
	r.Use(
		middleware.RequestID,
		middleware.Logger,
		httprate.Limit(5, time.Second),
	)

	// Initializes server routes and returns a completed http server.
	srv := handler.Routes(r, logger, authService, config)

	logger.Infof("runnig server at: %s", config.HTTPConfig.Addr)
	if err := srv.ListenAndServe(); err != nil {
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
