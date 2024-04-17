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
	"github.com/gogapopp/Skoof/internal/handler/middlewares"
	"github.com/gogapopp/Skoof/internal/lib/logger"
	"github.com/gogapopp/Skoof/internal/service"
	"github.com/gogapopp/Skoof/internal/storage/postgres"
)

func main() {
	logger, err := logger.New()
	if err != nil {
		logger.Fatal(err)
	}
	config, err := config.New(".env")
	if err != nil {
		logger.Fatal(err)
	}

	db, err := postgres.New(config.PGConfig.DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	service := service.New(config.PASS_SECRET, config.JWT_SECRET, db)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(httprate.Limit(3, time.Second))

	r.Get("/", handler.HomePage(logger))

	r.Get("/signin", handler.SignInPage(logger, service))
	r.Post("/signin", handler.SignInPage(logger, service))

	r.Get("/signup", handler.SignUpPage(logger, service))
	r.Post("/signup", handler.SignUpPage(logger, service))

	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)

		r.Get("/skoof", handler.SkoofPage(logger))
		r.Post("/skoof", handler.SkoofPage(logger))
	})

	srv := http.Server{
		Addr:    config.HTTPConfig.Addr,
		Handler: r,
	}
	log.Println("server running at", config.HTTPConfig.Addr)
	log.Fatal(srv.ListenAndServe())
}
