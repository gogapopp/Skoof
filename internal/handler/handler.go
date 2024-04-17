package handler

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gogapopp/Skoof/internal/config"
	"github.com/gogapopp/Skoof/internal/handler/middlewares"
	"go.uber.org/zap"
)

const ReadHeaderTimeout = 10 * time.Second

func Routes(r *chi.Mux, logger *zap.SugaredLogger, authService authService, config *config.Config) *http.Server {
	r.Get("/", HomePage(logger))

	r.Get("/signin", SignInPage(logger, authService))
	r.Post("/signin", SignInPage(logger, authService))

	r.Get("/signup", SignUpPage(logger, authService))
	r.Post("/signup", SignUpPage(logger, authService))

	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)

		r.Get("/skoof", SkoofPage(logger))
		r.Post("/skoof", SkoofPage(logger))
	})

	srv := &http.Server{
		Addr:              config.HTTPConfig.Addr,
		Handler:           r,
		ReadHeaderTimeout: ReadHeaderTimeout,
	}

	return srv
}
