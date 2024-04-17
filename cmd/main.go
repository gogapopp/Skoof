package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
	"github.com/gogapopp/Skoof/internal/config"
	"github.com/gogapopp/Skoof/internal/handler"
	"github.com/gogapopp/Skoof/internal/lib/logger"
	"github.com/gogapopp/Skoof/internal/service"
	"github.com/gogapopp/Skoof/internal/storage/postgres"
)

func main() {
	var (
		logger = must(logger.New())
		config = must(config.New(".env"))
		db     = must(postgres.New(config.PGConfig.DSN))

		authService = service.New(config.PASS_SECRET, config.JWT_SECRET, db)
		r           = chi.NewRouter()
	)
	defer db.Close()
	// initializes middlewares
	mdwrs(r, middleware.RequestID, middleware.Logger, httprate.Limit(5, time.Second))
	// initializes server routes and returns a completed http server
	srv := handler.Routes(r, logger, authService, config)
	log.Println("server running at", config.HTTPConfig.Addr)
	log.Fatal(srv.ListenAndServe())
}

func mdwrs(r *chi.Mux, middlewares ...func(http.Handler) http.Handler) {
	for _, mdwr := range middlewares {
		r.Use(mdwr)
	}
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}
